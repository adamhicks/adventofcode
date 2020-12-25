package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testInput = `5764801
17807724`

func TestExample(t *testing.T) {
	card, door, err := parseInput(testInput)
	require.NoError(t, err)

	cSecret := findMultiple(7, 20201227, card)
	dSecret := findMultiple(7, 20201227, door)

	assert.Equal(t, 8, cSecret)
	assert.Equal(t, 11, dSecret)
}

func TestLoopN(t *testing.T) {
	card := loopN(7, 20201227, 8)
	door := loopN(7, 20201227, 11)

	assert.Equal(t, 5764801, card)
	assert.Equal(t, 17807724, door)
}
