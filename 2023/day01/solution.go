package day01

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var inputString string

type input []string

func parseInput(s string) input {
	return strings.Split(strings.TrimSpace(s), "\n")
}

var testString1 = `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet
`

func numberFromDigit(c rune) (int, bool) {
	if c >= '0' && c <= '9' {
		return int(c) - int('0'), true
	}
	return 0, false
}

func getLineValue(l string) int {
	var nums []int
	for _, c := range l {
		if v, ok := numberFromDigit(c); ok {
			nums = append(nums, v)
		}
	}
	return nums[0]*10 + nums[len(nums)-1]
}

func runPartOne(s input) {
	var sum int
	for _, l := range s {
		sum += getLineValue(l)
	}
	fmt.Println(sum)
}

var testString2 = `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen
`

var numbers = map[string]int{
	"one": 1, "two": 2, "three": 3, "four": 4,
	"five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9,
}

func endsInWord(s string) (int, bool) {
	for word, val := range numbers {
		if strings.HasSuffix(s, word) {
			return val, true
		}
	}
	return 0, false
}

func startsWithWord(s string) (int, bool) {
	for word, val := range numbers {
		if strings.HasPrefix(s, word) {
			return val, true
		}
	}
	return 0, false
}

func firstNumber(line string) int {
	var buf string
	for _, c := range line {
		if v, ok := numberFromDigit(c); ok {
			return v
		}
		buf += string(c)
		if v, ok := endsInWord(buf); ok {
			return v
		}
	}
	return 0
}

func lastNumber(line string) int {
	var buf string
	for i := len(line) - 1; i >= 0; i-- {
		c := rune(line[i])
		if v, ok := numberFromDigit(c); ok {
			return v
		}
		buf = string(c) + buf
		if v, ok := startsWithWord(buf); ok {
			return v
		}
	}
	return 0
}

func getLineValueWithWords(l string) int {
	first, last := firstNumber(l), lastNumber(l)
	return first*10 + last
}

func runPartTwo(s input) {
	var sum int
	for _, line := range s {
		sum += getLineValueWithWords(line)
	}
	fmt.Println(sum)
}

type Solution struct{}

func (Solution) TestPart1() {
	runPartOne(parseInput(testString1))
}

func (Solution) RunPart1() {
	runPartOne(parseInput(inputString))
}

func (Solution) TestPart2() {
	runPartTwo(parseInput(testString2))
}
func (Solution) RunPart2() {
	runPartTwo(parseInput(inputString))
}
