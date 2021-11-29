package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func getTestInput() []int {
	testInput := []int{
		1721,
		979,
		366,
		299,
		675,
		1456,
	}
	sort.Ints(testInput)
	return testInput
}

func TestFindPairSum(t *testing.T) {
	fac, err := findPairSum(getTestInput(), 2020)
	require.NoError(t, err)

	var sum int
	for _, i := range fac {
		sum += i
	}
	assert.Equal(t, sum, 2020)
}

func TestFindTreble(t *testing.T) {
	fac, err := findTrebleSum(getTestInput(), 2020)
	require.NoError(t, err)

	var sum int
	for _, i := range fac {
		sum += i
	}
	assert.Equal(t, 2020, sum)
}

func TestPartOne(t *testing.T) {
	nums, err := fetchInput()
	require.NoError(t, err)

	fac, err := findPairSum(nums, 2020)
	require.NoError(t, err)

	var sum int
	for _, i := range fac {
		sum += i
	}
	assert.Equal(t, 2020, sum)
}

func TestPartTwo(t *testing.T) {
	nums, err := fetchInput()
	require.NoError(t, err)

	fac, err := findTrebleSum(nums, 2020)
	require.NoError(t, err)

	var sum int
	for _, i := range fac {
		sum += i
	}
	assert.Equal(t, 2020, sum)
}
