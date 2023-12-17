package day17

import (
	"container/heap"
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

var testString1 = `2413432311323
3215453535623
3255245654254
3446585845452
4546657867536
1438598798454
4457876987766
3637877979653
4654967986887
4564679986453
1224686865563
2546548887735
4322674655533
`

type Crucible struct {
	Pos, Dir       aoc.Vec2
	WithoutTurning int
	Heat           int
	IsUltra        bool
}

func (c Crucible) Go(dir aoc.Vec2) Crucible {
	c.Pos = c.Pos.Add(dir)
	if dir == c.Dir {
		c.WithoutTurning++
	} else {
		c.WithoutTurning = 1
	}
	c.Dir = dir
	return c
}

func nextPaths(c Crucible, maxim aoc.Vec2) []Crucible {
	var ret []Crucible
	if (!c.IsUltra && c.WithoutTurning < 3) ||
		(c.IsUltra && c.WithoutTurning < 10) {
		ret = append(ret, c.Go(c.Dir))
	}
	if c.IsUltra && c.WithoutTurning < 4 {
		return ret
	}
	l := c.Go(c.Dir.TurnLeft())
	if l.Pos.InRange(maxim) {
		ret = append(ret, l)
	}
	r := c.Go(c.Dir.TurnRight())
	if r.Pos.InRange(maxim) {
		ret = append(ret, r)
	}
	return ret
}

type CrucibleQueue struct {
	Heat      map[Crucible]int
	Crucibles []Crucible
}

func (q *CrucibleQueue) Len() int { return len(q.Crucibles) }

func (q *CrucibleQueue) Less(i, j int) bool {
	return q.Crucibles[i].Heat < q.Crucibles[j].Heat
}

func (q *CrucibleQueue) Swap(i, j int) {
	q.Crucibles[i], q.Crucibles[j] = q.Crucibles[j], q.Crucibles[i]
}

func (q *CrucibleQueue) Push(x any) {
	q.Crucibles = append(q.Crucibles, x.(Crucible))
}

func (q *CrucibleQueue) Pop() any {
	old := q.Crucibles
	n := len(old)
	item := old[n-1]
	q.Crucibles = old[:n-1]
	return item
}

type searchKey struct {
	Pos, Dir       aoc.Vec2
	WithoutTurning int
}

func searchForPath(start []Crucible, heat [][]int) Crucible {
	maxim := aoc.Vec2{X: len(heat[0]), Y: len(heat)}
	tgt := aoc.Vec2{X: maxim.X - 1, Y: maxim.Y - 1}

	queue := CrucibleQueue{Crucibles: start}

	least := make(map[searchKey]int)
	heap.Init(&queue)

	for queue.Len() > 0 {
		i := heap.Pop(&queue)
		c := i.(Crucible)

		for _, p := range nextPaths(c, maxim) {
			if !p.Pos.InRange(maxim) {
				continue
			}
			p.Heat = c.Heat + heat[p.Pos.Y][p.Pos.X]
			if p.Pos == tgt {
				return p
			}
			k := searchKey{Pos: p.Pos, Dir: p.Dir, WithoutTurning: p.WithoutTurning}
			if l, ok := least[k]; ok && l <= p.Heat {
				continue
			}
			least[k] = p.Heat
			heap.Push(&queue, p)
		}
	}
	panic("no solution found")
}

func parseHeat(s input) ([][]int, error) {
	heat := make([][]int, len(s))
	for y, l := range s {
		heat[y] = make([]int, len(l))
		for x, r := range l {
			v, err := strconv.Atoi(string(r))
			if err != nil {
				return nil, err
			}
			heat[y][x] = v
		}
	}
	return heat, nil
}

func runPartOne(s input) error {
	heat, err := parseHeat(s)
	if err != nil {
		return err
	}
	var zero aoc.Vec2
	start := []Crucible{
		{Pos: zero, Dir: aoc.East, WithoutTurning: 1},  // Path: []aoc.Vec2{zero}},
		{Pos: zero, Dir: aoc.South, WithoutTurning: 1}, // Path: []aoc.Vec2{zero}},
	}
	c := searchForPath(start, heat)
	fmt.Println(c.Heat)
	return nil
}

var testString2 = testString1

func runPartTwo(s input) error {
	heat, err := parseHeat(s)
	if err != nil {
		return err
	}
	var zero aoc.Vec2
	start := []Crucible{
		{Pos: zero, Dir: aoc.East, WithoutTurning: 1, IsUltra: true},
		{Pos: zero, Dir: aoc.South, WithoutTurning: 1, IsUltra: true},
	}
	c := searchForPath(start, heat)
	fmt.Println(c.Heat)
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
