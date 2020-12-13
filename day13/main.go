package main

import (
	"fmt"
	"io/ioutil"
	"log"
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

func runPartTwo() error {
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
