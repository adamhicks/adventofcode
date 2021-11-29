package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testInput = `Player 1:
9
2
6
3
1

Player 2:
5
8
4
7
10`

func TestParse(t *testing.T) {
	decks, err := parseInput(testInput)
	require.NoError(t, err)

	assert.Equal(t, [][]int{
		{9, 2, 6, 3, 1},
		{5, 8, 4, 7, 10},
	}, decks)
}

func TestPlayGame(t *testing.T) {
	decks, err := parseInput(testInput)
	require.NoError(t, err)

	require.Len(t, decks, 2)
	res := playGame(decks[0], decks[1])
	s := score(res)

	assert.Equal(t, 306, s)
}

func TestPlayRecursiveGame(t *testing.T) {
	decks, err := parseInput(testInput)
	require.NoError(t, err)

	winner, res := playRecursiveGame(decks[0], decks[1])
	assert.Equal(t, 2, winner)
	assert.Equal(t, []int{7, 5, 6, 2, 4, 1, 10, 8, 9, 3}, res)
}

const infiniteGame = `Player 1:
43
19

Player 2:
2
29
14`

func TestInfiniteBreak(t *testing.T) {
	decks, err := parseInput(infiniteGame)
	require.NoError(t, err)

	winner, res := playRecursiveGame(decks[0], decks[1])
	assert.Equal(t, 1, winner)
	assert.Nil(t, res)
}
