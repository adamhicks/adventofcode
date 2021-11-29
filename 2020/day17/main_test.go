package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const testInput = `.#.
..#
###`

func TestParse(t *testing.T) {
	u := parseInput(testInput)
	exp := universe{
		{X: 1}:       true,
		{Y: 1, X: 2}: true,
		{Y: 2, X: 0}: true, {Y: 2, X: 1}: true, {Y: 2, X: 2}: true,
	}
	assert.Equal(t, exp, u)
}
func TestTick(t *testing.T) {
	u := parseInput(testInput)
	for i := 0; i < 6; i++ {
		u = tickGameOfLife(u, getNeighbours3d)
	}
	assert.Len(t, u, 112)
}

func TestTick4d(t *testing.T) {
	u := parseInput(testInput)
	for i := 0; i < 6; i++ {
		u = tickGameOfLife(u, getNeighbours4d)
	}
	assert.Len(t, u, 848)
}
