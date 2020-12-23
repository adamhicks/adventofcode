package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func parseInput(content string) ([]int, error) {
	var ret []int
	for _, c := range content {
		i, err := strconv.Atoi(string(c))
		if err != nil {
			return nil, err
		}
		ret = append(ret, i)
	}
	return ret, nil
}

type element struct {
	v    int
	next *element
}

type list struct {
	cur   *element
	max   int
	index map[int]*element
}

func createList(s []int) list {
	if len(s) == 0 {
		return list{}
	}
	first := &element{v: s[0]}
	index := map[int]*element{first.v: first}
	max := first.v

	cur := first
	for i := 1; i < len(s); i++ {
		this := &element{v: s[i]}
		cur.next = this
		cur = this

		index[this.v] = this
		if this.v > max {
			max = this.v
		}
	}
	cur.next = first

	return list{
		cur:   first,
		max:   max,
		index: index,
	}
}

func nextID(i, max int) int {
	if i < 1 {
		panic("invalid id")
	}
	if i == 1 {
		return max
	}
	return i - 1
}

func (l *list) move() {
	start := l.cur.next
	end := l.cur

	vals := make(map[int]bool)
	for i := 0; i < 3; i++ {
		end = end.next
		vals[end.v] = true
	}

	t := nextID(l.cur.v, l.max)
	for vals[t] {
		t = nextID(t, l.max)
	}

	target := l.index[t]

	l.cur.next = end.next
	end.next = target.next
	target.next = start

	l.cur = l.cur.next
}

func (l list) listNAfter(start int, n int) []int {
	e, ok := l.index[start]
	if !ok {
		panic("unknown starting point")
	}
	e = e.next

	var ret []int
	for i := 0; i < n; i++ {
		ret = append(ret, e.v)
		e = e.next
	}
	return ret
}

func (l list) ToArray() []int {
	e := l.cur
	ret := []int{e.v}
	cur := e.next
	for cur != e {
		ret = append(ret, cur.v)
		cur = cur.next
	}
	return ret
}

func getPartOneOutput(l list) string {
	var b strings.Builder
	for _, i := range l.listNAfter(1, 8) {
		fmt.Fprint(&b, i)
	}
	return b.String()
}

func runPartOne() error {
	vals, err := parseInput("952316487")
	if err != nil {
		return err
	}
	l := createList(vals)
	for i := 0; i < 100; i++ {
		l.move()
	}
	fmt.Println(getPartOneOutput(l))
	return nil
}

func (l *list) fillTo(to int) {
	cur := l.cur
	for cur.next != l.cur {
		cur = cur.next
	}
	for i := l.max + 1; i <= to; i++ {
		newEle := &element{v: i}
		cur.next = newEle
		cur = newEle
		l.index[i] = newEle
	}
	cur.next = l.cur

	l.max = to
}

func getPartTwoOutput(l list) int {
	mult := 1
	for _, v := range l.listNAfter(1, 2) {
		mult *= v
	}
	return mult
}

func runPartTwo() error {
	vals, err := parseInput("952316487")
	if err != nil {
		return err
	}
	l := createList(vals)
	l.fillTo(1000000)
	for i := 0; i < 10000000; i++ {
		l.move()
	}
	fmt.Println(getPartTwoOutput(l))
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
