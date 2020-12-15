package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample(t *testing.T) {
	testCases := []struct {
		seq []int
		n   int
		exp int
	}{
		{seq: []int{0, 3, 6}, n: 10, exp: 0},
		{seq: []int{2, 1, 3}, n: 2020, exp: 10},
		{seq: []int{1, 2, 3}, n: 2020, exp: 27},
		{seq: []int{2, 3, 1}, n: 2020, exp: 78},
		{seq: []int{3, 2, 1}, n: 2020, exp: 438},
		{seq: []int{3, 1, 2}, n: 2020, exp: 1836},
		{seq: []int{0, 3, 6}, n: 30000000, exp: 175594},
		{seq: []int{1, 3, 2}, n: 30000000, exp: 2578},
		{seq: []int{2, 1, 3}, n: 30000000, exp: 3544142},
		{seq: []int{1, 2, 3}, n: 30000000, exp: 261214},
		{seq: []int{2, 3, 1}, n: 30000000, exp: 6895259},
		{seq: []int{3, 2, 1}, n: 30000000, exp: 18},
		{seq: []int{3, 1, 2}, n: 30000000, exp: 362},
	}
	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			res := runSequence(tc.seq, tc.n)
			assert.Equal(t, tc.exp, res)
		})
	}
}
