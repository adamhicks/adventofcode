package main

import (
	"strconv"
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

func TestEGCD(t *testing.T) {
	a, b, _ := eGCD(29, 17)
	assert.Equal(t, -7, a)
	assert.Equal(t, 12, b)
}

func TestBrilliantCRT(t *testing.T) {
	// Example from https://brilliant.org/wiki/chinese-remainder-theorem/
	mnrs := []modAndRemainder{
		{Remain: 6, Mod: 7},
		{Remain: 4, Mod: 5},
		{Remain: 1, Mod: 3},
	}
	mr := chineseRemainder(mnrs)
	assert.Equal(t, 34, mr.Remain)
	assert.Equal(t, 105, mr.Mod)
}

func TestExample2(t *testing.T) {
	testCases := []struct{
		mnrs []modAndRemainder
		exp modAndRemainder
	}{
		{mnrs: []modAndRemainder{
			{Mod: 17, Remain: 0},
			{Mod: 13, Remain: 2},
			{Mod: 19, Remain: 3},
		},
			exp: modAndRemainder{Mod: 4199, Remain: 782},
		},
		{mnrs: []modAndRemainder{
			{Mod: 1789, Remain: 0},
			{Mod: 37, Remain: 1},
			{Mod: 47, Remain: 2},
			{Mod: 1889, Remain: 3},
		},
			exp: modAndRemainder{Mod: 5876813119, Remain: -1202161486},
		},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			mr := chineseRemainder(tc.mnrs)
			assert.Equal(t, tc.exp, mr)
		})
	}
}