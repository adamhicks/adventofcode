package day06

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var inputString string

type input []string

func parseInput(s string) input {
	return strings.Split(strings.TrimSpace(s), "\n")
}

var testString1 = `Time:      7  15   30
Distance:  9  40  200`

func parseLine(l string) ([]int, error) {
	var ret []int
	for _, v := range strings.Fields(l)[1:] {
		num, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		ret = append(ret, num)
	}
	return ret, nil
}

func isWin(hold, ti, di int) bool {
	return (ti-hold)*hold > di
}

func wins(ti, di int) int {
	var winAt int
	for ; winAt < ti; winAt++ {
		if isWin(winAt, ti, di) {
			break
		}
	}
	return ti - 2*winAt + 1
}

func runPartOne(s input) error {
	times, err := parseLine(s[0])
	if err != nil {
		return err
	}
	distances, err := parseLine(s[1])
	if err != nil {
		return err
	}

	total := 1
	for i, ti := range times {
		di := distances[i]
		wi := wins(ti, di)
		total *= wi
	}
	fmt.Println(total)
	return nil
}

var testString2 = testString1

func parseSingleRace(l string) (int, error) {
	_, val, _ := strings.Cut(l, ":")
	val = strings.ReplaceAll(val, " ", "")
	return strconv.Atoi(val)
}

func runPartTwo(s input) error {
	ti, err := parseSingleRace(s[0])
	if err != nil {
		return err
	}
	di, err := parseSingleRace(s[1])
	if err != nil {
		return err
	}
	wi := wins(ti, di)
	fmt.Println(wi)
	return nil
}

type Solution struct{}

func (Solution) TestPart1() error {
	return runPartOne(parseInput(testString1))
}

func (Solution) RunPart1() error {
	return runPartOne(parseInput(inputString))
}

func (Solution) TestPart2() error {
	return runPartTwo(parseInput(testString2))
}

func (Solution) RunPart2() error {
	return runPartTwo(parseInput(inputString))
}
