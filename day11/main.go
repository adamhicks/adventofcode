package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type tileType int

const (
	tileFloor    = 0
	tileSeat     = 1
	tileOccupied = 2
)

func parseInput(content string) [][]tileType {
	lines := strings.Split(content, "\n")
	ret := make([][]tileType, 0, len(lines))
	for _, l := range lines {
		if l == "" {
			continue
		}
		row := make([]tileType, len(l))
		for idx, c := range l {
			if c == 'L' {
				row[idx] = tileSeat
			}
		}
		ret = append(ret, row)
	}
	return ret
}

func listNeighbours(grid [][]tileType, row, col int) []tileType {
	var ret []tileType

	for y := -1; y <= 1; y++ {
		for x := -1; x <= 1; x++ {
			if x == 0 && y == 0 {
				continue
			}
			tRow := row + y
			tCol := col + x
			if tRow < 0 || tRow >= len(grid) {
				continue
			}
			if tCol < 0 || tCol >= len(grid[tRow]) {
				continue
			}
			ret = append(ret, grid[tRow][tCol])
		}
	}

	return ret
}

func getNextSeatState(prev [][]tileType, row, col int) tileType {
	s := prev[row][col]
	if s == tileFloor {
		return tileFloor
	}
	neighbours := listNeighbours(prev, row, col)
	var occu int
	for _, n := range neighbours {
		if n == tileOccupied {
			occu++
		}
	}
	if s == tileSeat && occu == 0 {
		return tileOccupied
	} else if s == tileOccupied && occu >= 4 {
		return tileSeat
	}
	return s
}

type getNext func(prev [][]tileType, row, col int) tileType

func tickGrid(grid [][]tileType, gn getNext) ([][]tileType, int) {
	nextGrid := make([][]tileType, 0, len(grid))
	for _, row := range grid {
		newRow := make([]tileType, len(row))
		nextGrid = append(nextGrid, newRow)
	}
	var changed int
	for y, row := range grid {
		for x := range row {
			s := gn(grid, y, x)
			nextGrid[y][x] = s
			if s != grid[y][x] {
				changed++
			}
		}
	}
	return nextGrid, changed
}

func tickUntilStale(grid [][]tileType, gn getNext) [][]tileType {
	for {
		nextGrid, changed := tickGrid(grid, gn)
		if changed == 0 {
			return grid
		}
		grid = nextGrid
	}
}

func countOccupied(grid [][]tileType) int {
	var i int
	for _, row := range grid {
		for _, s := range row {
			if s == tileOccupied {
				i++
			}
		}
	}
	return i
}

func printGrid(grid [][]tileType) {
	for _, row := range grid {
		for _, s := range row {
			r := "."
			switch s {
			case tileSeat:
				r = "L"
			case tileOccupied:
				r = "#"
			}
			fmt.Printf(r)
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func runPartOne() error {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		return err
	}
	grid := parseInput(string(content))
	grid = tickUntilStale(grid, getNextSeatState)
	fmt.Println(countOccupied(grid))
	return nil
}

func getUntilSeat(grid [][]tileType, row, col int, dX, dY int) tileType {
	x, y := col, row
	for {
		x += dX
		y += dY
		if y < 0 || y >= len(grid) {
			return tileFloor
		}
		if x < 0 || x >= len(grid[y]) {
			return tileFloor
		}
		if grid[y][x] != tileFloor {
			return grid[y][x]
		}
	}
}

func listLOSNeighbours(grid [][]tileType, row, col int) []tileType {
	var ret []tileType

	for y := -1; y <= 1; y++ {
		for x := -1; x <= 1; x++ {
			if x == 0 && y == 0 {
				continue
			}
			s := getUntilSeat(grid, row, col, x, y)
			ret = append(ret, s)
		}
	}

	return ret
}

func getNextSeatState2(prev [][]tileType, row, col int) tileType {
	s := prev[row][col]
	if s == tileFloor {
		return tileFloor
	}
	neighbours := listLOSNeighbours(prev, row, col)
	var occu int
	for _, n := range neighbours {
		if n == tileOccupied {
			occu++
		}
	}
	if s == tileSeat && occu == 0 {
		return tileOccupied
	} else if s == tileOccupied && occu >= 5 {
		return tileSeat
	}
	return s
}

func runPartTwo() error {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		return err
	}
	grid := parseInput(string(content))
	grid = tickUntilStale(grid, getNextSeatState2)
	fmt.Println(countOccupied(grid))
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
