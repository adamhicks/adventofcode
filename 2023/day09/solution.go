package day09

import (
	_ "embed"
	"fmt"
	"log"
	"strconv"
	"strings"
)

//go:embed input.txt
var inputString string

type input [][]int

func parseInput(s string) input {
	var ret input
	for _, l := range strings.Split(strings.TrimSpace(s), "\n") {
		var vals []int
		for _, num := range strings.Split(l, " ") {
			val, err := strconv.Atoi(num)
			if err != nil {
				log.Fatal(err)
			}
			vals = append(vals, val)
		}
		ret = append(ret, vals)
	}
	return ret
}

var testString1 = `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`

func extrapolateLine(vals []int) int {
	if len(vals) == 0 {
		return 0
	} else if len(vals) == 1 {
		return vals[0]
	}
	var diffs []int
	for i := 1; i < len(vals); i++ {
		diffs = append(diffs, vals[i]-vals[i-1])
	}
	return vals[len(vals)-1] + extrapolateLine(diffs)
}

func runPartOne(s input) error {
	var sum int
	for _, l := range s {
		sum += extrapolateLine(l)
	}
	fmt.Println(sum)
	return nil
}

var testString2 = testString1

func extrapolateLineStart(vals []int) int {
	if len(vals) == 0 {
		return 0
	} else if len(vals) == 1 {
		return vals[0]
	}
	var diffs []int
	for i := 1; i < len(vals); i++ {
		diffs = append(diffs, vals[i]-vals[i-1])
	}
	r := vals[0] - extrapolateLineStart(diffs)
	return r
}

func runPartTwo(s input) error {
	var sum int
	for _, l := range s {
		sum += extrapolateLineStart(l)
	}
	fmt.Println(sum)
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
