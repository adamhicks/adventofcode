package day21

import (
	_ "embed"
	"fmt"
	"github.com/adamhicks/adventofcode/2023/aoc"
	"math"
	"strings"
)

//go:embed input.txt
var inputString string

type input []string

func parseInput(s string) input {
	return strings.Split(strings.TrimSpace(s), "\n")
}

var testString1 = `...........
.....###.#.
.###.##..#.
..#.#...#..
....#.#....
.##..S####.
.##..#...#.
.......##..
.##.#.####.
.##..##.##.
...........`

func findStart(s input) aoc.Vec2 {
	for y, l := range s {
		for x, r := range l {
			if r == 'S' {
				return aoc.Vec2{X: x, Y: y}
			}
		}
	}
	panic("no start found")
}

func boulders(s input) map[aoc.Vec2]bool {
	ret := make(map[aoc.Vec2]bool)
	for y, l := range s {
		for x, r := range l {
			if r == '#' {
				ret[aoc.Vec2{X: x, Y: y}] = true
			}
		}
	}
	return ret
}

type isBoulderFunc func(aoc.Vec2) bool

func singleMap(boulders map[aoc.Vec2]bool, maxim aoc.Vec2) isBoulderFunc {
	return func(pos aoc.Vec2) bool {
		return !pos.InRange(maxim) || boulders[pos]
	}
}

func fillMap(bf isBoulderFunc, start map[aoc.Vec2]int, endStep int) map[aoc.Vec2]int {
	minStep := math.MaxInt
	incoming := make(map[int][]aoc.Vec2)
	for pos, step := range start {
		minStep = min(minStep, step)
		incoming[step] = append(incoming[step], pos)
	}
	wave := incoming[minStep]
	seen := make(map[aoc.Vec2]int)
	for _, p := range wave {
		seen[p] = minStep
	}

	for i := minStep + 1; i <= endStep && len(wave) > 0; i++ {
		var nxt []aoc.Vec2
		for _, p := range incoming[i] {
			if v, ok := seen[p]; ok && v < i {
				continue
			}
			seen[p] = i
			nxt = append(nxt, p)
		}

		for _, p := range wave {
			for _, n := range p.Orthogonal() {
				if bf(n) {
					continue
				}
				if _, ok := seen[n]; ok {
					continue
				}
				seen[n] = i
				nxt = append(nxt, n)
			}
		}
		wave = nxt
	}
	return seen
}

func countMod2(m map[aoc.Vec2]int, val int) int {
	var count int
	for _, i := range m {
		if i%2 == val%2 {
			count++
		}
	}
	return count
}

func runPartOne(s input) error {
	maxim := aoc.Vec2{X: len(s[0]), Y: len(s)}
	bf := singleMap(boulders(s), maxim)

	steps := 64
	positions := fillMap(bf, map[aoc.Vec2]int{findStart(s): 0}, steps)
	fmt.Println(countMod2(positions, steps))
	return nil
}

var testString2 = `...........
......##.#.
.###..#..#.
..#.#...#..
....#.#....
.....S.....
.##......#.
.......##..
.##.#.####.
.##...#.##.
...........`

func sq(x int) int {
	return x * x
}

func sumMiddle(posCount, n int) int {
	n++
	b := sq(n - ((n + 1) % 2))
	return posCount * b
}

func sumAlt(posCount, n int) int {
	n++
	b := sq(n - (n % 2))
	return posCount * b
}

func runPartTwo(s input) error {
	maxim := aoc.Vec2{X: len(s[0]), Y: len(s)}

	steps := 26501365
	blocks := (steps / maxim.X) - 1

	grid := fillMap(
		singleMap(boulders(s), maxim),
		map[aoc.Vec2]int{findStart(s): 0},
		math.MaxInt,
	)
	// middle blocks will match the step modulo
	middle := countMod2(grid, steps)
	// alternate blocks will be offset from the step modulo
	alt := countMod2(grid, steps+1)

	var sum int
	sum += sumMiddle(middle, blocks)
	sum += sumAlt(alt, blocks)

	cover := blocks * maxim.X

	// complete grids can leave one or two partial grids on each diagonal
	extraEdge := steps - cover
	extraEdgeBlocks := extraEdge / maxim.X
	if extraEdge%maxim.X > 0 {
		extraEdgeBlocks++
	}
	for _, corner := range []aoc.Vec2{
		{X: 0, Y: 0},
		{X: 0, Y: maxim.Y - 1},
		{X: maxim.X - 1, Y: 0},
		{X: maxim.X - 1, Y: maxim.Y - 1},
	} {
		for extra := 0; extra < extraEdgeBlocks; extra++ {
			start := cover + extra*maxim.X
			cornerFill := fillMap(
				singleMap(boulders(s), maxim),
				map[aoc.Vec2]int{corner: start + 1},
				steps,
			)
			cornerCount := countMod2(cornerFill, steps)
			sum += (blocks + extra) * cornerCount
		}
	}
	// the north, south, east, and west tips can also make partial grids
	tipStep := cover + maxim.X/2
	extraTips := steps - tipStep
	extraTipBlocks := extraTips / maxim.X
	if extraTips%maxim.X > 0 {
		extraTipBlocks++
	}
	for _, tip := range []aoc.Vec2{
		{X: 0, Y: maxim.Y / 2},           // middle left
		{X: maxim.X - 1, Y: maxim.Y / 2}, // middle right
		{X: maxim.X / 2, Y: 0},           // middle top
		{X: maxim.X / 2, Y: maxim.Y - 1}, // middle bottom
	} {
		for extra := 0; extra < extraTipBlocks; extra++ {
			start := tipStep + extra*maxim.X
			tipFill := fillMap(
				singleMap(boulders(s), maxim),
				map[aoc.Vec2]int{tip: start + 1},
				steps,
			)
			tipCount := countMod2(tipFill, steps)
			sum += tipCount
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
