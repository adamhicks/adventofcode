package day25

import (
	_ "embed"
	"fmt"
	"sort"
	"strings"
)

//go:embed input.txt
var inputString string

type input []string

func parseInput(s string) input {
	return strings.Split(strings.TrimSpace(s), "\n")
}

var testString1 = `jqt: rhn xhk nvd
rsh: frs pzl lsr
xhk: hfx
cmg: qnr nvd lhk bvb
rhn: xhk bvb hfx
bvb: xhk hfx
pzl: lsr hfx nvd
qnr: nvd
ntq: jqt hfx bvb xhk
nvd: lhk
lsr: lhk
rzs: qnr cmg lsr rsh
frs: qnr lhk lsr`

func routesBetween(from, to string, nodes map[string][]string, ignore map[wire]bool) [][]string {
	var ret [][]string
	seen := map[string]bool{from: true}
	q := [][]string{{from}}
	for len(q) > 0 {
		p := q[0]
		q = q[1:]
		last := p[len(p)-1]
		if last == to {
			ret = append(ret, p)
			continue
		}
		for _, n := range nodes[last] {
			if ignore[wire{last, n}] || ignore[wire{n, last}] {
				continue
			}
			if seen[n] && n != to {
				continue
			}
			seen[n] = true
			p1 := make([]string, len(p), len(p)+1)
			copy(p1, p)
			p1 = append(p1, n)
			q = append(q, p1)
		}
	}
	return ret
}

//func followWires(wires [][2]string, ignore map[int]bool) (map[string]int, int) {
//	var nxtColour int
//	sets := make(map[string]int)
//
//	for i, w := range wires {
//		if ignore[i] {
//			continue
//		}
//		if c, ok := sets[w[0]]; ok {
//			sets[w[1]] = c
//		} else if c, ok := sets[w[1]]; ok {
//			sets[w[0]] = c
//		} else {
//			nxtColour++
//			sets[w[0]] = nxtColour
//			sets[w[1]] = nxtColour
//		}
//	}
//	return sets, nxtColour
//}

func parseWires(s input) []wire {
	var wires []wire
	for _, l := range s {
		from, tos, _ := strings.Cut(l, ": ")
		for _, to := range strings.Fields(tos) {
			wires = append(wires, [...]string{from, to})
		}
	}
	return wires
}

func score(s map[string]int) int {
	sCount := make(map[int]int)
	for _, v := range s {
		sCount[v]++
	}
	ret := 1
	for _, count := range sCount {
		ret *= count
	}
	return ret
}

type wire [2]string

func getKeyWires(wires []wire, nodes map[string][]string, ignore map[wire]bool) []wire {
	v := make(map[wire]int)
	for _, w := range wires {
		if ignore[w] {
			continue
		}
		from, to := w[0], w[1]
		routes := routesBetween(from, to, nodes, ignore)
		d := make(map[int]int)
		for _, r := range routes {
			d[len(r)]++
		}
		if len(d) == 1 {
			fmt.Println("found breaker", d)
		}
		if d[2] == 1 && len(d) <= 2 {
			for _, dist := range d {
				if dist > 3 {
					v[w] = dist
					break
				}
			}
		}
	}
	ret := make([]wire, 0, len(v))
	for k := range v {
		ret = append(ret, k)
	}
	sort.Slice(ret, func(i, j int) bool {
		wi, wj := ret[i], ret[j]
		return v[wi] > v[wj]
	})
	return ret
}

func someKey[K comparable, V any](m map[K]V) K {
	var ret K
	for k := range m {
		ret = k
		break
	}
	return ret
}

func distinct(nodes map[string][]string, ignore map[wire]bool) ([2][]string, bool) {
	start := someKey(nodes)
	seen := map[string]bool{start: true}
	q := []string{start}
	for len(q) > 0 {
		nxt := q[0]
		q = q[1:]
		for _, n := range nodes[nxt] {
			if seen[n] {
				continue
			}
			if ignore[wire{nxt, n}] || ignore[wire{n, nxt}] {
				continue
			}
			seen[n] = true
			q = append(q, n)
		}
	}
	if len(seen) == len(nodes) {
		return [2][]string{}, false
	}
	var a, b []string
	for n := range nodes {
		if seen[n] {
			a = append(a, n)
		} else {
			b = append(b, n)
		}
	}
	return [2][]string{a, b}, true
}

func findTripleDisconnect(wires []wire) ([3]wire, bool) {
	node := make(map[string][]string)
	for _, w := range wires {
		node[w[0]] = append(node[w[0]], w[1])
		node[w[1]] = append(node[w[1]], w[0])
	}

	keyWires := getKeyWires(wires, node, nil)
	for i := 0; i < len(keyWires)-1; i++ {
		for j := i + 1; j < len(keyWires); j++ {
			a, b := keyWires[i], keyWires[j]
			for _, w := range wires {
				if w == a || w == b {
					continue
				}
				ignore := map[wire]bool{a: true, b: true, w: true}
				_, ok := distinct(node, ignore)
				if ok {
					return [...]wire{a, b, w}, true
				}
			}
			fmt.Println(a, b, "no breaker")
		}
	}
	return [3]wire{}, false
}

func runPartOne(s input) error {
	wires := parseWires(s)
	trip, found := findTripleDisconnect(wires)
	if !found {
		panic("no disconnect found")
	}
	fmt.Println(trip)
	return nil
}

var testString2 = testString1

func runPartTwo(s input) error {
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
