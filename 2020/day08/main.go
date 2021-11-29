package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

// State is the execution state (memory)
type State struct {
	Pos, Acc            int
	HasFixedInstruction bool
	FixedPos            int
}

// Instruction modify the state
type Instruction interface {
	Execute(State) State
}

// Nop is a no-op
type Nop int64

// Execute no-op
func (Nop) Execute(s State) State {
	s.Pos++
	return s
}

// Jump moves the pointer to a relative instruction
type Jump int

// Execute jump
func (j Jump) Execute(s State) State {
	s.Pos += int(j)
	return s
}

// Acc adds to the acc memory state
type Acc int64

// Execute accumulate
func (a Acc) Execute(s State) State {
	s.Acc += int(a)
	s.Pos++
	return s
}

func parseInstruction(s string) (Instruction, error) {
	parts := strings.Split(s, " ")
	i, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, err
	}

	switch parts[0] {
	case "nop":
		return Nop(i), nil
	case "jmp":
		return Jump(i), nil
	case "acc":
		return Acc(i), err
	}
	return nil, fmt.Errorf("unknown instruction '%s'", s)
}

func parseInput(content string) ([]Instruction, error) {
	var ret []Instruction
	for _, l := range strings.Split(content, "\n") {
		if l == "" {
			continue
		}
		i, err := parseInstruction(l)
		if err != nil {
			return nil, err
		}
		ret = append(ret, i)
	}
	return ret, nil
}

func detectLoop(prog []Instruction) State {
	var s State
	been := make([]bool, len(prog))

	for {
		if been[s.Pos] {
			return s
		}
		been[s.Pos] = true
		s = prog[s.Pos].Execute(s)
	}
}

func runPartOne() error {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		return err
	}
	ins, err := parseInput(string(content))
	if err != nil {
		return err
	}
	s := detectLoop(ins)
	fmt.Println(s.Acc)
	return nil
}

func detectCorruptExit(prog []Instruction) (State, error) {
	// We use breadth first so if anything loops then
	// we still get to the answer
	states := []State{{}}
	for len(states) > 0 {
		cur := states[0]
		states = states[1:]

		if cur.Pos == len(prog) {
			return cur, nil
		}
		i := prog[cur.Pos]
		states = append(states, i.Execute(cur))
		if !cur.HasFixedInstruction {
			alt := cur
			alt.HasFixedInstruction = true
			alt.FixedPos = cur.Pos
			switch v := i.(type) {
			case Nop:
				alt = Jump(v).Execute(alt)
			case Jump:
				alt = Nop(0).Execute(alt)
			}
			states = append(states, alt)
		}
	}
	return State{}, errors.New("no fix detected")
}

func runPartTwo() error {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		return err
	}
	ins, err := parseInput(string(content))
	if err != nil {
		return err
	}
	s, err := detectCorruptExit(ins)
	if err != nil {
		return err
	}
	fmt.Println(s.Acc)
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
