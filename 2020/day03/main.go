package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type vector struct {
	right, down int
}

type position struct {
	x, y int
}

func (p position) Apply(v vector) position {
	return position{
		x: p.x + v.right,
		y: p.y + v.down,
	}
}

type GridContent int

const (
	gridEmpty GridContent = 0
	gridTree  GridContent = 1
)

type mountain struct {
	Height, Width int
	Grid          [][]GridContent
}

func (m mountain) getAt(p position) GridContent {
	realX := p.x % m.Width
	realY := p.y % m.Height

	return m.Grid[realY][realX]
}

func traverseMountain(m mountain, start position, v vector) []GridContent {
	var bump []GridContent
	pos := start
	for pos.y < m.Height {
		bump = append(bump, m.getAt(pos))
		pos = pos.Apply(v)
	}
	return bump
}

func parseInput(s string) mountain {
	s = strings.TrimSpace(s)
	lines := strings.Split(s, "\n")
	if len(lines) == 0 {
		panic("invalid grid")
	}

	height := len(lines)
	width := len(lines[0])
	gridContent := make([][]GridContent, height)

	for y, row := range lines {
		if len(row) != width {
			panic("invalid grid")
		}
		gridRow := make([]GridContent, width)
		for x, c := range row {
			if c == '#' {
				gridRow[x] = gridTree
			}
		}
		gridContent[y] = gridRow
	}
	return mountain{
		Height: height,
		Width:  width,
		Grid:   gridContent,
	}
}

func countTrees(bumps []GridContent) int {
	var trees int
	for _, b := range bumps {
		if b == gridTree {
			trees++
		}
	}
	return trees
}

func runPartOne() error {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		return err
	}
	m := parseInput(string(content))
	bumps := traverseMountain(m, position{}, vector{right: 3, down: 1})
	fmt.Println(countTrees(bumps))
	return nil
}

func getSlopes() []vector {
	return []vector{
		{right: 1, down: 1},
		{right: 3, down: 1},
		{right: 5, down: 1},
		{right: 7, down: 1},
		{right: 1, down: 2},
	}
}

func runPartTwo() error {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		return err
	}
	m := parseInput(string(content))

	mult := 1
	for _, s := range getSlopes() {
		bumps := traverseMountain(m, position{}, s)
		mult *= countTrees(bumps)
	}

	fmt.Println(mult)
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
