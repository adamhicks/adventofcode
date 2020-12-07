package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type grammar map[string]map[string]int

func (g grammar) String() string {
	var b strings.Builder
	for from, produces := range g {
		fmt.Fprintf(&b, "%s -> ", from)
		if len(produces) == 0 {
			fmt.Fprintf(&b, "∅")
		} else {
			outs := make([]string, 0, len(produces))
			for to, n := range produces {
				outs = append(outs, strconv.Itoa(n)+" "+to)
			}
			fmt.Fprintf(&b, strings.Join(outs, " + "))
		}
		fmt.Fprintf(&b, "\n")
	}
	return b.String()
}

func parseInput(content string) (grammar, error) {
	re, err := regexp.Compile("(\\d+) (\\w+ \\w+)")
	if err != nil {
		return nil, err
	}

	g := make(grammar)

	for _, l := range strings.Split(content, "\n") {
		if l == "" {
			continue
		}
		parts := strings.SplitN(l, " ", 3)

		symbol := parts[0] + " " + parts[1]
		rule := make(map[string]int)

		for _, match := range re.FindAllStringSubmatch(parts[2], -1) {
			n, err := strconv.Atoi(match[1])
			if err != nil {
				return grammar{}, err
			}
			rule[match[2]] = n
		}
		g[symbol] = rule
	}
	return g, nil
}

func reverse(g grammar) grammar {
	rev := make(grammar)

	for from, produce := range g {
		for to := range produce {
			m, ok := rev[to]
			if !ok {
				m = make(map[string]int)
				rev[to] = m
			}
			m[from] = 1
		}
	}
	return rev
}

func findAllContaining(start string, g grammar) map[string]bool {
	visited := make(map[string]bool)

	toVisit := []string{start}
	for len(toVisit) > 0 {
		cur := toVisit[0]
		toVisit = toVisit[1:]

		visited[cur] = true
		for nxt := range g[cur] {
			if !visited[nxt] {
				toVisit = append(toVisit, nxt)
			}
		}
	}

	delete(visited, start)
	return visited
}

func runPartOne() error {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		return err
	}
	g, err := parseInput(string(content))
	if err != nil {
		return err
	}
	g = reverse(g)
	connected := findAllContaining("shiny gold", g)
	fmt.Println(len(connected))
	return nil
}

func sumContents(start string, g grammar) int {
	var sum int
	type visit struct {
		s string
		n int
	}
	toVisit := []visit{{start, 1}}

	for len(toVisit) > 0 {
		cur := toVisit[0]
		toVisit = toVisit[1:]

		sum += cur.n
		for nxt, m := range g[cur.s] {
			toVisit = append(toVisit, visit{
				s: nxt, n: cur.n * m,
			})
		}
	}
	// Remove the original bag
	return sum - 1
}

func runPartTwo() error {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		return err
	}
	g, err := parseInput(string(content))
	if err != nil {
		return err
	}
	sum := sumContents("shiny gold", g)
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
