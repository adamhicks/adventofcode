package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testInput = `mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
mem[8] = 11
mem[7] = 101
mem[8] = 0`

func TestParse(t *testing.T) {
	ins, err := parseInput(testInput)
	require.NoError(t, err)
	assert.Len(t, ins, 4)
}

func TestShift(t *testing.T) {
	var m intMask
	m.shift(1, 0)
	m.shift(0, 1)
	m.shift(0, 0)
	assert.Equal(t, int64(4), m.set)
	assert.Equal(t, int64(2), m.clear)
}

func TestExecuteMask(t *testing.T) {
	p := newProgramV1()

	op := setMask{mask: "XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X"}
	run(p, []instruction{op})

	assert.Equal(t, int64(64), p.valMask.set)
	assert.Equal(t, int64(2), p.valMask.clear)
}

func TestExecuteVal(t *testing.T) {
	p := newProgramV1()
	p.valMask = intMask{set: 64, clear: 2}

	op := setValue{address: 8, value: 11}
	run(p, []instruction{op})

	assert.Len(t, p.mem, 1)
	assert.Equal(t, int64(73), p.mem[8])
}

func TestExample(t *testing.T) {
	ins, err := parseInput(testInput)
	require.NoError(t, err)

	p := newProgramV1()
	run(p, ins)
	assert.Equal(t, int64(165), p.sum())
}

func TestMaskV2(t *testing.T) {
	p := newProgramV2()
	p.setMask("10X01")

	assert.Equal(t, []intMask{
		{set: 17, clear: 4},
		{set: 21, clear: 0},
	}, p.addrMasks,
	)
}

const example2 = `mask = 000000000000000000000000000000X1001X
mem[42] = 100
mask = 00000000000000000000000000000000X0XX
mem[26] = 1`

func TestExample2(t *testing.T) {
	ins, err := parseInput(example2)
	require.NoError(t, err)

	p := newProgramV2()
	run(p, ins)
	assert.Equal(t, int64(208), p.sum())
}
