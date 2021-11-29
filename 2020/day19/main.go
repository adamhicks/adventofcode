package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

type cnfGrammar struct {
	rules          []nonTermRule
	terminals      []termRule
	maxNonTerminal int
}

func (g cnfGrammar) Normalise() cnfGrammar {
	sort.Slice(g.rules, func(i, j int) bool {
		return g.rules[i].From < g.rules[j].From
	})
	sort.Slice(g.terminals, func(i, j int) bool {
		return g.terminals[i].From < g.terminals[j].From
	})
	return g
}

func (g cnfGrammar) NonTerminals() []int {
	uniq := make(map[int]bool)
	for _, r := range g.rules {
		uniq[r.From] = true
	}
	for _, r := range g.terminals {
		uniq[r.From] = true
	}
	ret := make([]int, 0, len(uniq))
	for i := range uniq {
		ret = append(ret, i)
	}
	sort.Ints(ret)
	return ret
}

func (g cnfGrammar) Minimalise() cnfGrammar {
	lookup := make(map[int]int)
	for i, nt := range g.NonTerminals() {
		lookup[nt] = i
	}

	rules := make([]nonTermRule, 0, len(g.rules))
	for _, r := range g.rules {
		rules = append(rules, nonTermRule{
			From:  lookup[r.From],
			ToOne: lookup[r.ToOne],
			ToTwo: lookup[r.ToTwo],
		})
	}
	terms := make([]termRule, 0, len(g.terminals))
	for _, r := range g.terminals {
		terms = append(terms, termRule{
			From: lookup[r.From],
			To:   r.To,
		})
	}
	return cnfGrammar{
		rules:          rules,
		terminals:      terms,
		maxNonTerminal: len(lookup),
	}.Normalise()
}

func (g cnfGrammar) String() string {
	var b strings.Builder
	for _, r := range g.rules {
		fmt.Fprintln(&b, r)
	}
	for _, r := range g.terminals {
		fmt.Fprintln(&b, r)
	}
	return b.String()
}

func (g cnfGrammar) Matches(str string) bool {
	n := len(str)
	if n == 0 {
		return false
	}

	cyk := make([][][]bool, n)
	for i := range cyk {
		cyk2 := make([][]bool, n)
		for j := range cyk2 {
			cyk2[j] = make([]bool, g.maxNonTerminal)
		}
		cyk[i] = cyk2
	}

	for s := 1; s <= n; s++ {
		for _, r := range g.terminals {
			if r.To == str[s-1:s] {
				cyk[0][s-1][r.From] = true
			}
		}
	}

	for l := 2; l <= n; l++ {
		for s := 1; s <= n-l+1; s++ {
			for p := 1; p <= l-1; p++ {
				for _, r := range g.rules {
					if cyk[p-1][s-1][r.ToOne] && cyk[l-p-1][s+p-1][r.ToTwo] {
						cyk[l-1][s-1][r.From] = true
					}
				}
			}
		}
	}

	return cyk[n-1][0][0]
}

type nonTermRule struct {
	From  int
	ToOne int
	ToTwo int
}

func (r nonTermRule) String() string {
	return fmt.Sprintf("%d -> %d %d", r.From, r.ToOne, r.ToTwo)
}

type termRule struct {
	From int
	To   string
}

func (r termRule) String() string {
	return fmt.Sprintf("%d -> %s", r.From, r.To)
}

func parseInput(content string) (map[int]string, []string, error) {
	parts := strings.Split(content, "\n\n")

	rules := make(map[int]string)

	for _, s := range strings.Split(parts[0], "\n") {
		if s == "" {
			continue
		}
		ruleParts := strings.Split(s, ": ")
		srcStr := ruleParts[0]
		srcNum, err := strconv.Atoi(srcStr)
		if err != nil {
			return nil, nil, err
		}
		rules[srcNum] = ruleParts[1]
	}

	var tests []string
	for _, s := range strings.Split(parts[1], "\n") {
		if s == "" {
			continue
		}
		tests = append(tests, s)
	}
	return rules, tests, nil
}

func parseTerminal(s string) string {
	return strings.ReplaceAll(s, "\"", "")
}

func parseProduction(s string) ([][]int, error) {
	var ret [][]int
	for _, r := range strings.Split(s, " | ") {
		var prod []int
		for _, iStr := range strings.Split(r, " ") {
			i, err := strconv.Atoi(iStr)
			if err != nil {
				return nil, err
			}
			prod = append(prod, i)
		}
		ret = append(ret, prod)
	}
	return ret, nil
}

func parseGrammar(rulesIn map[int]string) (cnfGrammar, error) {
	unitRules := make(map[int][]int)
	var terms []termRule
	var rules []nonTermRule
	maxNT := 200

	for src, s := range rulesIn {
		if strings.Contains(s, "\"") {
			terms = append(terms, termRule{From: src, To: parseTerminal(s)})
			continue
		}
		prods, err := parseProduction(s)
		if err != nil {
			return cnfGrammar{}, err
		}
		for _, prod := range prods {
			switch len(prod) {
			case 1:
				unitRules[src] = append(unitRules[src], prod[0])
			case 2:
				rules = append(rules,
					nonTermRule{From: src, ToOne: prod[0], ToTwo: prod[1]},
				)
			case 3:
				newRule := maxNT
				maxNT++
				rules = append(rules,
					nonTermRule{From: src, ToOne: prod[0], ToTwo: newRule},
					nonTermRule{From: newRule, ToOne: prod[1], ToTwo: prod[2]},
				)
			default:
				panic("invalid production")
			}
		}
	}

	for src, prods := range unitRules {
		for _, p := range prods {
			for _, r := range rules {
				if r.From == p {
					rules = append(rules,
						nonTermRule{From: src, ToOne: r.ToOne, ToTwo: r.ToTwo},
					)
				}
			}
			for _, r := range terms {
				if r.From == p {
					terms = append(terms,
						termRule{From: src, To: r.To},
					)
				}
			}
		}
	}

	ret := cnfGrammar{
		rules:          rules,
		maxNonTerminal: maxNT,
		terminals:      terms,
	}

	return ret.Normalise(), nil
}

func countMatches(g cnfGrammar, tests []string) int {
	var s int
	for _, t := range tests {
		if g.Matches(t) {
			s++
		}
	}
	return s
}

func runPartOne() error {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		return err
	}
	rules, tests, err := parseInput(string(content))
	if err != nil {
		return err
	}
	g, err := parseGrammar(rules)
	if err != nil {
		return err
	}
	g = g.Minimalise()
	sum := countMatches(g, tests)
	fmt.Println(sum)
	return nil
}

func replaceRules(rules map[int]string) {
	rules[8] = "42 | 42 8"
	rules[11] = "42 31 | 42 11 31"
}

func runPartTwo() error {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		return err
	}
	rules, tests, err := parseInput(string(content))
	if err != nil {
		return err
	}
	replaceRules(rules)
	g, err := parseGrammar(rules)
	if err != nil {
		return err
	}
	g = g.Minimalise()
	sum := countMatches(g, tests)
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
