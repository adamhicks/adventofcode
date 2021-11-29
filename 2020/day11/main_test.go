package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const testInput = `L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`

func TestExample(t *testing.T) {
	grid := parseInput(testInput)
	grid = tickUntilStale(grid, getNextSeatState)
	assert.Equal(t, 37, countOccupied(grid))
}

func TestExample2(t *testing.T) {
	grid := parseInput(testInput)
	grid = tickUntilStale(grid, getNextSeatState2)
	assert.Equal(t, 26, countOccupied(grid))
}
