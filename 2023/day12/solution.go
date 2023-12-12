package day12

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

var testString1 = `???.### 1,1,3
.??..??...?##. 1,1,3
?#?#?#?#?#?#?#? 1,3,1,6
????.#...#... 4,1,1
????.######..#####. 1,6,5
?###???????? 3,2,1`

func parseLine(l string) (string, []int, error) {
	spr, nums, _ := strings.Cut(l, " ")
	var conds []int
	for _, n := range strings.Split(nums, ",") {
		v, err := strconv.Atoi(n)
		if err != nil {
			return "", nil, err
		}
		conds = append(conds, v)
	}
	return spr, conds, nil
}

func runPartOne(s input) error {
	var sum int
	for _, l := range s {
		spr, conds, err := parseLine(l)
		if err != nil {
			return err
		}
		sum += possibleCount(spr, conds)
	}
	fmt.Println(sum)
	return nil
}

var testString2 = testString1

func unfoldSprings(s string) string {
	ret := s
	for i := 0; i < 4; i++ {
		ret = ret + "?" + s
	}
	return ret
}

func unfoldConditions(c []int) []int {
	ret := make([]int, len(c)*5)
	for i := 0; i < len(ret); i += len(c) {
		copy(ret[i:], c)
	}
	return ret
}

func passCount(s string, conditions []int) (int, bool) {
	var cur int
	var cIdx int
	for _, r := range s {
		switch r {
		case '.':
			if cur > 0 {
				if cur == conditions[cIdx] {
					cIdx++
					cur = 0
				} else {
					return 0, false
				}
			}
		case '#':
			cur++
			if cIdx >= len(conditions) || cur > conditions[cIdx] {
				return 0, false
			}
		}
	}
	if cur > 0 && cur == conditions[cIdx] {
		cIdx++
	}
	return cIdx, true
}

type solution struct {
	String    string
	Satisfied int
	Count     int
}

func condenseSolutions(solutions []solution) []solution {
	satisfied := make(map[int][]solution)
	var ret []solution
	for _, s := range solutions {
		if s.String[len(s.String)-1] == '.' {
			satisfied[s.Satisfied] = append(satisfied[s.Satisfied], s)
		} else {
			ret = append(ret, s)
		}
	}

	for _, sols := range satisfied {
		first := sols[0]
		for i := 1; i < len(sols); i++ {
			first.Count += sols[i].Count
		}
		ret = append(ret, first)
	}
	return ret
}

func possibleCount(s string, cond []int) int {
	sols := []solution{{Count: 1}}
	for _, r := range s {
		var nxt []solution
		switch r {
		case '?':
			for _, old := range sols {
				s1, s2 := old, old
				s1.String += "."
				s2.String += "#"
				nxt = append(nxt, s1, s2)
			}
		default:
			for _, old := range sols {
				old.String += string(r)
				nxt = append(nxt, old)
			}
		}
		sols = nil
		for _, newSol := range nxt {
			count, ok := passCount(newSol.String, cond)
			if !ok {
				continue
			}
			newSol.Satisfied = count
			sols = append(sols, newSol)
		}
		sols = condenseSolutions(sols)
	}
	var sum int
	for _, sol := range sols {
		if sol.Satisfied < len(cond) {
			continue
		}
		sum += sol.Count
	}
	return sum
}

func runPartTwo(s input) error {
	var sum int
	for _, l := range s {
		spr, conds, err := parseLine(l)
		if err != nil {
			return err
		}
		spr = unfoldSprings(spr)
		conds = unfoldConditions(conds)
		poss := possibleCount(spr, conds)
		sum += poss
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
