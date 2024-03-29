package main

import (
	"context"
	_ "embed"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	"text/template"
	"time"
)

var nycTimezone *time.Location

func init() {
	tz, err := time.LoadLocation("America/New_York")
	if err != nil {
		panic(err)
	}
	nycTimezone = tz
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:        "start",
				Description: "start a problem",
				ArgsUsage:   "[days e.g. 1 2 3]",
				Action:      start,
			},
			{
				Name:        "run",
				Description: "run a solution",
				ArgsUsage:   "[days e.g. 1 2 3]",
				Action:      run,
			},
		},
	}
	ctx := context.Background()
	if err := app.RunContext(ctx, os.Args); err != nil {
		log.Fatal(err)
	}
}

func getDays(args []string) ([]int, error) {
	var days []int
	for _, v := range args {
		d, err := strconv.Atoi(v)
		if err != nil {
			return nil, fmt.Errorf("invalid day number %s: %w", v, err)
		}
		if d < 1 || d > 31 {
			return nil, fmt.Errorf("invalid day %d", d)
		}
		days = append(days, d)
	}
	if len(days) == 0 {
		day, err := getCurrentDay()
		if err != nil {
			return nil, err
		}
		days = append(days, day)
	}
	return days, nil
}

func getCurrentDay() (int, error) {
	t := time.Now().In(nycTimezone)
	if t.Month() != 12 {
		return 0, fmt.Errorf("we are not in December yet, calm down")
	}
	return t.Day(), nil
}

func start(ctx *cli.Context) error {
	token, err := getToken()
	if err != nil {
		return err
	}
	days, err := getDays(ctx.Args().Slice())
	if err != nil {
		return err
	}
	for _, day := range days {
		err := startDay(ctx.Context, token, 2023, day)
		if err != nil {
			return err
		}
	}
	return generateSolutionMap()
}

type DaySolution interface {
	TestPart1() error
	RunPart1() error
	TestPart2() error
	RunPart2() error
}

func run(ctx *cli.Context) error {
	days, err := getDays(ctx.Args().Slice())
	if err != nil {
		return err
	}
	for _, day := range days {
		s, ok := daySolutions[day]
		if !ok {
			return fmt.Errorf("solution for day %d not present", day)
		}
		runSolution(day, s)
	}
	return nil
}

func fmtPartName(day, part int, sample bool) string {
	n := fmt.Sprintf("day %d part %d", day, part)
	if sample {
		n += " - sample"
	}
	return n
}

func runSolution(day int, s DaySolution) {
	runPart(fmtPartName(day, 1, true), s.TestPart1)
	runPart(fmtPartName(day, 1, false), s.RunPart1)
	runPart(fmtPartName(day, 2, true), s.TestPart2)
	runPart(fmtPartName(day, 2, false), s.RunPart2)
}

func runPart(name string, f func() error) {
	start := time.Now()
	log.Println(name)
	err := f()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(time.Since(start))
	log.Println("----")
}

func getToken() (string, error) {
	t := os.Getenv("GITHUB_TOKEN")
	if t == "" {
		return "", fmt.Errorf("no github token")
	}
	return t, nil
}

func startDay(ctx context.Context, token string, y, d int) error {
	log.Println("setting up day", d)
	dir := getDayDir(d)
	if err := makeDir(dir); err != nil {
		return err
	}
	if err := createSolutionFile(dir); err != nil {
		return err
	}
	in, err := downloadDayInput(ctx, token, y, d)
	if err != nil {
		return err
	}
	if err := writeInputFile(dir, in); err != nil {
		return err
	}
	fmt.Println("Good Luck!!!")
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d", y, d)
	fmt.Println(url)
	return nil
}

func getDayDir(d int) string {
	return fmt.Sprintf("day%02d", d)
}

func dirExists(dir string) bool {
	i, err := os.Stat(dir)
	if os.IsNotExist(err) {
		return false
	} else if err != nil {
		panic(err)
	}
	return i.IsDir()
}

func makeDir(dir string) error {
	if dirExists(dir) {
		return fmt.Errorf("directory %s already exists", dir)
	}
	return os.Mkdir(dir, 0o777)
}

//go:embed solution.gotmpl
var solutionFile string
var solutionTemplate = template.Must(template.New("solution").Parse(solutionFile))

type solution struct {
	PackageName string
}

func createSolutionFile(dir string) error {
	p := path.Join(dir, "solution.go")
	f, err := os.Create(p)
	if err != nil {
		return err
	}
	s := solution{PackageName: path.Base(dir)}
	return solutionTemplate.Execute(f, s)
}

func downloadDayInput(ctx context.Context, token string, y int, d int) (io.ReadCloser, error) {
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", y, d)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.AddCookie(&http.Cookie{Name: "session", Value: token})
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("unexpected response [%d]: '%s'", resp.StatusCode, string(body))
	}
	return resp.Body, nil
}

func writeInputFile(dir string, in io.ReadCloser) error {
	p := path.Join(dir, "input.txt")
	f, err := os.Create(p)
	if err != nil {
		return err
	}
	_, err = io.Copy(f, in)
	return err
}

var mapTemplateString = `// DO NOT EDIT: File generated by aoc 2023
package main

import (
{{- range $day, $package := . }}
	"github.com/adamhicks/adventofcode/2023/{{$package}}"
{{- end }}
)

var daySolutions = map[int]DaySolution{
{{- range $day, $package := . }}
	{{$day}}: {{$package}}.Solution{},
{{- end }}
}
`
var mapTemplate = template.Must(template.New("days").Parse(mapTemplateString))

func generateSolutionMap() error {
	m := make(map[int]string)
	for i := 1; i <= 31; i++ {
		dir := getDayDir(i)
		if dirExists(dir) {
			m[i] = dir
		}
	}
	f, err := os.Create("days.go")
	if err != nil {
		return err
	}
	return mapTemplate.Execute(f, m)
}
