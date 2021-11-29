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
	SetParent(expr)
	ChangeChild(from, to expr)
	Parent() expr
	Format() string
	Value() int
}

type basicExpr struct {
	parent expr
}

func (b *basicExpr) SetParent(e expr) {
	b.parent = e
}

func (b basicExpr) Parent() expr {
	return b.parent
}

type operationType int

const (
	opAddition operationType = 1
	opMultiply operationType = 2
)

type binaryOp struct {
	basicExpr
	op          operationType
	left, right expr
}

func (o *binaryOp) ChangeChild(from, to expr) {
	if o.left == from {
		o.left = to
	} else if o.right == from {
		o.right = to
	} else {
		panic("invalid child")
	}
}

func (o binaryOp) String() string {
	switch o.op {
	case opAddition:
		return "bop +"
	case opMultiply:
		return "bop *"
	}
	panic("invalid op")
}

func (o binaryOp) Format() string {
	var c string
	switch o.op {
	case opAddition:
		c = "+"
	case opMultiply:
		c = "*"
	default:
		panic("bad operator")
	}
	return fmt.Sprintf("%s %s %s", o.left.Format(), c, o.right.Format())
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
	basicExpr
	parent expr
	child  expr
}

func (p *parenthesis) ChangeChild(from, to expr) {
	if p.child == from {
		p.child = to
	} else {
		panic("invalid child")
	}
}

func (p parenthesis) String() string {
	return "par"
}

func (p parenthesis) Format() string {
	return "(" + p.child.Format() + ")"
}

func (p *parenthesis) Value() int {
	return p.child.Value()
}

type literal struct {
	basicExpr
	val int
}

func (l literal) ChangeChild(from, to expr) {
	panic("no children")
}

func (l literal) String() string {
	return "literal " + l.Format()
}

func (l literal) Format() string {
	return strconv.Itoa(l.val)
}

func (l *literal) Value() int {
	return l.val
}

func parseExpression(s string) (expr, error) {
	var cur expr
	var stack []expr

	for _, c := range s {
		switch c {
		case ' ':
			continue
		case '+':
			bop := &binaryOp{left: cur, op: opAddition}
			cur.SetParent(bop)
			cur = bop
		case '*':
			bop := &binaryOp{left: cur, op: opMultiply}
			cur.SetParent(bop)
			cur = bop
		case '(':
			stack = append([]expr{cur}, stack...)
			cur = nil
		case ')':
			p := &parenthesis{child: cur}
			cur.SetParent(p)
			cur = p

			top := stack[0]
			stack = stack[1:]
			if top != nil {
				bop, ok := top.(*binaryOp)
				if !ok {
					return nil, errors.New("invalid stack element")
				}
				bop.right = cur
				p.SetParent(bop)
				cur = top
			}
		default:
			i, err := strconv.Atoi(string(c))
			if err != nil {
				return nil, err
			}
			l := &literal{val: i}
			if cur == nil {
				cur = l
			} else {
				bop, ok := cur.(*binaryOp)
				if !ok {
					return nil, fmt.Errorf("ended with invalid expression '%s'", string(c))
				}
				bop.right = l
				l.SetParent(bop)
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
		left.SetParent(plus.Parent())
		plus.SetParent(left)

		plus.left = left.right
		left.right = plus
		changed = true
	}
	right, ok := plus.right.(*binaryOp)
	if ok && right.op == opMultiply {
		if e == plus {
			e = right
		}
		right.SetParent(plus.Parent())
		plus.SetParent(right)

		plus.right = right.left
		right.left = plus
		changed = true
	}
	return e, changed
}

func manipulatePrecedence(e expr) expr {
	// use a root pointer that wont be swapped
	root := &parenthesis{child: e}
	e.SetParent(root)
	toVisit := []expr{e}

	for len(toVisit) > 0 {
		e := toVisit[0]
		toVisit = toVisit[1:]

		p := e.Parent()
		ne, changed := maybeApplyPrecedence(e)
		if changed {
			p.ChangeChild(e, ne)
			toVisit = append(toVisit, p)
		} else {
			switch o := e.(type) {
			case *parenthesis:
				toVisit = append(toVisit, o.child)
			case *binaryOp:
				toVisit = append(toVisit, o.left, o.right)
			}
		}
	}
	e = root.child
	e.SetParent(nil)
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
