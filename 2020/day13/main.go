package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

func parseInput(content string) (int, []int, error) {
	lines := strings.Split(content, "\n")
	i, err := strconv.Atoi(lines[0])
	if err != nil {
		return 0, nil, err
	}
	var times []int
	for _, t := range strings.Split(lines[1], ",") {
		var j int
		if t != "x" {
			j, err = strconv.Atoi(t)
			if err != nil {
				return 0, nil, err
			}
		}
		times = append(times, j)
	}
	return i, times, nil
}

func getNextBus(i int, times []int) (int, int) {
	var minTTW, busNo int
	for _, t := range times {
		if t == 0 {
			continue
		}
		timeToWait := t - (i % t)
		if timeToWait < minTTW || minTTW == 0 {
			minTTW = timeToWait
			busNo = t
		}
	}
	return minTTW, busNo
}

func runPartOne() error {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		return err
	}
	i, l, err := parseInput(string(content))
	if err != nil {
		return err
	}
	ttw, bus := getNextBus(i, l)
	fmt.Println(ttw * bus)
	return nil
}

func eGCD(a, b int) (int, int, int) {
	oldR, r := a, b
	oldS, s := 1, 0
	oldT, t := 0, 1

	for r != 0 {
		q := oldR / r
		oldR, r = r, oldR-q*r
		oldS, s = s, oldS-q*s
		oldT, t = t, oldT-q*t
	}
	return oldS, oldT, oldR
}

type modAndRemainder struct {
	Mod, Remain int
}

func getModAndRemainders(buses []int) []modAndRemainder {
	var ret []modAndRemainder
	for i, b := range buses {
		if b == 0 {
			continue
		}
		ret = append(ret, modAndRemainder{Mod: b, Remain: i})
	}
	// Sort by highest modulo
	sort.Slice(ret, func(i, j int) bool {
		return ret[i].Mod > ret[j].Mod
	})
	return ret
}

func chineseRemainder(mnrs []modAndRemainder) modAndRemainder {
	cur := mnrs[0]
	for _, mr := range mnrs[1:] {
		rem := mr.Remain - cur.Remain
		a, _, _ := eGCD(cur.Mod, mr.Mod)
		rem *= a
		if rem < 0 {
			rem = mr.Mod - rem
		}
		rem %= mr.Mod
		cur = modAndRemainder{
			Mod:    cur.Mod * mr.Mod,
			Remain: cur.Mod * rem + cur.Remain,
		}
	}
	return cur
}

func runPartTwo() error {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		return err
	}
	_, buses, err := parseInput(string(content))
	if err != nil {
		return err
	}
	mnrs := getModAndRemainders(buses)
	mr := chineseRemainder(mnrs)
	fmt.Println(mr.Mod - mr.Remain)
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
