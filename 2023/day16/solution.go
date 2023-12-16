package day16

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

var testString1 = `.|...\....
|.-.\.....
.....|-...
........|.
..........
.........\
..../.\\..
.-.-/..|..
.|....-|.\
..//.|....`

var interact = map[rune]map[aoc.Vec2][]aoc.Vec2{
	'/': {
		aoc.North: []aoc.Vec2{aoc.East},
		aoc.South: []aoc.Vec2{aoc.West},
		aoc.East:  []aoc.Vec2{aoc.North},
		aoc.West:  []aoc.Vec2{aoc.South},
	},
	'\\': {
		aoc.North: []aoc.Vec2{aoc.West},
		aoc.South: []aoc.Vec2{aoc.East},
		aoc.East:  []aoc.Vec2{aoc.South},
		aoc.West:  []aoc.Vec2{aoc.North},
	},
	'-': {
		aoc.North: []aoc.Vec2{aoc.East, aoc.West},
		aoc.South: []aoc.Vec2{aoc.East, aoc.West},
		aoc.East:  []aoc.Vec2{aoc.East},
		aoc.West:  []aoc.Vec2{aoc.West},
	},
	'|': {
		aoc.North: []aoc.Vec2{aoc.North},
		aoc.South: []aoc.Vec2{aoc.South},
		aoc.East:  []aoc.Vec2{aoc.North, aoc.South},
		aoc.West:  []aoc.Vec2{aoc.North, aoc.South},
	},
}

type ray struct {
	Pos, Dir aoc.Vec2
}

func next(r ray, m map[aoc.Vec2]rune) []ray {
	c := m[r.Pos]
	i, ok := interact[c]
	if !ok {
		return []ray{{Pos: r.Pos.Add(r.Dir), Dir: r.Dir}}
	}
	var ret []ray
	for _, d := range i[r.Dir] {
		nxt := ray{Pos: r.Pos.Add(d), Dir: d}
		ret = append(ret, nxt)
	}
	return ret
}

func inRange(p, maxim aoc.Vec2) bool {
	return p.X >= 0 && p.X < maxim.X && p.Y >= 0 && p.Y < maxim.Y
}

func energised(s ray, maxim aoc.Vec2, m map[aoc.Vec2]rune) int {
	ener := make(map[aoc.Vec2]bool)
	done := make(map[ray]bool)
	q := []ray{s}
	for len(q) > 0 {
		n := q[0]
		q = q[1:]
		ener[n.Pos] = true
		for _, nxt := range next(n, m) {
			if !inRange(nxt.Pos, maxim) || done[nxt] {
				continue
			}
			done[nxt] = true
			q = append(q, nxt)
		}
	}
	return len(ener)
}

func getMap(s input) map[aoc.Vec2]rune {
	m := make(map[aoc.Vec2]rune)
	for y, l := range s {
		for x, r := range l {
			if r == '.' {
				continue
			}
			m[aoc.Vec2{X: x, Y: y}] = r
		}
	}
	return m
}

func runPartOne(s input) error {
	maxim := aoc.Vec2{X: len(s[0]), Y: len(s)}
	m := getMap(s)
	fmt.Println(energised(ray{Dir: aoc.East}, maxim, m))
	return nil
}

var testString2 = testString1

func runPartTwo(s input) error {
	maxim := aoc.Vec2{X: len(s[0]), Y: len(s)}
	m := getMap(s)

	var most int
	for x := 0; x < maxim.X; x++ {
		down := ray{Pos: aoc.Vec2{X: x}, Dir: aoc.South}
		most = max(most, energised(down, maxim, m))

		up := ray{Pos: aoc.Vec2{X: x, Y: maxim.Y - 1}, Dir: aoc.North}
		most = max(most, energised(up, maxim, m))
	}
	for y := 0; y < maxim.Y; y++ {
		left := ray{Pos: aoc.Vec2{Y: y}, Dir: aoc.East}
		most = max(most, energised(left, maxim, m))

		right := ray{Pos: aoc.Vec2{X: maxim.X - 1, Y: y}, Dir: aoc.West}
		most = max(most, energised(right, maxim, m))
	}

	fmt.Println(most)
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
