package main

import (
	"errors"
	"fmt"
	"log"
	"sort"
	"strconv"

	"github.com/adamhicks/adventofcode/util"
)

func convertToInts(lines []string) ([]int, error) {
	nums := make([]int, 0, len(lines))
	for _, l := range lines {
		i, err := strconv.Atoi(l)
		if err != nil {
			return nil, err
		}
		nums = append(nums, i)
	}
	return nums, nil
}

func findPairSum(nums []int, sum int) ([]int, error) {
	if len(nums) == 0 {
		return nil, errors.New("empty nums")
	}
	i, j := 0, len(nums)-1

	for i != j {
		thisSum := nums[i] + nums[j]
		if thisSum == sum {
			return []int{nums[i], nums[j]}, nil
		} else if thisSum > sum {
			j--
		} else {
			i++
		}
	}
	return nil, errors.New("no pair found")
}

func findTrebleSum(nums []int, sum int) ([]int, error) {
	for i, m := range nums {
		left := sum - m
		if left <= 0 {
			return nil, errors.New("single number larger than sum")
		}
		il, err := findPairSum(nums[i+1:], left)
		if err != nil {
			continue
		}
		return append(il, m), nil
	}
	return nil, errors.New("no treble found")
}

func fetchInput() ([]int, error) {
	lines, err := util.ReadInput("input.txt")
	if err != nil {
		return nil, err
	}
	nums, err := convertToInts(lines)
	if err != nil {
		return nil, err
	}
	sort.Ints(nums)
	return nums, nil
}

func main() {
	nums, err := fetchInput()
	if err != nil {
		log.Fatal(err)
	}
	factors, err := findTrebleSum(nums, 2020)
	if err != nil {
		log.Fatal(err)
	}
	var prod int64 = 1
	for _, i := range factors {
		prod *= int64(i)
	}
	fmt.Println(factors, prod)
}
