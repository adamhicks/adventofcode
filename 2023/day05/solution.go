package day05

import (
	_ "embed"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

//go:embed input.txt
var inputString string

type input [][]string

func parseInput(s string) input {
	sections := strings.Split(strings.TrimSpace(s), "\n\n")
	var ret input
	for _, sec := range sections {
		ret = append(ret, strings.Split(sec, "\n"))
	}
	return ret
}

var testString1 = `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4`

func numbers(line string) ([]int, error) {
	nums := strings.Fields(line)
	ret := make([]int, 0, len(nums))
	for _, n := range nums {
		val, err := strconv.Atoi(n)
		if err != nil {
			return nil, err
		}
		ret = append(ret, val)
	}
	return ret, nil
}

func runMap(start []int, sec []string) ([]int, error) {
	current := make(map[int]int)
	for _, s := range start {
		current[s] = s
	}
	mapped := make(map[int]int)

	for _, mapLine := range sec {
		nums, err := numbers(mapLine)
		if err != nil {
			return nil, err
		}
		dst, src, size := nums[0], nums[1], nums[2]
		for k := range current {
			if k >= src && k < src+size {
				mapped[k] = k - src + dst
				delete(current, k)
			}
		}
	}
	for k, v := range current {
		mapped[k] = v
	}
	var ret []int
	for _, s := range start {
		ret = append(ret, mapped[s])
	}
	return ret, nil
}

func runPartOne(s input) {
	seeds, err := numbers(strings.TrimPrefix(s[0][0], "seeds: "))
	if err != nil {
		log.Fatal(err)
	}
	current := seeds
	for _, sec := range s[1:] {
		current, err = runMap(current, sec[1:])
		if err != nil {
			log.Fatal(err)
		}
	}
	m := math.MaxInt
	for _, v := range current {
		m = min(m, v)
	}
	fmt.Println(m)
}

var testString2 = testString1

type IntRange struct {
	Start, Size int
}

func (i IntRange) End() int {
	return i.Start + i.Size
}

func (i IntRange) IsZero() bool {
	return i.Start == 0 && i.Size == 0
}

func splitRange(r, sub IntRange) (IntRange, []IntRange) {
	if r.Start >= sub.End() || r.End() <= sub.Start {
		return IntRange{}, []IntRange{r}
	}
	overlap := IntRange{Start: max(r.Start, sub.Start)}
	overlap.Size = min(r.End(), sub.End()) - overlap.Start
	var left []IntRange
	if r.Start < overlap.Start {
		left = append(left, IntRange{Start: r.Start, Size: overlap.Start - r.Start})
	}
	if r.End() > overlap.End() {
		left = append(left, IntRange{Start: overlap.End(), Size: r.End() - overlap.End()})
	}
	tot := overlap.Size
	for _, l := range left {
		tot += l.Size
	}
	return overlap, left
}

func runMapTwo(start []IntRange, sec []string) ([]IntRange, error) {
	current := start
	var ret []IntRange
	for _, l := range sec {
		nums, err := numbers(l)
		if err != nil {
			return nil, err
		}
		var next []IntRange
		dst, src, size := nums[0], nums[1], nums[2]
		srcRange := IntRange{Start: src, Size: size}

		for _, r := range current {
			overlap, left := splitRange(r, srcRange)
			next = append(next, left...)
			if !overlap.IsZero() {
				dstRange := IntRange{Start: overlap.Start + dst - src, Size: overlap.Size}
				ret = append(ret, dstRange)
			}
		}
		current = next
	}
	return append(ret, current...), nil
}

func runPartTwo(s input) {
	seeds, err := numbers(strings.TrimPrefix(s[0][0], "seeds: "))
	if err != nil {
		log.Fatal(err)
	}
	var current []IntRange
	for i := 0; i < len(seeds); i += 2 {
		current = append(current, IntRange{Start: seeds[i], Size: seeds[i+1]})
	}
	for _, sec := range s[1:] {
		current, err = runMapTwo(current, sec[1:])
		if err != nil {
			log.Fatal(err)
		}
	}
	m := math.MaxInt
	for _, v := range current {
		m = min(m, v.Start)
	}
	fmt.Println(m)
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
