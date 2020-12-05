package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseSeat(t *testing.T) {
	i := parseAddress("FBFBBFF")
	assert.Equal(t, 44, i)
	j := parseAddress("RLR")
	assert.Equal(t, 5, j)

	r, c := parseSeat("FBFBBFFRLR")
	assert.Equal(t, 44, r)
	assert.Equal(t, 5, c)

	assert.Equal(t, 357, getSeatID("FBFBBFFRLR"))
	assert.Equal(t, 567, getSeatID("BFFFBBFRRR"))
	assert.Equal(t, 119, getSeatID("FFFBBBFRRR"))
	assert.Equal(t, 820, getSeatID("BBFFBBFRLL"))
}
