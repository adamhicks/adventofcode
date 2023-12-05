package day03

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

//go:embed input.txt
var inputString string

type input []string

func parseInput(s string) input {
	return strings.Split(strings.TrimSpace(s), "\n")
}

var testString1 = `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

type coord struct {
	X, Y int
}

func neighbours(c coord) []coord {
	return []coord{
		{X: c.X - 1, Y: c.Y - 1},
		{X: c.X, Y: c.Y - 1},
		{X: c.X + 1, Y: c.Y - 1},
		{X: c.X - 1, Y: c.Y},
		{X: c.X + 1, Y: c.Y},
		{X: c.X - 1, Y: c.Y + 1},
		{X: c.X, Y: c.Y + 1},
		{X: c.X + 1, Y: c.Y + 1},
	}
}

func runPartOne(s input) error {
	symbols := make(map[coord]rune)
	for y, l := range s {
		for x, c := range l {
			if !unicode.IsDigit(c) && c != '.' {
				symbols[coord{X: x, Y: y}] = c
			}
		}
	}
	var sum int
	for y, line := range s {
		line = line + "."
		idx := -1
		var isPart bool
		for x, chr := range line {
			if unicode.IsDigit(chr) {
				if idx == -1 {
					idx = x
				}
				for _, c := range neighbours(coord{X: x, Y: y}) {
					if _, ok := symbols[c]; ok {
						isPart = true
					}
				}
			} else if idx != -1 {
				s := line[idx:x]
				num, err := strconv.Atoi(s)
				if err != nil {
					return err
				}
				if isPart {
					sum += num
				}
				idx = -1
				isPart = false
			}
		}
	}
	fmt.Println(sum)
	return nil
}

var testString2 = `
467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..
`

func runPartTwo(s input) error {
	symbols := make(map[coord]rune)
	for y, l := range s {
		for x, c := range l {
			if !unicode.IsDigit(c) && c != '.' {
				symbols[coord{X: x, Y: y}] = c
			}
		}
	}

	gears := make(map[coord][]int)

	for y, line := range s {
		line = line + "."
		idx := -1

		touchGears := make(map[coord]struct{})

		for x, chr := range line {
			if unicode.IsDigit(chr) {
				if idx == -1 {
					idx = x
				}
				for _, c := range neighbours(coord{X: x, Y: y}) {
					if symbols[c] == '*' {
						touchGears[c] = struct{}{}
					}
				}
			} else if idx != -1 {
				s := line[idx:x]
				num, err := strconv.Atoi(s)
				if err != nil {
					return err
				}
				for c := range touchGears {
					gears[c] = append(gears[c], num)
				}
				idx = -1
				touchGears = make(map[coord]struct{})
			}
		}
	}

	var sum int
	for _, nums := range gears {
		if len(nums) != 2 {
			continue
		}
		sum += nums[0] * nums[1]
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
