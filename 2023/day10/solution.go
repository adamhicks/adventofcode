package day10

import (
	_ "embed"
	"fmt"
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

type coord struct {
	X, Y int
}

func connectsTo(r rune, c coord) ([2]coord, bool) {
	switch r {
	case '|':
		return [...]coord{{c.X, c.Y - 1}, {c.X, c.Y + 1}}, true
	case '-':
		return [...]coord{{c.X - 1, c.Y}, {c.X + 1, c.Y}}, true
	case 'L':
		return [...]coord{{c.X, c.Y - 1}, {c.X + 1, c.Y}}, true
	case 'J':
		return [...]coord{{c.X, c.Y - 1}, {c.X - 1, c.Y}}, true
	case '7':
		return [...]coord{{c.X - 1, c.Y}, {c.X, c.Y + 1}}, true
	case 'F':
		return [...]coord{{c.X, c.Y + 1}, {c.X + 1, c.Y}}, true
	}
	return [...]coord{{}, {}}, false
}

func neighbours(c coord) []coord {
	return []coord{
		{c.X, c.Y - 1},
		{c.X - 1, c.Y},
		{c.X + 1, c.Y},
		{c.X, c.Y + 1},
	}
}

func connect(maze input) map[coord][2]coord {
	ret := make(map[coord][2]coord)
	for y, l := range maze {
		for x, r := range l {
			c := coord{x, y}
			con, ok := connectsTo(r, c)
			if !ok {
				continue
			}
			ret[c] = con
		}
	}
	return ret
}

func fixStart(s coord, maze map[coord][2]coord) {
	var cur []coord
	for _, n := range neighbours(s) {
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
	maze[s] = [2]coord{cur[0], cur[1]}
}

func runMaze(s coord, maze map[coord][2]coord) map[coord][2]coord {
	fixStart(s, maze)
	loop := make(map[coord][2]coord)
	cur := []coord{s}
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

func findStart(s input) coord {
	for y, l := range s {
		for x, r := range l {
			if r == 'S' {
				return coord{x, y}
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

var (
	North = coord{Y: -1}
	East  = coord{X: 1}
	South = coord{Y: 1}
	West  = coord{X: -1}
)

func plus(a, b coord) coord {
	return coord{a.X + b.X, a.Y + b.Y}
}

func dir(a, b coord) coord {
	return coord{X: b.X - a.X, Y: b.Y - a.Y}
}

func toPaint(c coord, nxt coord) (coord, coord) {
	switch dir(c, nxt) {
	case North:
		return plus(c, West), plus(c, East)
	case East:
		return plus(c, North), plus(c, South)
	case South:
		return plus(c, East), plus(c, West)
	case West:
		return plus(c, South), plus(c, North)
	}
	fmt.Println(c, nxt)
	panic("unreachable")
}

func inRange(c coord, maximum coord) bool {
	return c.X >= 0 && c.X < maximum.X && c.Y >= 0 && c.Y < maximum.Y
}

func fill(painted map[coord]int, s coord, maximum coord) {
	todo := []coord{s}
	for len(todo) > 0 {
		c := todo[0]
		todo = todo[1:]
		for _, n := range neighbours(c) {
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

func paint(from, to coord, painted map[coord]int) []coord {
	var toFill []coord
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

	painted := make(map[coord]int)
	var toFill []coord

	cur := start
	nxt := loop[start][0]
	prev := loop[start][1]

	for {
		painted[cur] = 0
		toFill = append(toFill, paint(cur, nxt, painted)...)
		wasGoing := dir(prev, cur)
		nowGoing := dir(cur, nxt)
		if wasGoing != nowGoing {
			proj := plus(cur, wasGoing)
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
	maximum := coord{len(s[0]), len(s)}
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
