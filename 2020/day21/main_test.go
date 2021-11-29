package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const testInput = `mxmxvkd kfcds sqjhc nhms (contains dairy, fish)
trh fvjkl sbzzf mxmxvkd (contains dairy)
sqjhc fvjkl (contains soy)
sqjhc mxmxvkd sbzzf (contains fish)`

func TestParse(t *testing.T) {
	allFood := parseInput(testInput)
	assert.Len(t, allFood, 4)
}

func TestGetCommon(t *testing.T) {
	f := parseInput(testInput)
	allCommon := findAllergenCommonIngredients(f)
	matched := matchIngredients(allCommon)

	exp := map[string]string{
		"mxmxvkd": "dairy",
		"sqjhc":   "fish",
		"fvjkl":   "soy",
	}
	assert.Equal(t, exp, matched)
}

func TestCountUnmatched(t *testing.T) {
	f := parseInput(testInput)
	allCommon := findAllergenCommonIngredients(f)
	matched := matchIngredients(allCommon)

	i := countUnmatched(f, matched)
	assert.Equal(t, 5, i)
}

func TestArrange(t *testing.T) {
	f := parseInput(testInput)

	allCommon := findAllergenCommonIngredients(f)
	matched := matchIngredients(allCommon)
	s := arrangeIngredients(matched)

	assert.Equal(t, "mxmxvkd,sqjhc,fvjkl", s)
}
