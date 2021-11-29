package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/golang-collections/collections/set"
)

type group []*set.Set

func parseInput(s string) []group {
	var ret []group
	var cur group
	for _, l := range strings.Split(s, "\n") {
		if l == "" {
			ret = append(ret, cur)
			cur = nil
			continue
		}

		q := set.New()
		for _, c := range l {
			q.Insert(c)
		}
		cur = append(cur, q)
	}
	return ret
}

func sumGroupsDistinct(gs []group) int {
	var i int
	for _, g := range gs {
		s := set.New()
		for _, q := range g {
			s = s.Union(q)
		}
		i += s.Len()
	}
	return i
}

func runPartOne() error {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		return err
	}
	gs := parseInput(string(content))
	fmt.Println(sumGroupsDistinct(gs))
	return nil
}

func sumGroupsIntersect(gs []group) int {
	var i int
	for _, g := range gs {
		if len(g) == 0 {
			continue
		}
		s := g[0]
		for i := 1; i < len(g); i++ {
			s = s.Intersection(g[i])
		}
		i += s.Len()
	}
	return i
}

func runPartTwo() error {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		return err
	}
	gs := parseInput(string(content))
	fmt.Println(sumGroupsIntersect(gs))
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
