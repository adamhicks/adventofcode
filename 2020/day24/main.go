package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type direction int

const (
	dirWest      direction = 1
	dirNorthWest direction = 2
	dirNorthEast direction = 3
	dirEast      direction = 4
	dirSouthEast direction = 5
	dirSouthWest direction = 6
)

func (d direction) String() string {
	switch d {
	case dirWest:
		return "w"
	case dirNorthWest:
		return "nw"
	case dirNorthEast:
		return "ne"
	case dirEast:
		return "e"
	case dirSouthEast:
		return "se"
	case dirSouthWest:
		return "sw"
	default:
		return "unknown!"
	}
}

type coord struct {
	x, y int
}

func (c coord) neighbours() []coord {
	return []coord{
		move(c, dirEast),
		move(c, dirNorthEast),
		move(c, dirNorthWest),
		move(c, dirWest),
		move(c, dirSouthWest),
		move(c, dirSouthEast),
	}
}

func move(c coord, dir direction) coord {
	switch dir {
	case dirWest:
		c.x -= 2
	case dirNorthWest:
		c.x--
		c.y--
	case dirNorthEast:
		c.x++
		c.y--
	case dirEast:
		c.x += 2
	case dirSouthEast:
		c.x++
		c.y++
	case dirSouthWest:
		c.x--
		c.y++
	default:
		panic("unknown direction")
	}
	return c
}

func parseInput(content string) [][]direction {
	var ret [][]direction
	for _, l := range strings.Split(content, "\n") {
		if l == "" {
			continue
		}
		var instructs []direction
		var s string
		for _, c := range l {
			s = s + string(c)
			if c == 's' || c == 'n' {
				continue
			}
			switch s {
			case "w":
				instructs = append(instructs, dirWest)
			case "nw":
				instructs = append(instructs, dirNorthWest)
			case "ne":
				instructs = append(instructs, dirNorthEast)
			case "e":
				instructs = append(instructs, dirEast)
			case "se":
				instructs = append(instructs, dirSouthEast)
			case "sw":
				instructs = append(instructs, dirSouthWest)
			default:
				panic("unknown string: " + s)
			}
			s = ""
		}
		ret = append(ret, instructs)
	}
	return ret
}

func findBlackTiles(instructs [][]direction) map[coord]bool {
	black := make(map[coord]bool)

	for _, ins := range instructs {
		var c coord
		for _, i := range ins {
			c = move(c, i)
		}
		black[c] = !black[c]
	}
	for k, v := range black {
		if !v {
			delete(black, k)
		}
	}
	return black
}

func runPartOne() error {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		return err
	}
	ins := parseInput(string(content))
	b := findBlackTiles(ins)
	fmt.Println(len(b))
	return nil
}

func tick(black map[coord]bool) map[coord]bool {
	visited := make(map[coord]bool, len(black))
	nextBlack := make(map[coord]bool, len(black))

	toVisit := make([]coord, 0, len(black))
	for k := range black {
		toVisit = append(toVisit, k)
	}

	for len(toVisit) > 0 {
		cur := toVisit[0]
		toVisit = toVisit[1:]

		var bt int
		for _, n := range cur.neighbours() {
			if black[n] {
				bt++
			}
			if black[cur] && !visited[n] {
				toVisit = append(toVisit, n)
				visited[n] = true
			}
		}

		if bt == 2 {
			nextBlack[cur] = true
		} else if bt == 1 && black[cur] {
			nextBlack[cur] = true
		}
	}

	return nextBlack
}

func runPartTwo() error {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		return err
	}
	ins := parseInput(string(content))
	s := findBlackTiles(ins)

	for i := 0; i < 100; i++ {
		s = tick(s)
	}
	fmt.Println(len(s))
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
