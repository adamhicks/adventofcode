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

func isVerticalReflected(s []string, idx int) int {
	var diff int
	for _, l := range s {
		i, j := idx, idx+1
		for {
			if i < 0 || j >= len(l) {
				break
			}
			if l[i] != l[j] {
				diff++
			}
			i--
			j++
		}
	}
	return diff
}

func isHorizontalReflected(s []string, idx int) int {
	var diff int
	i, j := idx, idx+1
	for {
		if i < 0 || j >= len(s) {
			break
		}
		l1, l2 := s[i], s[j]
		for r := range l1 {
			if l1[r] != l2[r] {
				diff++
			}
		}
		i--
		j++
	}
	return diff
}

func getReflection(b []string, expDiff int) (int, int) {
	for i := 0; i < len(b[0])-1; i++ {
		if isVerticalReflected(b, i) == expDiff {
			return i + 1, 0
		}
	}
	for i := 0; i < len(b)-1; i++ {
		if isHorizontalReflected(b, i) == expDiff {
			return 0, i + 1
		}
	}
	return 0, 0
}

func runPartOne(s input) error {
	var sum int
	for _, b := range s {
		v, h := getReflection(b, 0)
		if v > 0 {
			sum += v
		}
		if h > 0 {
			sum += h * 100
		}
	}
	fmt.Println(sum)
	return nil
}

var testString2 = testString1

func runPartTwo(s input) error {
	var sum int
	for _, b := range s {
		v, h := getReflection(b, 1)
		if v > 0 {
			sum += v
		}
		if h > 0 {
			sum += h * 100
		}
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
