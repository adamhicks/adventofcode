package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const testInput = `
..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#
`

func TestParseInput(t *testing.T) {
	m := parseInput(testInput)
	assert.Equal(t, 11, m.Height)
	assert.Equal(t, 11, m.Width)
	assert.Equal(t, gridTree, m.getAt(position{x: 2, y: 0}))
	assert.Equal(t, gridEmpty, m.getAt(position{x: 0, y: 0}))
	assert.Equal(t, gridTree, m.getAt(position{x: 0, y: 12}))
	assert.Equal(t, gridTree, m.getAt(position{x: 15, y: 1}))
	assert.Equal(t, gridTree, m.getAt(position{x: 26, y: 12}))
}

func TestTraverse(t *testing.T) {
	m := parseInput(testInput)
	bump := traverseMountain(m, position{}, vector{right: 3, down: 1})

	assert.Equal(t, 7, countTrees(bump))
}

func TestPartTwo(t *testing.T) {
	m := parseInput(testInput)

	mult := 1
	for _, s := range getSlopes() {
		bump := traverseMountain(m, position{}, s)
		mult *= countTrees(bump)
	}
	assert.Equal(t, 336, mult)
}
