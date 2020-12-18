package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type expr interface {
	Value() int
}

type operationType int

const (
	opAddition operationType = 1
	opMultiply operationType = 2
)

type binaryOp struct {
	op          operationType
	left, right expr
}

func (o binaryOp) String() string {
	var c string
	switch o.op {
	case opAddition:
		c = "+"
	case opMultiply:
		c = "*"
	default:
		panic("bad operator")
	}
	return fmt.Sprintf("(%s %s %s)", o.left, c, o.right)
}

func (o *binaryOp) Value() int {
	switch o.op {
	case opAddition:
		return o.left.Value() + o.right.Value()
	case opMultiply:
		return o.left.Value() * o.right.Value()
	}
	panic("bad operator")
}

type parenthesis struct {
	child expr
}

func (p parenthesis) String() string {
	return fmt.Sprintf("(%s)", p.child)
}

func (p *parenthesis) Value() int {
	return p.child.Value()
}

type literal int

func (l literal) String() string {
	return strconv.Itoa(int(l))
}

func (l *literal) Value() int {
	return int(*l)
}

func parseExpression(s string) (expr, error) {
	var cur expr
	var stack []expr

	for _, c := range s {
		switch c {
		case ' ':
			continue
		case '+':
			cur = &binaryOp{left: cur, op: opAddition}
		case '*':
			cur = &binaryOp{left: cur, op: opMultiply}
		case '(':
			stack = append([]expr{cur}, stack...)
			cur = nil
		case ')':
			top := stack[0]
			stack = stack[1:]
			cur = &parenthesis{child: cur}
			if top != nil {
				bop, ok := top.(*binaryOp)
				if !ok {
					return nil, errors.New("invalid stack element")
				}
				bop.right = cur
				cur = top
			}
		default:
			i, err := strconv.Atoi(string(c))
			if err != nil {
				return nil, err
			}
			l := literal(i)
			if cur == nil {
				cur = &l
			} else {
				bop, ok := cur.(*binaryOp)
				if !ok {
					return nil, fmt.Errorf("ended with invalid expression '%s'", string(c))
				}
				bop.right = &l
			}
		}
	}
	if len(stack) > 0 {
		return nil, errors.New("expected ')'")
	}
	return cur, nil
}

func parseInput(content string) []string {
	var ret []string
	for _, l := range strings.Split(string(content), "\n") {
		if l == "" {
			continue
		}
		ret = append(ret, l)
	}
	return ret
}

func runPartOne() error {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		return err
	}
	exprs := parseInput(string(content))
	var sum int
	for _, exp := range exprs {
		e, err := parseExpression(exp)
		if err != nil {
			return err
		}
		sum += e.Value()
	}
	fmt.Println(sum)
	return nil
}

// maybeApplyPrecedence returns the new parent expression
func maybeApplyPrecedence(e expr) (expr, bool) {
	plus, ok := e.(*binaryOp)
	if !ok || plus.op != opAddition {
		return e, false
	}

	var changed bool

	left, ok := plus.left.(*binaryOp)
	if ok && left.op == opMultiply {
		e = left
		plus.left = left.right
		left.right = plus
		changed = true
	}
	right, ok := plus.right.(*binaryOp)
	if ok && right.op == opMultiply {
		if e == plus {
			e = right
		}
		plus.right = right.left
		right.left = plus
		changed = true
	}
	return e, changed
}

func manipulatePrecedence(e expr) expr {
	// repeating this loop until no more changes
	// hack to deal with recursive tree shifts
	// solution is probably to use parent pointers
	for {
		// use a root pointer that wont be swapped
		root := &parenthesis{child: e}
		toVisit := []expr{root}

		var anyChanged bool

		for len(toVisit) > 0 {
			e := toVisit[0]
			toVisit = toVisit[1:]

			switch o := e.(type) {
			case *parenthesis:
				var a bool
				o.child, a = maybeApplyPrecedence(o.child)
				toVisit = append(toVisit, o.child)
				if a {
					anyChanged = true
				}
			case *binaryOp:
				var a, b bool
				o.left, a = maybeApplyPrecedence(o.left)
				o.right, b = maybeApplyPrecedence(o.right)
				toVisit = append(toVisit, o.left, o.right)
				if a || b {
					anyChanged = true
				}
			}
		}
		e = root.child
		if !anyChanged {
			break
		}
	}
	return e
}

func runPartTwo() error {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		return err
	}
	exprs := parseInput(string(content))
	var sum int
	for _, exp := range exprs {
		e, err := parseExpression(exp)
		if err != nil {
			return err
		}
		e = manipulatePrecedence(e)
		sum += e.Value()
	}
	fmt.Println(sum)
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
