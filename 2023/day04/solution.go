package day04

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var inputString string

type input []string

func parseInput(s string) input {
	return strings.Split(strings.TrimSpace(s), "\n")
}

var testString1 = `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11
`

func parseFixedNumbers(line string) ([]int, error) {
	var ret []int
	for idx := 0; idx < len(line); idx += 3 {
		s := line[idx : idx+3]
		s = strings.TrimSpace(s)
		num, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		ret = append(ret, num)
	}
	return ret, nil
}

func parseCard(line string) (int, []int, []int, error) {
	_, card, _ := strings.Cut(line, ":")
	winners, got, _ := strings.Cut(card, " |")

	winNums, err := parseFixedNumbers(winners)
	if err != nil {
		return 0, nil, nil, err
	}
	gotNums, err := parseFixedNumbers(got)
	if err != nil {
		return 0, nil, nil, err
	}
	return 0, winNums, gotNums, nil
}

func numWinners(winners, got []int) int {
	win := make(map[int]bool)
	for _, w := range winners {
		win[w] = true
	}
	var sum int
	for _, v := range got {
		if !win[v] {
			continue
		}
		sum++
	}
	return sum
}

func runPartOne(s input) error {
	var sum int
	for _, l := range s {
		_, win, got, err := parseCard(l)
		if err != nil {
			return err
		}
		wins := numWinners(win, got)
		var val int
		if wins > 0 {
			val = 1 << (wins - 1)
		}
		sum += val
	}
	fmt.Println(sum)
	return nil
}

var testString2 = `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11
`

func runPartTwo(s input) error {
	cards := make([]int, len(s))
	for i, l := range s {
		// add the initial card
		cards[i]++
		_, win, got, err := parseCard(l)
		if err != nil {
			return err
		}
		wins := numWinners(win, got)
		for j := 0; j < wins; j++ {
			cards[i+j+1] += cards[i]
		}
	}
	var sum int
	for _, c := range cards {
		sum += c
	}
	fmt.Println(sum)
	return nil
}

type Solution struct{}

func (Solution) TestPart1() error {
	return runPartOne(parseInput(testString1))
}

func (Solution) RunPart1() error {
	return runPartOne(parseInput(inputString))
}

func (Solution) TestPart2() error {
	return runPartTwo(parseInput(testString2))
}

func (Solution) RunPart2() error {
	return runPartTwo(parseInput(inputString))
}
