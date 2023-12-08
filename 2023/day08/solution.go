package day08

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var inputString string

func parseInput(s string) (string, []string) {
	dir, rest, _ := strings.Cut(strings.TrimSpace(s), "\n\n")
	return dir, strings.Split(rest, "\n")
}

var testString1 = `LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)
`

type Node struct {
	Name        string
	Left, Right *Node
}

func getOrCreate(m map[string]*Node, name string) *Node {
	node, ok := m[name]
	if ok {
		return node
	}
	node = &Node{Name: name}
	m[name] = node
	return node
}

func parseNodes(nodes []string) map[string]*Node {
	ret := make(map[string]*Node)
	for _, n := range nodes {
		name, children, _ := strings.Cut(n, " = ")
		parent := getOrCreate(ret, name)
		leftName, rightName, _ := strings.Cut(strings.Trim(children, "()"), ", ")
		left := getOrCreate(ret, leftName)
		parent.Left = left
		right := getOrCreate(ret, rightName)
		parent.Right = right
	}
	return ret
}

func follow(from string, done func(*Node) bool, dirs string, tree map[string]*Node) int {
	current := tree[from]
	var steps int
	for !done(current) {
		idx := steps % len(dirs)
		switch dirs[idx] {
		case 'L':
			current = current.Left
		case 'R':
			current = current.Right
		}
		steps++
	}
	return steps
}

func runPartOne(dirs string, nodes []string) error {
	done := func(n *Node) bool {
		return n.Name == "ZZZ"
	}
	fmt.Println(follow("AAA", done, dirs, parseNodes(nodes)))
	return nil
}

var testString2 = `LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)
`

func findEndCycle(start string, dirs string, tree map[string]*Node) int {
	seen := make(map[string]int)
	var cycle int
	var counter int
	done := func(n *Node) bool {
		defer func() { counter++ }()
		if n.Name[2] != 'Z' {
			return false
		}
		if _, ok := seen[n.Name]; ok {
			cycle = counter - seen[n.Name]
			return true
		} else {
			seen[n.Name] = counter
		}
		return false
	}
	follow(start, done, dirs, tree)
	return cycle
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

func runPartTwo(dirs string, nodes []string) error {
	var start []string
	tree := parseNodes(nodes)
	for n := range tree {
		if n[2] == 'A' {
			start = append(start, n)
		}
	}
	steps := 1
	for _, s := range start {
		cycle := findEndCycle(s, dirs, tree)
		steps = lcm(steps, cycle)
	}
	fmt.Println(steps)
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
