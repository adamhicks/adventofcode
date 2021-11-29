package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIntersect(t *testing.T) {
	testCases := []struct {
		name         string
		iOne, iTwo   interval
		expIntersect bool
		expUnion     interval
	}{
		{name: "overlap",
			iOne:         interval{low: 10, high: 20},
			iTwo:         interval{low: 15, high: 22},
			expIntersect: true,
			expUnion:     interval{low: 10, high: 22},
		},
		{name: "subset",
			iOne:         interval{low: 10, high: 20},
			iTwo:         interval{low: 15, high: 17},
			expIntersect: true,
			expUnion:     interval{low: 10, high: 20},
		},
		{name: "superset",
			iOne:         interval{low: 10, high: 20},
			iTwo:         interval{low: 8, high: 22},
			expIntersect: true,
			expUnion:     interval{low: 8, high: 22},
		},
		{name: "touching",
			iOne:         interval{low: 10, high: 20},
			iTwo:         interval{low: 20, high: 30},
			expIntersect: true,
			expUnion:     interval{low: 10, high: 20},
		},
		{name: "separate",
			iOne:         interval{low: 10, high: 20},
			iTwo:         interval{low: 21, high: 30},
			expIntersect: false,
			expUnion:     interval{low: 10, high: 30},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expIntersect, tc.iOne.Intersects(tc.iTwo))
		})
	}
}

const testInput = `class: 1-3 or 5-7
row: 6-11 or 33-44
seat: 13-40 or 45-50

your ticket:
7,1,14

nearby tickets:
7,3,47
40,4,50
55,2,20
38,6,12`

func TestParseInput(t *testing.T) {
	ti, err := parseInput(testInput)
	require.NoError(t, err)

	exp := ticketInfo{
		fields: []field{
			{name: "class", intervals: []interval{
				{low: 1, high: 4}, {low: 5, high: 8},
			}},
			{name: "row", intervals: []interval{
				{low: 6, high: 12}, {low: 33, high: 45},
			}},
			{name: "seat", intervals: []interval{
				{low: 13, high: 41}, {low: 45, high: 51},
			}},
		},
		yourTicket: []int{7, 1, 14},
		nearbyTickets: [][]int{
			{7, 3, 47},
			{40, 4, 50},
			{55, 2, 20},
			{38, 6, 12},
		},
	}
	assert.Equal(t, exp, ti)
}

func TestSimplify(t *testing.T) {
	ti, err := parseInput(testInput)
	require.NoError(t, err)

	ints := simplifyFields(ti.fields)
	exp := []interval{
		{low: 1, high: 4},
		{low: 5, high: 12},
		{low: 13, high: 51},
	}
	assert.Equal(t, exp, ints)
}

func TestCheckTicket(t *testing.T) {
	ti, err := parseInput(testInput)
	require.NoError(t, err)

	ints := simplifyFields(ti.fields)

	var invalidVals [][]int
	for _, tick := range ti.nearbyTickets {
		invalidVals = append(invalidVals, checkTicketSimple(tick, ints))
	}

	exp := [][]int{nil, {4}, {55}, {12}}
	assert.Equal(t, exp, invalidVals)
}
