package day11

import (
	_ "embed"
	"fmt"
	"github.com/adamhicks/adventofcode/2023/aoc"
	"strings"
)

//go:embed input.txt
var inputString string

type input []string

func parseInput(s string) input {
	return strings.Split(strings.TrimSpace(s), "\n")
}

var testString1 = `...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`

func blanksCrossed(a, b int, counts []int) int {
	var zeroes int

	if a > b {
		a, b = b, a
	}

	for _, v := range counts[a:b] {
		if v == 0 {
			zeroes++
		}
	}
	return zeroes
}

func parseMap(s input) ([]aoc.Vec2, []int, []int) {
	rowCount := make([]int, len(s))
	colCount := make([]int, len(s[0]))

	var gals []aoc.Vec2
	for y, l := range s {
		for x, r := range l {
			if r != '#' {
				continue
			}
			rowCount[y]++
			colCount[x]++
			gals = append(gals, aoc.Vec2{X: x, Y: y})
		}
	}
	return gals, rowCount, colCount
}

func runPartOne(s input) error {
	gals, rowCount, colCount := parseMap(s)

	var sum int
	for i := 0; i < len(gals); i++ {
		for j := i + 1; j < len(gals); j++ {
			a, b := gals[i], gals[j]
			sum += a.Distance(b)
			sum += blanksCrossed(a.Y, b.Y, rowCount)
			sum += blanksCrossed(a.X, b.X, colCount)
		}
	}
	fmt.Println(sum)

	return nil
}

var testString2 = testString1

func runPartTwo(s input) error {
	gals, rowCount, colCount := parseMap(s)

	var sum int
	for i := 0; i < len(gals); i++ {
		for j := i + 1; j < len(gals); j++ {
			a, b := gals[i], gals[j]
			sum += a.Distance(b)
			sum += blanksCrossed(a.Y, b.Y, rowCount) * 999999
			sum += blanksCrossed(a.X, b.X, colCount) * 999999
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
