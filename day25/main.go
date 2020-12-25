package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func parseInput(content string) (int, int, error) {
	lines := strings.Split(content, "\n")
	i1, err := strconv.Atoi(lines[0])
	if err != nil {
		return 0, 0, err
	}
	i2, err := strconv.Atoi(lines[1])
	if err != nil {
		return 0, 0, err
	}
	return i1, i2, nil
}

func findMultiple(s, mod, target int) int {
	v := s
	i := 1
	for v != target {
		v = (v * s) % mod
		i++
	}
	return i
}

func loopN(s, mod, n int) int {
	v := s
	for i := 1; i < n; i++ {
		v = (v * s) % mod
	}
	return v
}

func runPartOne() error {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		return err
	}
	card, door, err := parseInput(string(content))
	if err != nil {
		return err
	}
	cSecret := findMultiple(7, 20201227, card)
	key := loopN(door, 20201227, cSecret)
	fmt.Println(key)
	return nil
}

func main() {
	if err := runPartOne(); err != nil {
		log.Fatal(err)
	}
}
