package day07

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var inputString string

type input []string

func parseInput(s string) input {
	return strings.Split(strings.TrimSpace(s), "\n")
}

var testString1 = `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483
`

var jokes bool

func ofAKind(a string) [2]int {
	cards := make(map[rune]int)
	for _, r := range a {
		cards[r]++
	}
	if jokes && cards['J'] > 0 {
		var maxCard rune
		var maxCount int
		for card, count := range cards {
			if card == 'J' || count <= maxCount {
				continue
			}
			maxCount = count
			maxCard = card
		}
		cards[maxCard] += cards['J']
		delete(cards, 'J')
	}
	var pairsOrBetter []int
	for _, count := range cards {
		if count > 1 {
			pairsOrBetter = append(pairsOrBetter, count)
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(pairsOrBetter)))
	var ret [2]int
	if len(pairsOrBetter) > 0 {
		ret[0] = pairsOrBetter[0]
	}
	if len(pairsOrBetter) > 1 {
		ret[1] = pairsOrBetter[1]
	}
	return ret
}

var value = map[rune]int{
	'A': 13,
	'K': 12,
	'Q': 11,
	'J': 10,
	'T': 9,
	'9': 8,
	'8': 7,
	'7': 6,
	'6': 5,
	'5': 4,
	'4': 3,
	'3': 2,
	'2': 1,
}

// kindsCmp works on kinds in a hand
// returns > 0 if a is better than b
// < 0 if b is better than a
// and 0 if a == b
func kindsCmp(a, b [2]int) int {
	if a == b {
		return 0
	}
	for i := 0; i < len(a); i++ {
		d := a[i] - b[i]
		if d == 0 {
			continue
		}
		return d
	}
	panic("unreachable")
}

// cardCmp returns > 0 if a is better than b
// < 0 if b is better than a
// and 0 if a == b
func cardCmp(a, b rune) int {
	return value[a] - value[b]
}

func isBetterThan(a, b string) bool {
	kCmp := kindsCmp(ofAKind(a), ofAKind(b))
	if kCmp != 0 {
		return kCmp > 0
	}
	for i, c := range a {
		cmp := cardCmp(c, rune(b[i]))
		if cmp > 0 {
			return true
		} else if cmp < 0 {
			return false
		}
	}
	return false
}

func runPartOne(s input) error {
	sort.Slice(s, func(i, j int) bool {
		handA, _, _ := strings.Cut(s[i], " ")
		handB, _, _ := strings.Cut(s[j], " ")
		return !isBetterThan(handA, handB)
	})
	var sum int
	for i, l := range s {
		_, val, _ := strings.Cut(l, " ")
		ante, err := strconv.Atoi(val)
		if err != nil {
			return err
		}
		sum += ante * (i + 1)
	}
	fmt.Println(sum)
	return nil
}

var testString2 = testString1

func runPartTwo(s input) error {
	value['J'] = 0
	jokes = true

	sort.Slice(s, func(i, j int) bool {
		handA, _, _ := strings.Cut(s[i], " ")
		handB, _, _ := strings.Cut(s[j], " ")
		return !isBetterThan(handA, handB)
	})
	var sum int
	for i, l := range s {
		_, val, _ := strings.Cut(l, " ")
		ante, err := strconv.Atoi(val)
		if err != nil {
			return err
		}
		sum += ante * (i + 1)
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
	//return nil
	return runPartTwo(parseInput(inputString))
}
