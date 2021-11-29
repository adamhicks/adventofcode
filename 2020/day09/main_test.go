package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testInput = `35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576
`

func TestExample(t *testing.T) {
	vals, err := parseInput(testInput)
	require.NoError(t, err)
	assert.Len(t, vals, 20)

	wrong, err := findFirstNoneSum(vals, 5)
	require.NoError(t, err)
	assert.Equal(t, 127, wrong)
}

func TestSumRange(t *testing.T) {
	vals, err := parseInput(testInput)
	require.NoError(t, err)
	assert.Len(t, vals, 20)

	r, err := findSumRange(vals, 127)
	require.NoError(t, err)

	assert.Equal(t, []int{15, 25, 47, 40}, r)
}

func TestSumMinMax(t *testing.T) {
	vals := []int{15, 25, 47, 40}
	assert.Equal(t, 62, sumMinMax(vals))
}
