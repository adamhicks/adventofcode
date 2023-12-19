package day19

import (
	_ "embed"
	"fmt"
	"github.com/adamhicks/adventofcode/2023/aoc"
	"strconv"
	"strings"
)

//go:embed input.txt
var inputString string

func parseInput(s string) ([]string, []string) {
	parts := strings.Split(strings.TrimSpace(s), "\n\n")
	return strings.Split(parts[0], "\n"), strings.Split(parts[1], "\n")
}

var testString1 = `px{a<2006:qkq,m>2090:A,rfg}
pv{a>1716:R,A}
lnx{m>1548:A,A}
rfg{s<537:gd,x>2440:R,A}
qs{s>3448:A,lnx}
qkq{x<1416:A,crn}
crn{x>2662:A,R}
in{s<1351:px,qqz}
qqz{s>2770:qs,m<1801:hdj,R}
gd{a>3333:R,R}
hdj{m>838:A,pv}

{x=787,m=2655,a=1222,s=2876}
{x=1679,m=44,a=2067,s=496}
{x=2036,m=264,a=79,s=2244}
{x=2461,m=1339,a=466,s=291}
{x=2127,m=1623,a=2188,s=1013}`

type state map[string]int

func parseStates(s []string) []state {
	var ret []state
	for _, l := range s {
		l = strings.Trim(l, "{}")
		m := make(state)
		for _, p := range strings.Split(l, ",") {
			name, valStr, _ := strings.Cut(p, "=")
			val := aoc.Must(strconv.Atoi(valStr))
			m[name] = val
		}
		ret = append(ret, m)
	}
	return ret
}

type rule struct {
	Var    string
	Cmp    rune
	Val    int
	Target string
}

type flow struct {
	Rules   []rule
	Default string
}

func parseFlows(flows []string) map[string]flow {
	ret := make(map[string]flow)
	for _, l := range flows {
		i := strings.Index(l, "{")
		name := l[:i]
		l = l[i+1 : len(l)-1]
		var f flow
		for _, rulePart := range strings.Split(l, ",") {
			j := strings.Index(rulePart, ":")
			if j < 0 {
				f.Default = rulePart
				continue
			}
			r := rule{
				Var:    rulePart[:1],
				Cmp:    rune(rulePart[1]),
				Val:    aoc.Must(strconv.Atoi(rulePart[2:j])),
				Target: rulePart[j+1:],
			}
			f.Rules = append(f.Rules, r)
		}
		ret[name] = f
	}
	return ret
}

func eval(f flow, s state) string {
	for _, r := range f.Rules {
		v := s[r.Var]
		if r.Cmp == '>' && v > r.Val {
			return r.Target
		} else if r.Cmp == '<' && v < r.Val {
			return r.Target
		}
	}
	return f.Default
}

func isFlowStateAccepted(flows map[string]flow, in state) bool {
	cur := "in"
	for {
		cur = eval(flows[cur], in)
		switch cur {
		case "A":
			return true
		case "R":
			return false
		}
	}
}

func runPartOne(flows, states []string) error {
	f := parseFlows(flows)
	var sum int
	for _, s := range parseStates(states) {
		if isFlowStateAccepted(f, s) {
			for _, v := range s {
				sum += v
			}
		}
	}
	fmt.Println(sum)
	return nil
}

var testString2 = testString1

func copyState(a state) state {
	ret := make(state)
	for k, v := range a {
		ret[k] = v
	}
	return ret
}

func perms(from, to state) int {
	p := 1
	for k := range to {
		d := to[k] - from[k]
		if d <= 0 {
			return 0
		}
		p *= d + 1
	}
	return p
}

func flowSplit(name string, flows map[string]flow, from, to state) (int, int) {
	if name == "A" {
		return perms(from, to), 0
	} else if name == "R" {
		return 0, perms(from, to)
	}
	var accept, reject int
	f := flows[name]
	for _, r := range f.Rules {
		innerFrom, innerTo := copyState(from), copyState(to)
		if r.Cmp == '<' {
			innerTo[r.Var] = r.Val - 1
			from[r.Var] = r.Val
		} else {
			innerFrom[r.Var] = r.Val + 1
			to[r.Var] = r.Val
		}
		if innerFrom[r.Var] < innerTo[r.Var] {
			acc, rej := flowSplit(r.Target, flows, innerFrom, innerTo)
			accept += acc
			reject += rej
		}
	}
	a, r := flowSplit(f.Default, flows, from, to)
	accept += a
	reject += r
	return accept, reject
}

func runPartTwo(flows, _ []string) error {
	f := parseFlows(flows)
	from := state{"x": 1, "m": 1, "a": 1, "s": 1}
	to := state{"x": 4000, "m": 4000, "a": 4000, "s": 4000}
	a, _ := flowSplit("in", f, from, to)
	fmt.Println(a)
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
