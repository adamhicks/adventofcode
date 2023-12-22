package day22

import (
	_ "embed"
	"fmt"
	"github.com/adamhicks/adventofcode/2023/aoc"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var inputString string

type input []string

func parseInput(s string) input {
	return strings.Split(strings.TrimSpace(s), "\n")
}

var testString1 = `1,0,1~1,2,1
0,0,2~2,0,2
0,2,3~2,2,3
0,0,4~0,2,4
2,0,5~2,2,5
0,1,6~2,1,6
1,1,8~1,1,9`

var ada = `1,0,1~1,2,1
0,0,2~2,0,2
0,2,3~2,2,3
0,0,4~0,2,4
2,0,5~2,2,5
0,1,6~2,1,6
1,1,8~1,1,9`

type Cube struct {
	aoc.Rect
	Z, Height int
}

func (c Cube) Top() int {
	return c.Z + c.Height
}

func overlaps(above, below Cube) bool {
	return !above.Overlap(below.Rect).IsZero()
}

func parseCoOrd(s string) [3]int {
	var ret [3]int
	for i, p := range strings.SplitN(s, ",", 3) {
		ret[i] = aoc.Must(strconv.Atoi(p))
	}
	return ret
}

func cube(from, to [3]int) Cube {
	return Cube{
		Rect: aoc.Rect{
			From: aoc.Vec2{X: from[0], Y: from[1]},
			To:   aoc.Vec2{X: to[0] + 1, Y: to[1] + 1},
		},
		Z: from[2], Height: to[2] - from[2] + 1,
	}
}

func parseCubes(s input) []Cube {
	var ret []Cube
	for _, l := range s {
		fromStr, toStr, _ := strings.Cut(l, "~")
		ret = append(ret, cube(parseCoOrd(fromStr), parseCoOrd(toStr)))
	}
	return ret
}

func sortCubesByBottom(c []Cube) {
	sort.Slice(c, func(i, j int) bool {
		if c[i].Z != c[j].Z {
			return c[i].Z < c[j].Z
		}
		return c[i].Top() > c[i].Top()
	})
}

func sortCubesByTop(c []Cube) {
	sort.Slice(c, func(i, j int) bool {
		ti, tj := c[i].Top(), c[j].Top()
		return ti < tj
	})
}

func dropCube(dropping Cube, cubes []Cube) int {
	for i := len(cubes) - 1; i >= 0; i-- {
		c := cubes[i]
		if overlaps(dropping, c) {
			return c.Top()
		}
	}
	return 1
}

func dependGraph(cubes []Cube) ([][]int, [][]int) {
	carries := make([][]int, len(cubes))
	depends := make([][]int, len(cubes))
	for i := 0; i < len(cubes)-1; i++ {
		c := cubes[i]
		var carry []int
		for j := i + 1; j < len(cubes); j++ {
			b := cubes[j]
			if c.Top() < b.Z {
				break
			}
			if c.Z == b.Z {
				continue
			}
			if overlaps(c, b) {
				carry = append(carry, j)
				depends[j] = append(depends[j], i)
			}
		}
		carries[i] = carry
	}
	return carries, depends
}

func canRemove(i int, carries [][]int, dependsOn [][]int) bool {
	for _, c := range carries[i] {
		if len(dependsOn[c]) == 1 {
			return false
		}
	}
	return true
}

func runPartOne(s input) error {
	starting := parseCubes(s)
	sortCubesByBottom(starting)
	cubes := make([]Cube, 0, len(starting))

	for _, c := range starting {
		c.Z = dropCube(c, cubes)
		cubes = append(cubes, c)
		sortCubesByTop(cubes)
	}
	sortCubesByBottom(cubes)

	carries, depends := dependGraph(cubes)

	var count int
	for i := range carries {
		if canRemove(i, carries, depends) {
			count++
		}
	}
	fmt.Println(count)

	return nil
}

var testString2 = testString1

func falls(i int, dependsOn [][]int, gone map[int]bool) bool {
	for _, d := range dependsOn[i] {
		if !gone[d] {
			return false
		}
	}
	return true
}

func countFall(i int, carries, dependsOn [][]int) int {
	gone := map[int]bool{i: true}
	q := make([]int, len(carries[i]))
	copy(q, carries[i])
	for len(q) > 0 {
		n := q[0]
		q = q[1:]
		if falls(n, dependsOn, gone) {
			gone[n] = true
			for _, c := range carries[n] {
				q = append(q, c)
			}
		}
	}
	return len(gone) - 1
}

func runPartTwo(s input) error {
	starting := parseCubes(s)
	sortCubesByBottom(starting)
	cubes := make([]Cube, 0, len(starting))

	for _, c := range starting {
		c.Z = dropCube(c, cubes)
		cubes = append(cubes, c)
		sortCubesByTop(cubes)
	}

	sortCubesByBottom(cubes)

	carries, depends := dependGraph(cubes)

	var sum int
	for i := range cubes {
		sum += countFall(i, carries, depends)
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
