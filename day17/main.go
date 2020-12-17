package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type coord struct {
	X, Y, Z, W int
}

type universe map[coord]bool

func (u universe) String() string {

	var min, max coord
	for c := range u {
		if c.X < min.X {
			min.X = c.X
		}
		if c.Y < min.Y {
			min.Y = c.Y
		}
		if c.Z < min.Z {
			min.Z = c.Z
		}
		if c.W < min.W {
			min.W = c.W
		}
		if c.X > max.X {
			max.X = c.X
		}
		if c.Y > max.Y {
			max.Y = c.Y
		}
		if c.Z > max.Z {
			max.Z = c.Z
		}
		if c.W > max.W {
			max.W = c.W
		}
	}

	var b strings.Builder

	for w := min.W; w <= max.W; w++ {
		for z := min.Z; z <= max.Z; z++ {
			fmt.Fprintf(&b, "z = %d, w = %d\n", z, w)
			for y := min.Y; y <= max.Y; y++ {
				for x := min.X; x <= max.X; x++ {
					if u[coord{X: x, Y: y, Z: z, W: w}] {
						fmt.Fprint(&b, "#")
					} else {
						fmt.Fprint(&b, ".")
					}
				}
				fmt.Fprint(&b, "\n")
			}
			fmt.Fprint(&b, "\n")
		}
	}

	return b.String()
}

func getNeighbours3d(c coord) []coord {
	var neighbours []coord
	for z := -1; z <= 1; z++ {
		for y := -1; y <= 1; y++ {
			for x := -1; x <= 1; x++ {
				if x == 0 && y == 0 && z == 0 {
					continue
				}
				neighbours = append(neighbours, coord{
					X: c.X + x,
					Y: c.Y + y,
					Z: c.Z + z,
				})
			}
		}
	}
	return neighbours
}

type neighbourFunc func(coord) []coord

func tickGameOfLife(prev universe, neighFun neighbourFunc) universe {
	nextUni := make(universe)
	visited := make(universe)
	toConsider := make([]coord, 0, len(prev))
	for k := range prev {
		toConsider = append(toConsider, k)
	}

	for len(toConsider) > 0 {
		cur := toConsider[0]
		toConsider = toConsider[1:]

		curActive := prev[cur]
		var activeNeighbours int
		for _, n := range neighFun(cur) {
			if prev[n] {
				activeNeighbours++
			}
			// Don't need to explore this nodes neighbours
			if !curActive {
				continue
			}
			if _, ok := visited[n]; ok {
				continue
			}
			toConsider = append(toConsider, n)
			visited[n] = true
		}

		var active bool
		if curActive {
			active = activeNeighbours == 2 || activeNeighbours == 3
		} else {
			active = activeNeighbours == 3
		}
		if active {
			nextUni[cur] = true
		}
	}
	return nextUni
}

func parseInput(content string) universe {
	u := make(universe)
	for y, l := range strings.Split(content, "\n") {
		for x, c := range l {
			if c == '#' {
				u[coord{X: x, Y: y, Z: 0}] = true
			}
		}
	}
	return u
}

func runPartOne() error {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		return err
	}
	uni := parseInput(string(content))
	for i := 0; i < 6; i++ {
		uni = tickGameOfLife(uni, getNeighbours3d)
	}
	fmt.Println(len(uni))
	return nil
}

func getNeighbours4d(c coord) []coord {
	var neighbours []coord
	for w := -1; w <= 1; w++ {
		for z := -1; z <= 1; z++ {
			for y := -1; y <= 1; y++ {
				for x := -1; x <= 1; x++ {
					if x == 0 && y == 0 && z == 0 && w == 0 {
						continue
					}
					neighbours = append(neighbours, coord{
						X: c.X + x,
						Y: c.Y + y,
						Z: c.Z + z,
						W: c.W + w,
					})
				}
			}
		}
	}
	return neighbours
}

func runPartTwo() error {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		return err
	}
	uni := parseInput(string(content))
	for i := 0; i < 6; i++ {
		uni = tickGameOfLife(uni, getNeighbours4d)
	}
	fmt.Println(len(uni))
	return nil
}

func main() {
	if err := runPartOne(); err != nil {
		log.Fatal(err)
	}
	if err := runPartTwo(); err != nil {
		log.Fatal(err)
	}
}
