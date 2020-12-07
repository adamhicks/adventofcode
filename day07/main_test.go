package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testInput = `light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.
`

func TestParseInput(t *testing.T) {
	g, err := parseInput(testInput)
	require.NoError(t, err)
	assert.Len(t, g, 9)
}

func TestReverse(t *testing.T) {
	g, err := parseInput(testInput)
	require.NoError(t, err)

	rev := reverse(g)
	// reduces to 7 because of 2 terminal symbols
	assert.Len(t, rev, 7)
}

func TestFindContaining(t *testing.T) {
	g, err := parseInput(testInput)
	require.NoError(t, err)

	g = reverse(g)
	all := findAllContaining("shiny gold", g)

	assert.Equal(t, all, map[string]bool{
		"bright white": true, "muted yellow": true,
		"light red": true, "dark orange": true,
	})
}

func TestSumAll(t *testing.T) {
	g, err := parseInput(testInput)
	require.NoError(t, err)

	sum := sumContents("shiny gold", g)
	assert.Equal(t, 32, sum)
}
