package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func parseInput() ([][]int) {
	var ret [][]int
	for _, row := range strings.Split(input, "\n") {
		r := make([]int, len(row))
		for i, c := range row {
			if c == '1' {
				r[i] = 1
			}
		}
		if len(row) == 0 {
			continue
		}
		ret = append(ret, r)
	}
	return ret
}

func accumulateOnes(bin [][]int) []int {
	out := make([]int, len(bin[0]))
	for _, row := range bin {
		for i, v := range row {
			if v == 1 {
				out[i]++
			}
		}
	}
	return out
}

var testIn = [][]int{
	{0,0,1,0,0},
	{1,1,1,1,0},
	{1,0,1,1,0},
	{1,0,1,1,1},
	{1,0,1,0,1},
	{0,1,1,1,1},
	{0,0,1,1,1},
	{1,1,1,0,0},
	{1,0,0,0,0},
	{1,1,0,0,1},
	{0,0,0,1,0},
	{0,1,0,1,0},
}

func test1() {
	g, e := getPower(testIn)
	fmt.Println(g, e)
}

func test2() {
	ox := toDec(filterRows(testIn, true))
	co2 := toDec(filterRows(testIn, false))
	fmt.Println(ox, co2)
}

func toDec(bin []int) int {
	var ret int
	for _, v := range bin {
		ret <<= 1;
		if v == 1 {
			ret++
		}
	}
	return ret
}

func getPower(bin [][]int) (int, int) {
	ones := accumulateOnes(bin)
	maj := len(bin) / 2

	gamma := make([]int, len(ones))
	epsilon := make([]int, len(ones))
	for i, v := range ones {
		if v > maj {
			gamma[i] = 1;
		} else {
			epsilon[i] = 1;
		}
	}
	return toDec(gamma), toDec(epsilon)
}

func filterRows(bin [][]int, selMaj bool) []int {
	kept := bin
	for idx := 0; len(kept) > 1; idx++ {
		var ones int
		for _, v := range kept {
			if v[idx] == 1 {
				ones++
			}
		}
		l := len(kept)/2
		majOnes := ones > l
		if len(kept)%2 == 0 && ones == l {
			majOnes = true
		}
		selOnes := majOnes == selMaj
		var nextKept [][]int
		for _, v := range kept {
			if v[idx] == 1 && selOnes || v[idx] == 0 && !selOnes {
				nextKept = append(nextKept, v)
			}
		}
		kept = nextKept
	}
	return kept[0]
}

func part1() {
	g, e := getPower(parseInput())
	fmt.Println(g * e)
}

func part2() {
	in := parseInput()
	ox := toDec(filterRows(in, true))
	co2 := toDec(filterRows(in, false))
	fmt.Println(ox*co2)
}

func main() {
	part1()
	// test1()
	part2()
	// test2()
}