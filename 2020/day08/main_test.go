package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testInput = `nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6
`

func TestParseInput(t *testing.T) {
	ins, err := parseInput(testInput)
	require.NoError(t, err)
	assert.Len(t, ins, 9)
}

func TestDetectLoop(t *testing.T) {
	ins, err := parseInput(testInput)
	require.NoError(t, err)

	s := detectLoop(ins)
	assert.Equal(t, 1, s.Pos)
	assert.Equal(t, 5, s.Acc)
}

func TestDetectCorrupt(t *testing.T) {
	ins, err := parseInput(testInput)
	require.NoError(t, err)

	s, err := detectCorruptExit(ins)
	require.NoError(t, err)

	assert.Equal(t, 9, s.Pos)
	assert.Equal(t, 8, s.Acc)
	assert.Equal(t, true, s.HasFixedInstruction)
	assert.Equal(t, 7, s.FixedPos)
}
