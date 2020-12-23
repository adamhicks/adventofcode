package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateList(t *testing.T) {
	vals := []int{1, 4, 6, 2, 7, 9}
	l := createList(vals)
	assert.Equal(t, vals, l.ToArray())
}

const testInput = `389125467`

func TestMove(t *testing.T) {
	vals, err := parseInput(testInput)
	require.NoError(t, err)

	l := createList(vals)
	for i := 0; i < 10; i++ {
		l.move()
	}

	exp := []int{8, 3, 7, 4, 1, 9, 2, 6, 5}
	assert.Equal(t, exp, l.ToArray())
}

func TestGetOutput(t *testing.T) {
	l := createList([]int{3, 8, 9, 1, 2, 5, 4, 6, 7})
	for i := 0; i < 10; i++ {
		l.move()
	}

	assert.Equal(t, "92658374", getPartOneOutput(l))

	for i := 0; i < 90; i++ {
		l.move()
	}

	assert.Equal(t, "67384529", getPartOneOutput(l))
}

func TestFillList(t *testing.T) {
	l := createList([]int{1})
	l.fillTo(10)
	assert.Equal(t,
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		l.ToArray(),
	)
}

func TestPartTwo(t *testing.T) {
	vals, err := parseInput(testInput)
	require.NoError(t, err)

	l := createList(vals)
	l.fillTo(1000000)

	for i := 0; i < 10000000; i++ {
		l.move()
	}

	i := getPartTwoOutput(l)
	assert.Equal(t, 149245887792, i)
}
