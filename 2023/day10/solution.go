package day10

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

var testString1 = `..F7.
.FJ|.
SJ.L7
|F--J
LJ...`

func connectsTo(r rune, c aoc.Vec2) ([2]aoc.Vec2, bool) {
	switch r {
	case '|':
		return [...]aoc.Vec2{{c.X, c.Y - 1}, {c.X, c.Y + 1}}, true
	case '-':
		return [...]aoc.Vec2{{c.X - 1, c.Y}, {c.X + 1, c.Y}}, true
	case 'L':
		return [...]aoc.Vec2{{c.X, c.Y - 1}, {c.X + 1, c.Y}}, true
	case 'J':
		return [...]aoc.Vec2{{c.X, c.Y - 1}, {c.X - 1, c.Y}}, true
	case '7':
		return [...]aoc.Vec2{{c.X - 1, c.Y}, {c.X, c.Y + 1}}, true
	case 'F':
		return [...]aoc.Vec2{{c.X, c.Y + 1}, {c.X + 1, c.Y}}, true
	}
	return [...]aoc.Vec2{{}, {}}, false
}

func connect(maze input) map[aoc.Vec2][2]aoc.Vec2 {
	ret := make(map[aoc.Vec2][2]aoc.Vec2)
	for y, l := range maze {
		for x, r := range l {
			c := aoc.Vec2{X: x, Y: y}
			con, ok := connectsTo(r, c)
			if !ok {
				continue
			}
			ret[c] = con
		}
	}
	return ret
}

func fixStart(s aoc.Vec2, maze map[aoc.Vec2][2]aoc.Vec2) {
	var cur []aoc.Vec2
	for _, n := range s.Orthogonal() {
		con := maze[n]
		for _, c := range con {
			if c == s {
				cur = append(cur, n)
			}
		}
	}
	if len(cur) != 2 {
		panic("whoops")
	}
	maze[s] = [2]aoc.Vec2{cur[0], cur[1]}
}

func runMaze(s aoc.Vec2, maze map[aoc.Vec2][2]aoc.Vec2) map[aoc.Vec2][2]aoc.Vec2 {
	fixStart(s, maze)
	loop := make(map[aoc.Vec2][2]aoc.Vec2)
	cur := []aoc.Vec2{s}
	for len(cur) > 0 {
		nxt := cur[0]
		cur = cur[1:]
		loop[nxt] = maze[nxt]
		for _, c := range maze[nxt] {
			if _, ok := loop[c]; ok {
				continue
			}
			cur = append(cur, c)
		}
	}

	return loop
}

func findStart(s input) aoc.Vec2 {
	for y, l := range s {
		for x, r := range l {
			if r == 'S' {
				return aoc.Vec2{X: x, Y: y}
			}
		}
	}
	panic("no start")
}

func runPartOne(s input) error {
	start := findStart(s)
	loop := runMaze(start, connect(s))
	fmt.Println(len(loop) / 2)
	return nil
}

var testString2 = `FF7FSF7F7F7F7F7F---7
L|LJ||||||||||||F--J
FL-7LJLJ||||||LJL-77
F--JF--7||LJLJ7F7FJ-
L---JF-JLJ.||-FJLJJ7
|F|F-JF---7F7-L7L|7|
|FFJF7L7F-JF7|JL---7
7-L-JL7||F7|L7F-7F7|
L.L7LFJ|||||FJL7||LJ
L7JLJL-JLJLJL--JLJ.L`

func toPaint(c aoc.Vec2, nxt aoc.Vec2) (aoc.Vec2, aoc.Vec2) {
	switch nxt.Sub(c) {
	case aoc.North:
		return c.Add(aoc.West), c.Add(aoc.East)
	case aoc.East:
		return c.Add(aoc.North), c.Add(aoc.South)
	case aoc.South:
		return c.Add(aoc.East), c.Add(aoc.West)
	case aoc.West:
		return c.Add(aoc.South), c.Add(aoc.North)
	}
	fmt.Println(c, nxt)
	panic("unreachable")
}

func inRange(c aoc.Vec2, maximum aoc.Vec2) bool {
	return c.X >= 0 && c.X < maximum.X && c.Y >= 0 && c.Y < maximum.Y
}

func fill(painted map[aoc.Vec2]int, s aoc.Vec2, maximum aoc.Vec2) {
	todo := []aoc.Vec2{s}
	for len(todo) > 0 {
		c := todo[0]
		todo = todo[1:]
		for _, n := range c.Orthogonal() {
			if _, ok := painted[n]; ok {
				continue
			}
			if !inRange(n, maximum) {
				continue
			}
			painted[n] = painted[s]
			todo = append(todo, n)
		}
	}
}

func paint(from, to aoc.Vec2, painted map[aoc.Vec2]int) []aoc.Vec2 {
	var toFill []aoc.Vec2
	left, right := toPaint(from, to)
	if _, ok := painted[left]; !ok {
		painted[left] = 1
		toFill = append(toFill, left)
	}
	if _, ok := painted[right]; !ok {
		painted[right] = 2
		toFill = append(toFill, right)
	}
	return toFill
}

func runPartTwo(s input) error {
	start := findStart(s)
	loop := runMaze(start, connect(s))

	painted := make(map[aoc.Vec2]int)
	var toFill []aoc.Vec2

	cur := start
	nxt := loop[start][0]
	prev := loop[start][1]

	for {
		painted[cur] = 0
		toFill = append(toFill, paint(cur, nxt, painted)...)
		wasGoing := cur.Sub(prev)
		nowGoing := nxt.Sub(cur)
		if wasGoing != nowGoing {
			proj := cur.Add(wasGoing)
			toFill = append(toFill, paint(cur, proj, painted)...)
		}

		prev = cur
		a, b := loop[nxt][0], loop[nxt][1]
		if a == cur {
			cur = nxt
			nxt = b
		} else {
			cur = nxt
			nxt = a
		}

		for _, c := range loop[nxt] {
			if c == cur {
				continue
			}
		}
		if cur == start {
			break
		}
	}
	maximum := aoc.Vec2{X: len(s[0]), Y: len(s)}
	for _, f := range toFill {
		fill(painted, f, maximum)
	}

	var ones, twos int
	for _, c := range painted {
		switch c {
		case 1:
			ones++
		case 2:
			twos++
		}
	}
	fmt.Println(min(ones, twos))
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
