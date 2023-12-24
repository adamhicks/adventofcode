package day23

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

var testString1 = `#.#####################
#.......#########...###
#######.#########.#.###
###.....#.>.>.###.#.###
###v#####.#v#.###.#.###
###.>...#.#.#.....#...#
###v###.#.#.#########.#
###...#.#.#.......#...#
#####.#.#.#######.#.###
#.....#.#.#.......#...#
#.#####.#.#.#########v#
#.#...#...#...###...>.#
#.#.#v#######v###.###v#
#...#.>.#...>.>.#.###.#
#####v#.#.###v#.#.###.#
#.....#...#...#.#.#...#
#.#########.###.#.#.###
#...###...#...#...#.###
###.###.#.###v#####v###
#...#...#.#.>.>.#.>.###
#.###.###.#.###.#.#v###
#.....###...###...#...#
#####################.#`

func checkDir(from, to aoc.Vec2, req aoc.Vec2) bool {
	return to.Sub(from) == req
}

type path struct {
	pos   aoc.Vec2
	seen  map[aoc.Vec2]bool
	total int
}

type traverseFunc func(from, to aoc.Vec2, c byte) bool

func canTraverse(from, to aoc.Vec2, c byte) bool {
	switch c {
	case '>':
		return checkDir(from, to, aoc.East)
	case '<':
		return checkDir(from, to, aoc.West)
	case 'v':
		return checkDir(from, to, aoc.South)
	case '^':
		return checkDir(from, to, aoc.North)
	}
	return true
}

func allPaths(graph map[aoc.Vec2]map[aoc.Vec2]int, start, end aoc.Vec2) []path {
	var done []path
	q := []path{{
		pos:  start,
		seen: map[aoc.Vec2]bool{start: true},
	}}
	for len(q) > 0 {
		n := q[0]
		q = q[1:]

		if n.pos == end {
			done = append(done, n)
			continue
		}
		for to, val := range graph[n.pos] {
			if n.seen[to] {
				continue
			}
			tot := n.total + val + 1
			nxtPath := path{pos: to, total: tot}
			nxtPath.seen = make(map[aoc.Vec2]bool, len(n.seen)+1)
			for p := range n.seen {
				nxtPath.seen[p] = true
			}
			nxtPath.seen[to] = true
			q = append(q, nxtPath)
		}
	}
	return done
}

func startEnd(s input) (aoc.Vec2, aoc.Vec2) {
	start := aoc.Vec2{
		X: strings.Index(s[0], "."), Y: 0,
	}
	end := aoc.Vec2{
		X: strings.Index(s[len(s)-1], "."), Y: len(s) - 1,
	}
	return start, end
}

func runPartOne(s input) error {
	start, end := startEnd(s)
	graph := buildGraph(s, start, end, canTraverse)
	paths := allPaths(graph, start, end)
	var most int
	for _, p := range paths {
		most = max(most, p.total)
	}
	fmt.Println(most)
	return nil
}

var testString2 = testString1

func canGoUphill(_, _ aoc.Vec2, _ byte) bool {
	return true
}

func followUntilJunction(pos, dir aoc.Vec2, s input, canGo traverseFunc, end aoc.Vec2) ([]aoc.Vec2, []aoc.Vec2) {
	maxim := aoc.Vec2{X: len(s[0]), Y: len(s)}
	var p []aoc.Vec2

	pos, prev := pos.Add(dir), pos

	for c := 0; ; c++ {
		if pos == end {
			return append(p, pos), nil
		}
		var nxt []aoc.Vec2
		for _, n := range pos.Orthogonal() {
			if !n.InRange(maxim) || n == prev {
				continue
			}
			if s[n.Y][n.X] != '#' && canGo(pos, n, s[pos.Y][pos.X]) {
				nxt = append(nxt, n)
			}
		}
		p = append(p, pos)
		switch len(nxt) {
		case 1:
			pos, prev = nxt[0], pos
		case 0:
			return nil, nil
		default:
			return p, nxt
		}
	}
}

func buildGraph(s input, start, end aoc.Vec2, tf traverseFunc) map[aoc.Vec2]map[aoc.Vec2]int {
	type toExplore struct {
		Pos, Dir aoc.Vec2
	}
	graph := map[aoc.Vec2]map[aoc.Vec2]int{
		start: make(map[aoc.Vec2]int),
	}
	nodes := []toExplore{{Pos: start, Dir: aoc.South}}
	for len(nodes) > 0 {
		st := nodes[0]
		nodes = nodes[1:]

		p, next := followUntilJunction(st.Pos, st.Dir, s, tf, end)
		if len(p) == 0 {
			continue
		}
		to := p[len(p)-1]
		graph[st.Pos][to] = len(p) - 1
		if to == end {
			continue
		}
		if _, done := graph[to]; done {
			continue
		}
		graph[to] = make(map[aoc.Vec2]int)
		for _, n := range next {
			nodes = append(nodes, toExplore{Pos: to, Dir: n.Sub(to)})
		}
	}
	return graph
}

func runPartTwo(s input) error {
	start, end := startEnd(s)

	graph := buildGraph(s, start, end, canGoUphill)

	for from, edges := range graph {
		for to, val := range edges {
			if to != end {
				graph[to][from] = val
			}
		}
	}
	var high int
	for _, p := range allPaths(graph, start, end) {
		high = max(high, p.total)
	}
	fmt.Println(high)
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
