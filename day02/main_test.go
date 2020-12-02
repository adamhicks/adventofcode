package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testInput = `1-3 a: abcde
1-3 b: cdefg
2-9 c: ccccccccc
`

func TestReadAll(t *testing.T) {
	l, err := ReadAll([]byte(testInput))
	require.NoError(t, err)

	exp := []InputLine{
		{1, 3, 'a', "abcde"},
		{1, 3, 'b', "cdefg"},
		{2, 9, 'c', "ccccccccc"},
	}
	assert.Equal(t, exp, l)
}

func TestPartOneValid(t *testing.T) {
	testCases := []struct {
		name     string
		in       InputLine
		expValid bool
	}{
		{name: "1", in: InputLine{1, 3, 'a', "abcde"}, expValid: true},
		{name: "2", in: InputLine{1, 3, 'b', "cdefg"}, expValid: false},
		{name: "3", in: InputLine{2, 9, 'c', "ccccccccc"}, expValid: true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expValid, PartOneValid(tc.in))
		})
	}
}

func TestPartOne(t *testing.T) {
	ins, err := ReadAll([]byte(testInput))
	require.NoError(t, err)

	var sum int
	for _, in := range ins {
		if PartOneValid(in) {
			sum++
		}
	}
	assert.Equal(t, 2, sum)
}

func TestPartTwoValid(t *testing.T) {
	testCases := []struct {
		name     string
		in       InputLine
		expValid bool
	}{
		{name: "1", in: InputLine{1, 3, 'a', "abcde"}, expValid: true},
		{name: "2", in: InputLine{1, 3, 'b', "cdefg"}, expValid: false},
		{name: "3", in: InputLine{2, 9, 'c', "ccccccccc"}, expValid: false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expValid, PartTwoValid(tc.in))
		})
	}
}
