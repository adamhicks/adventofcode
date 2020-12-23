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

func (l *list) move() {
	start := l.cur.next
	end := start.next.next

	t := l.cur.v
	for {
		t--
		if t < 1 {
			t = l.max
		}
		if t != start.v && t != start.next.v && t != end.v {
			break
		}
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
	e := l.listNAfter(1, 2)
	return e[0] * e[1]
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
