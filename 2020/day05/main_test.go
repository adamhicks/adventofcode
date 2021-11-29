package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseSeat(t *testing.T) {
	i := parseBinary("FBFBBFF")
	assert.Equal(t, 44, i)
	j := parseBinary("RLR")
	assert.Equal(t, 5, j)

	assert.Equal(t, 357, parseBinary("FBFBBFFRLR"))
	assert.Equal(t, 567, parseBinary("BFFFBBFRRR"))
	assert.Equal(t, 119, parseBinary("FFFBBBFRRR"))
	assert.Equal(t, 820, parseBinary("BBFFBBFRLL"))
}
