package main

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestString(t *testing.T) {
	s := "1 + (2 * 3 + 4)"
	n, err := parseExpression(s)
	require.NoError(t, err)
	assert.Equal(t, s, fmt.Sprint(n))
}

func TestExample(t *testing.T) {
	testCases := []struct {
		expression string
		expValue   int
	}{
		{expression: "1 + (2 * 3) + (4 * (5 + 6))", expValue: 51},
		{expression: "2 * 3 + (4 * 5)", expValue: 26},
		{expression: "5 + (8 * 3 + 9 + 3 * 4 * 3)", expValue: 437},
		{expression: "5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", expValue: 12240},
		{expression: "((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", expValue: 13632},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			e, err := parseExpression(tc.expression)
			require.NoError(t, err)
			assert.Equal(t, tc.expValue, e.Value(), "%s vs %s", tc.expression, e)
		})
	}
}

func TestManip(t *testing.T) {
	s := "9 * 2 + 3 + 4"
	e, err := parseExpression(s)
	require.NoError(t, err)

	e = manipulatePrecedence(e)
	fmt.Println(e)
	assert.Equal(t, 81, e.Value())
}

func TestManipulate(t *testing.T) {
	testCases := []struct {
		expression string
		expValue   int
	}{
		{expression: "1 + (2 * 3) + (4 * (5 + 6))", expValue: 51},
		{expression: "2 * 3 + (4 * 5)", expValue: 46},
		{expression: "5 + (8 * 3 + 9 + 3 * 4 * 3)", expValue: 1445},
		{expression: "5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", expValue: 669060},
		{expression: "((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", expValue: 23340},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			e, err := parseExpression(tc.expression)
			require.NoError(t, err)
			e = manipulatePrecedence(e)
			assert.Equal(t, tc.expValue, e.Value(), "%s vs %s", tc.expression, e)
		})
	}
}
