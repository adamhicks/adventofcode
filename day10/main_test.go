package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	smallExample = `16
10
15
5
1
11
7
19
6
12
4
`
	largerExample = `28
33
18
42
31
14
46
20
48
47
24
23
49
45
19
38
39
11
1
32
25
35
8
17
7
9
4
2
34
10
3
`
)

func TestExample(t *testing.T) {
	vals, err := parseInput(smallExample)
	require.NoError(t, err)

	ones, threes := countOneAndThreeDiffs(vals)
	assert.Equal(t, 7, ones)
	assert.Equal(t, 5, threes)
}

func TestExample2(t *testing.T) {
	vals, err := parseInput(largerExample)
	require.NoError(t, err)

	ones, threes := countOneAndThreeDiffs(vals)
	assert.Equal(t, 22, ones)
	assert.Equal(t, 10, threes)
}

func TestCountPermutations(t *testing.T) {
	vals, err := parseInput(smallExample)
	require.NoError(t, err)

	i := countPermutations(vals)
	assert.Equal(t, int64(8), i)
}

func TestCountPermutations2(t *testing.T) {
	vals, err := parseInput(largerExample)
	require.NoError(t, err)

	i := countPermutations(vals)
	assert.Equal(t, int64(19208), i)
}
