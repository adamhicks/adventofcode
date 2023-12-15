package day15

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var inputString string

type input []string

func parseInput(s string) input {
	return strings.Split(strings.TrimSpace(s), ",")
}

var testString1 = `rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7`

func hash(s string) int {
	var ret int
	for _, r := range s {
		ret += int(r)
		ret *= 17
		ret %= 256
	}
	return ret
}

func runPartOne(s input) error {
	var sum int
	for _, h := range s {
		sum += hash(h)
	}
	fmt.Println(sum)
	return nil
}

var testString2 = testString1

type lens struct {
	Label, Focal string
}

func remove(box []lens, label string) []lens {
	i := slices.IndexFunc(box, func(l lens) bool {
		return l.Label == label
	})
	if i < 0 {
		return box
	}
	copy(box[i:], box[i+1:])
	return box[:len(box)-1]
}

func add(newLens lens, box []lens) []lens {
	i := slices.IndexFunc(box, func(l lens) bool {
		return l.Label == newLens.Label
	})
	if i < 0 {
		return append(box, newLens)
	}
	box[i] = newLens
	return box
}

func hashmap(s input) [256][]lens {
	var out [256][]lens
	for _, l := range s {
		if strings.HasSuffix(l, "-") {
			label := l[:len(l)-1]
			v := hash(label)
			out[v] = remove(out[v], label)
		} else {
			label, focal, _ := strings.Cut(l, "=")
			v := hash(label)
			out[v] = add(lens{Label: label, Focal: focal}, out[v])
		}
	}
	return out
}

func runPartTwo(s input) error {
	hm := hashmap(s)
	var sum int
	for boxIdx, b := range hm {
		for lensIdx, l := range b {
			fl, err := strconv.Atoi(l.Focal)
			if err != nil {
				return err
			}
			sum += (boxIdx + 1) * (lensIdx + 1) * fl
		}
	}
	fmt.Println(sum)
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
