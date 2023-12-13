package day13

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var inputString string

type input [][]string

func parseInput(s string) input {
	var ret input
	for _, b := range strings.Split(strings.TrimSpace(s), "\n\n") {
		ret = append(ret, strings.Split(b, "\n"))
	}
	return ret
}

var testString1 = `#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.

#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#
`

func reflectDistance(s []string, idx int) int {
	var diff int
	for _, l := range s {
		for i, j := idx, idx+1; i >= 0 && j < len(l); i, j = i-1, j+1 {
			if l[i] != l[j] {
				diff++
			}
		}
	}
	return diff
}

func flipped(b []string) []string {
	ret := make([]string, len(b[0]))
	for _, l := range b {
		for x, r := range l {
			ret[x] += string(r)
		}
	}
	return ret
}

func findReflection(b []string, expDiff int) (int, bool) {
	for i := 0; i < len(b[0])-1; i++ {
		if reflectDistance(b, i) == expDiff {
			return i, true
		}
	}
	return 0, false
}

func scoreReflection(b []string, expDiff int) int {
	if i, ok := findReflection(b, expDiff); ok {
		return i + 1
	}
	if i, ok := findReflection(flipped(b), expDiff); ok {
		return (i + 1) * 100
	}
	panic("no reflection found")
}

func runPartOne(s input) error {
	var sum int
	for _, b := range s {
		sum += scoreReflection(b, 0)
	}
	fmt.Println(sum)
	return nil
}

var testString2 = testString1

func runPartTwo(s input) error {
	var sum int
	for _, b := range s {
		sum += scoreReflection(b, 1)
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
