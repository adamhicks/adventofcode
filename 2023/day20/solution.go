package day20

import (
	_ "embed"
	"fmt"
	"github.com/adamhicks/adventofcode/2023/aoc"
	"strings"
)

//go:embed input.txt
var inputString string

type input []string

func parseInput(s string) input {
	return strings.Split(strings.TrimSpace(s), "\n")
}

var testString1 = `broadcaster -> a, b, c
%a -> b
%b -> c
%c -> inv
&inv -> a`

type pulse struct {
	From   string
	Target string
	High   bool
}

type state struct {
	Inputs       map[string][]string
	Outputs      map[string][]string
	FlipFlops    map[string]bool
	Conjunctions map[string]map[string]bool
}

func initState(s input) state {
	ret := state{
		Inputs:       make(map[string][]string),
		Outputs:      make(map[string][]string),
		FlipFlops:    make(map[string]bool),
		Conjunctions: make(map[string]map[string]bool),
	}
	for _, l := range s {
		name, toStr, _ := strings.Cut(l, " -> ")

		children := strings.Split(toStr, ", ")
		var key string
		switch name[0] {
		case '%':
			key = name[1:]
			ret.FlipFlops[key] = false
		case '&':
			key = name[1:]
			ret.Conjunctions[key] = make(map[string]bool)
		default:
			key = name
		}
		ret.Outputs[key] = children
	}
	for p, children := range ret.Outputs {
		for _, c := range children {
			ret.Inputs[c] = append(ret.Inputs[c], p)
		}
	}
	for n, ins := range ret.Conjunctions {
		for _, in := range ret.Inputs[n] {
			ins[in] = false
		}
	}
	return ret
}

func queuePulses(q []pulse, s state, from string, high bool) []pulse {
	for _, c := range s.Outputs[from] {
		q = append(q, pulse{Target: c, High: high, From: from})
	}
	return q
}

func run(s state, start []pulse) []pulse {
	q := start
	hist := make([]pulse, 0, 1000)
	for len(q) > 0 {
		p := q[0]
		hist = append(hist, p)
		q = q[1:]

		if p.Target == "broadcaster" {
			q = queuePulses(q, s, p.Target, p.High)
			continue
		}

		if on, ok := s.FlipFlops[p.Target]; ok {
			if p.High {
				continue
			}
			on = !on
			var sendHigh bool
			if on {
				sendHigh = true
			}
			s.FlipFlops[p.Target] = on
			q = queuePulses(q, s, p.Target, sendHigh)
		}
		if sub, ok := s.Conjunctions[p.Target]; ok {
			sub[p.From] = p.High
			var sendHigh bool
			for _, v := range sub {
				if !v {
					sendHigh = true
				}
			}
			q = queuePulses(q, s, p.Target, sendHigh)
		}
	}
	return hist
}

func runPartOne(s input) error {
	st := initState(s)
	var low, high int
	for i := 0; i < 1000; i++ {
		h := run(st, []pulse{{Target: "broadcaster", From: "button"}})
		for _, p := range h {
			if p.High {
				high++
			} else {
				low++
			}
		}
	}
	fmt.Println(high * low)
	return nil
}

var testString2 = `broadcaster -> a, b
%a -> a1
%a1 -> a2
%a2 -> a3
&a3 -> ada
%b -> b1
%b1 -> ada
&ada -> rx
`

func runPartTwo(s input) error {
	st := initState(s)
	target := st.Inputs["rx"][0]

	m := make(map[string]int)
	for i := 1; ; i++ {
		h := run(st, []pulse{{Target: "broadcaster", From: "button"}})
		for _, p := range h {
			if !p.High || p.Target != target {
				continue
			}
			if _, ok := m[p.From]; ok {
				continue
			}
			m[p.From] = i
		}
		if len(m) == len(st.Inputs[target]) {
			break
		}
	}

	align := 1
	for _, i := range m {
		align = aoc.LCM(align, i)
	}
	fmt.Println(align)
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
