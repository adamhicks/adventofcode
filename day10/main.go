package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

func parseInput(content string) ([]int, error) {
	var ret []int
	for _, l := range strings.Split(content, "\n") {
		if l == "" {
			continue
		}
		i, err := strconv.Atoi(l)
		if err != nil {
			return nil, err
		}
		ret = append(ret, i)
	}
	return ret, nil
}

func countOneAndThreeDiffs(vals []int) (int, int) {
	sort.Ints(vals)
	var ones, threes int
	var lastJoltage int
	for _, joltage := range vals {
		diff := joltage - lastJoltage
		lastJoltage = joltage

		if diff == 1 {
			ones++
		} else if diff == 3 {
			threes++
		}
	}
	// Last adapter is always +3
	threes++
	return ones, threes
}

func runPartOne() error {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		return err
	}
	vals, err := parseInput(string(content))
	if err != nil {
		return err
	}
	ones, threes := countOneAndThreeDiffs(vals)
	fmt.Println(ones * threes)
	return nil
}

func getPermVal(idx int, offset int, perms []int64, vals []int) int64 {
	checkIdx := idx - offset
	if checkIdx < 0 || vals[idx]-vals[checkIdx] > 3 {
		return 0
	}
	return perms[checkIdx]
}

func countPermutations(vals []int) int64 {
	sort.Ints(vals)
	vals = append([]int{0}, vals...)
	perms := make([]int64, len(vals))
	perms[0] = 1

	for idx := 1; idx < len(vals); idx++ {
		var sumP int64
		sumP += getPermVal(idx, 3, perms, vals)
		sumP += getPermVal(idx, 2, perms, vals)
		sumP += getPermVal(idx, 1, perms, vals)
		perms[idx] = sumP
	}
	return perms[len(perms)-1]
}

func runPartTwo() error {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		return err
	}
	vals, err := parseInput(string(content))
	if err != nil {
		return err
	}
	fmt.Println(countPermutations(vals))
	return nil
}

func main() {
	if err := runPartOne(); err != nil {
		log.Fatal(err)
	}
	if err := runPartTwo(); err != nil {
		log.Fatal(err)
	}
}
