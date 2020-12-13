package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testInput = `939
7,13,x,x,59,x,31,19`

func TestParse(t *testing.T) {
	i, l, err := parseInput(testInput)
	require.NoError(t, err)

	assert.Equal(t, 939, i)
	assert.Equal(t, []int{7, 13, 0, 0, 59, 0, 31, 19}, l)
}

func TestExample(t *testing.T) {
	i, l, err := parseInput(testInput)
	require.NoError(t, err)

	min, bus := getNextBus(i, l)
	assert.Equal(t, 5, min)
	assert.Equal(t, 59, bus)
}
