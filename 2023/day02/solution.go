package day02

import (
	_ "embed"
	"fmt"
	"log"
	"strconv"
	"strings"
)

//go:embed input.txt
var inputString string

type input []string

func parseInput(s string) input {
	return strings.Split(strings.TrimSpace(s), "\n")
}

var testString1 = `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
`

type colors struct {
	Red, Blue, Green int
}

func parseGame(line string) ([]colors, error) {
	var ret []colors
	for _, g := range strings.Split(line, "; ") {
		var seen colors
		for _, c := range strings.Split(g, ", ") {
			val, colStr, _ := strings.Cut(c, " ")
			num, err := strconv.Atoi(val)
			if err != nil {
				return nil, err
			}
			switch colStr {
			case "red":
				seen.Red = num
			case "blue":
				seen.Blue = num
			case "green":
				seen.Green = num
			}
		}
		ret = append(ret, seen)
	}
	return ret, nil
}

func isPossible(seen, max colors) bool {
	return seen.Red <= max.Red && seen.Blue <= max.Blue && seen.Green <= max.Green
}

func runPartOne(s input) {
	maxSeen := colors{Red: 12, Green: 13, Blue: 14}
	var sum int
	for _, l := range s {
		id, game, _ := strings.Cut(l, ": ")
		_, gVal, _ := strings.Cut(id, " ")
		gameID, err := strconv.Atoi(gVal)
		if err != nil {
			log.Fatal(err)
		}
		gameSeen, err := parseGame(game)
		if err != nil {
			log.Fatal(err)
		}
		possible := true
		for _, seen := range gameSeen {
			if !isPossible(seen, maxSeen) {
				possible = false
			}
		}
		if possible {
			sum += gameID
		}
	}
	fmt.Println(sum)
}

var testString2 = `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`

func minPossible(seen []colors) colors {
	var ret colors
	for _, s := range seen {
		ret.Red = max(ret.Red, s.Red)
		ret.Green = max(ret.Green, s.Green)
		ret.Blue = max(ret.Blue, s.Blue)
	}
	return ret
}

func runPartTwo(s input) {
	var sum int
	for _, l := range s {
		_, game, _ := strings.Cut(l, ": ")
		gameSeen, err := parseGame(game)
		if err != nil {
			log.Fatal(err)
		}
		minPoss := minPossible(gameSeen)
		sum += minPoss.Red * minPoss.Green * minPoss.Blue
	}
	fmt.Println(sum)
}

type Solution struct{}

func (Solution) TestPart1() {
	runPartOne(parseInput(testString1))
}

func (Solution) RunPart1() {
	runPartOne(parseInput(inputString))
}

func (Solution) TestPart2() {
	runPartTwo(parseInput(testString2))
}
func (Solution) RunPart2() {
	runPartTwo(parseInput(inputString))
}
