package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testInput = `F10
N3
F7
R90
F11`

func TestExample(t *testing.T) {
	ins, err := parseInput(testInput)
	require.NoError(t, err)
	assert.Len(t, ins, 5)
}

func TestFollowInstructions(t *testing.T) {
	ins, err := parseInput(testInput)
	require.NoError(t, err)

	pos := followInstructions(ins, position{}, dirEast)
	assert.Equal(t, 25, pos.distance())
}

func TestRotate(t *testing.T) {
	testCases := []struct {
		name   string
		in     position
		deg    int
		expOut position
	}{
		{name: "zero stays zero", deg: 90},
		{name: "zero deg", in: dirNorth, deg: 0, expOut: dirNorth},
		{name: "north to east", in: dirNorth, deg: 90, expOut: dirEast},
		{name: "east to south", in: dirEast, deg: 90, expOut: dirSouth},
		{name: "south to west", in: dirSouth, deg: 90, expOut: dirWest},
		{name: "west to north", in: dirWest, deg: 90, expOut: dirNorth},
		{name: "turn left", in: dirNorth, deg: -90, expOut: dirWest},
		{name: "90 deg", in: position{dNorth: 1, dEast: 1}, deg: 90, expOut: position{dNorth: -1, dEast: 1}},
		{name: "180 deg", in: position{dNorth: 1, dEast: 1}, deg: 180, expOut: position{dNorth: -1, dEast: -1}},
		{name: "270 deg", in: position{dNorth: 1, dEast: 1}, deg: 270, expOut: position{dNorth: 1, dEast: -1}},
		{name: "360 deg", in: position{dNorth: 1, dEast: 1}, deg: 360, expOut: position{dNorth: 1, dEast: 1}},
		{name: "normal", in: position{dNorth: 10, dEast: -1}, deg: 90, expOut: position{dNorth: 1, dEast: 10}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			out := rotate(tc.in, tc.deg)
			assert.Equal(t, tc.expOut, out)
		})
	}
}
func TestFollowInstructions2(t *testing.T) {
	ins, err := parseInput(testInput)
	require.NoError(t, err)

	pos := followInstructions2(ins, position{}, position{dNorth: 1, dEast: 10})
	assert.Equal(t, 286, pos.distance())
}
