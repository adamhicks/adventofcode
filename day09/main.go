package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
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

func checkSum(i int, window []int) bool {
	found := make(map[int]bool)
	for _, j := range window {
		k := i - j
		if found[k] {
			return true
		}
		found[j] = true
	}
	return false
}

func findFirstNoneSum(vals []int, preamble int) (int, error) {
	for i := preamble; i < len(vals); i++ {
		j := i - preamble
		if !checkSum(vals[i], vals[j:i]) {
			return vals[i], nil
		}
	}
	return 0, errors.New("no none sum found")
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
	wrong, err := findFirstNoneSum(vals, 25)
	if err != nil {
		return err
	}
	fmt.Println(wrong)
	return nil
}

func findSumRange(vals []int, target int) ([]int, error) {
	var i, j int
	var cur int
	for j <= len(vals) {
		if cur == target {
			return vals[i:j], nil
		}
		if cur < target {
			cur += vals[j]
			j++
		} else {
			cur -= vals[i]
			i++
		}
	}
	return nil, errors.New("range not found")
}

func sumMinMax(vals []int) int {
	if len(vals) == 0 {
		return 0
	}
	min, max := vals[0], vals[0]
	for _, v := range vals {
		if v > max {
			max = v
		} else if v < min {
			min = v
		}
	}
	return min + max
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
	wrong, err := findFirstNoneSum(vals, 25)
	if err != nil {
		return err
	}
	r, err := findSumRange(vals, wrong)
	if err != nil {
		return err
	}
	fmt.Println(sumMinMax(r))
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
