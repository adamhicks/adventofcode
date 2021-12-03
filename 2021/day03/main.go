package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func defaultInput() []int {
	return parseInput(input)
}

func parseInput(s string) []int {
	var ret []int
	for _, row := range strings.Split(s, "\n") {
		var v int
		for _, c := range row {
			v <<= 1
			if c == '1' {
				v++
			}
		}
		if v == 0 {
			continue
		}
		ret = append(ret, v)
	}
	return ret
}

func testInput() []int {
	t := `00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`
	return parseInput(t)
}

func test1() {
	gamma := getPower(testInput(), true)
	epsilon := getPower(testInput(), false)
	fmt.Println(gamma, epsilon)
}

func test2() {
	ox := filterNums(testInput(), true)
	co2 := filterNums(testInput(), false)
	fmt.Println(ox, co2)
}

func maxMask(nums []int) int {
	for mask := 1; ; mask <<=1 {
		var anyNonZero bool
		for _, n := range nums {
			if n & mask > 0 {
				anyNonZero = true
				break
			}
		}
		if !anyNonZero {
			return mask>>1
		}
	}
}

func majorityOnes(nums []int, mask int) bool {
	var ones int
	for _, v := range nums {
		if v & mask > 0 {
			ones++
		}
	}
	l := len(nums)/2
	if len(nums)%2 == 0 && ones == l {
		return true
	}
	return ones > l
}

func getPower(nums []int, selMaj bool) int {
	mask := maxMask(nums)
	var ret int
	for ; mask > 0; mask >>= 1 {
		ret <<= 1
		majOnes := majorityOnes(nums, mask)
		if majOnes == selMaj {
			ret++
		}
	}
	return ret
}

func filterNums(nums []int, selMaj bool) int {
	mask := maxMask(nums)
	var ret int

	for ; mask > 0; mask >>= 1 {
		if len(nums) == 1 {
			ret |= nums[0] & mask
			continue
		}

		majOnes := majorityOnes(nums, mask)
		selOnes := majOnes == selMaj

		if selOnes {
			ret |= mask
		}

		var nextNums []int
		for _, v := range nums {
			isOne := (v & mask) > 0
			if isOne == selOnes {
				nextNums = append(nextNums, v)
			}
		}
		nums = nextNums
	}
	return ret
}

func part1() {
	in := defaultInput()
	g := getPower(in, true)
	e := getPower(in, false)
	fmt.Println(g * e)
}

func part2() {
	in := defaultInput()
	ox := filterNums(in, true)
	co2 := filterNums(in, false)
	fmt.Println(ox*co2)
}

func main() {
	part1()
	// test1()
	part2()
	// test2()
}