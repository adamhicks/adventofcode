package day14

import (
	"bytes"
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

var testString1 = `O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....
`

func parseMap(s input) [][]byte {
	ret := make([][]byte, 0, len(s))
	for _, l := range s {
		row := make([]byte, len(l))
		for x, r := range l {
			if r != '.' {
				row[x] = byte(r)
			}
		}
		ret = append(ret, row)
	}
	return ret
}

func inRange(p, max aoc.Vec2) bool {
	return p.X >= 0 && p.X < max.X && p.Y >= 0 && p.Y < max.Y
}

func order(maxim aoc.Vec2, dir aoc.Vec2) []aoc.Vec2 {
	ret := make([]aoc.Vec2, 0, maxim.X*maxim.Y)
	switch dir {
	case aoc.North:
		for y := 0; y < maxim.Y; y++ {
			for x := 0; x < maxim.X; x++ {
				ret = append(ret, aoc.Vec2{X: x, Y: y})
			}
		}
	case aoc.East:
		for x := maxim.X - 1; x >= 0; x-- {
			for y := 0; y < maxim.Y; y++ {
				ret = append(ret, aoc.Vec2{X: x, Y: y})
			}
		}
	case aoc.South:
		for y := maxim.Y - 1; y >= 0; y-- {
			for x := 0; x < maxim.X; x++ {
				ret = append(ret, aoc.Vec2{X: x, Y: y})
			}
		}
	case aoc.West:
		for x := 0; x < maxim.X; x++ {
			for y := 0; y < maxim.Y; y++ {
				ret = append(ret, aoc.Vec2{X: x, Y: y})
			}
		}
	default:
		panic("unknown direction")
	}
	return ret
}

func rollBoulders(s [][]byte, dir aoc.Vec2) {
	maxim := aoc.Vec2{X: len(s[0]), Y: len(s)}
	for _, pos := range order(maxim, dir) {
		c := s[pos.Y][pos.X]
		if c != 'O' {
			continue
		}
		tgt := pos
		for {
			nxt := tgt.Add(dir)
			if !inRange(nxt, maxim) || s[nxt.Y][nxt.X] != 0 {
				break
			}
			tgt = nxt
		}
		s[tgt.Y][tgt.X] = c
		if tgt != pos {
			s[pos.Y][pos.X] = 0
		}
	}
}

func score(m [][]byte) int {
	var sum int
	for y, l := range m {
		for _, r := range l {
			if r == 'O' {
				sum += len(m) - y
			}
		}
	}
	return sum
}

func runPartOne(s input) error {
	m := parseMap(s)
	rollBoulders(m, aoc.North)
	fmt.Println(score(m))
	return nil
}

var testString2 = testString1

func spin(m [][]byte) {
	rollBoulders(m, aoc.North)
	rollBoulders(m, aoc.West)
	rollBoulders(m, aoc.South)
	rollBoulders(m, aoc.East)
}

func toString(m [][]byte) string {
	var ret strings.Builder
	for _, row := range m {
		row = bytes.ReplaceAll(row, []byte{0}, []byte{'.'})
		_, _ = ret.Write(row)
		_, _ = ret.WriteRune('\n')
	}
	return ret.String()
}

func spinN(m [][]byte, n int) {
	var spins int
	seen := map[string]int{toString(m): 0}
	for ; spins < n; spins++ {
		spin(m)
		s := toString(m)
		if v, ok := seen[s]; ok {
			left := n - spins
			cycle := spins - v
			spins = n - (left % cycle)
		} else {
			seen[s] = spins
		}
	}
}

func runPartTwo(s input) error {
	m := parseMap(s)
	spinN(m, 1000000000)
	fmt.Println(score(m))
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
