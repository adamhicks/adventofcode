package aoc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOverlap(t *testing.T) {
	testCases := []struct {
		name       string
		a, b       Rect
		expOverlap Rect
	}{
		{name: "distinct",
			a: Rect{From: Vec2{X: 0, Y: 0}, To: Vec2{X: 2, Y: 2}},
			b: Rect{From: Vec2{X: 4, Y: 0}, To: Vec2{X: 6, Y: 2}},
		},
		{name: "touching",
			a: Rect{From: Vec2{X: 0, Y: 0}, To: Vec2{X: 2, Y: 2}},
			b: Rect{From: Vec2{X: 2, Y: 0}, To: Vec2{X: 6, Y: 2}},
		},
		{name: "inside",
			a:          Rect{From: Vec2{X: 0, Y: 0}, To: Vec2{X: 20, Y: 20}},
			b:          Rect{From: Vec2{X: 5, Y: 5}, To: Vec2{X: 10, Y: 10}},
			expOverlap: Rect{From: Vec2{X: 5, Y: 5}, To: Vec2{X: 10, Y: 10}},
		},
		{name: "outside",
			a:          Rect{From: Vec2{X: 5, Y: 5}, To: Vec2{X: 10, Y: 10}},
			b:          Rect{From: Vec2{X: 0, Y: 0}, To: Vec2{X: 20, Y: 20}},
			expOverlap: Rect{From: Vec2{X: 5, Y: 5}, To: Vec2{X: 10, Y: 10}},
		},
		{name: "overlap",
			a:          Rect{From: Vec2{X: 0, Y: 0}, To: Vec2{X: 20, Y: 20}},
			b:          Rect{From: Vec2{X: 15, Y: 15}, To: Vec2{X: 30, Y: 30}},
			expOverlap: Rect{From: Vec2{X: 15, Y: 15}, To: Vec2{X: 20, Y: 20}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			over := tc.a.Overlap(tc.b)
			assert.Equal(t, tc.expOverlap, over)
		})
	}
}
