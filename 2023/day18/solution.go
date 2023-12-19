package day18

import (
	_ "embed"
	"fmt"
	"github.com/adamhicks/adventofcode/2023/aoc"
	"strconv"
	"strings"
)

//go:embed input.txt
var inputString string

type input []string

func parseInput(s string) input {
	return strings.Split(strings.TrimSpace(s), "\n")
}

var testString1 = `R 6 (#70c710)
D 5 (#0dc571)
L 2 (#5713f0)
D 2 (#d2c081)
R 2 (#59c680)
D 2 (#411b91)
L 5 (#8ceee2)
U 2 (#caa173)
L 1 (#1b58a2)
U 2 (#caa171)
R 2 (#7807d2)
U 3 (#a77fa3)
L 2 (#015232)
U 2 (#7a21e3)`

func decodeLine(s string) (aoc.Vec2, int64) {
	parts := strings.Fields(s)
	var dir aoc.Vec2
	switch parts[0] {
	case "U":
		dir = aoc.North
	case "D":
		dir = aoc.South
	case "R":
		dir = aoc.East
	case "L":
		dir = aoc.West
	default:
		panic("unknown direction")
	}
	length := aoc.Must(strconv.ParseInt(parts[1], 10, 64))
	return dir, length
}

func runPartOne(s input) error {
	fmt.Println(calcArea(s, decodeLine))
	return nil
}

var testString2 = testString1

func decodeHex(s string) (aoc.Vec2, int64) {
	s = strings.Trim(strings.Fields(s)[2], "()#")
	i := s[len(s)-1:]
	s = s[:len(s)-1]
	var dir aoc.Vec2
	switch i {
	case "0":
		dir = aoc.East
	case "1":
		dir = aoc.South
	case "2":
		dir = aoc.West
	case "3":
		dir = aoc.North
	}
	length := aoc.Must(strconv.ParseInt(s, 16, 64))
	return dir, length
}

type parseLineFunc func(string) (aoc.Vec2, int64)

func shoelace(points []aoc.Vec2) int64 {
	var sum int64
	for i := 0; i < len(points); i++ {
		j := (i + 1) % len(points)
		pI, pJ := points[i], points[j]
		sum += int64(pI.X*pJ.Y) - int64(pI.Y*pJ.X)
	}
	return int64(aoc.Abs(int(sum))) / 2
}

func calcArea(s input, parse parseLineFunc) int64 {
	var perimeter int64
	var pos aoc.Vec2
	points := make([]aoc.Vec2, 0, len(s)+1)
	points = append(points, pos)
	for _, l := range s {
		dir, length := parse(l)
		pos = pos.Add(dir.Mul(int(length)))
		perimeter += length
		points = append(points, pos)
	}
	area := shoelace(points)
	return area + perimeter/2 + 1
}

func runPartTwo(s input) error {
	fmt.Println(calcArea(s, decodeHex))
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
