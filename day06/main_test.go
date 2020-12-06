package main

import (
	"testing"

	"github.com/golang-collections/collections/set"
	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	var input = `abcx
abcy
abcz
`
	res := parseInput(input)
	assert.Equal(t, res, []group{{
		set.New('a', 'b', 'c', 'x'),
		set.New('a', 'b', 'c', 'y'),
		set.New('a', 'b', 'c', 'z'),
	}})
}

func TestSeparateGroups(t *testing.T) {
	var input = `abc

a
b
c

ab
ac

a
a
a
a

b
`
	res := parseInput(input)
	assert.Len(t, res, 5)
	assert.Equal(t, res, []group{
		{
			set.New('a', 'b', 'c'),
		},
		{
			set.New('a'),
			set.New('b'),
			set.New('c'),
		},
		{
			set.New('a', 'b'),
			set.New('a', 'c'),
		},
		{
			set.New('a'), set.New('a'), set.New('a'), set.New('a'),
		},
		{
			set.New('b'),
		},
	})
	assert.Equal(t, 11, sumGroupsDistinct(res))
}

func TestPartTwo(t *testing.T) {
	var input = `abc

a
b
c

ab
ac

a
a
a
a

b
`
	res := parseInput(input)
	assert.Equal(t, 6, sumGroupsIntersect(res))
}
