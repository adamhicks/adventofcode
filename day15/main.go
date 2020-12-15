package main

import (
	"fmt"
	"log"
)

func runSequence(starting []int, n int) int {
	lastSeen := make(map[int]int)
	for i, s := range starting[:len(starting)-1] {
		lastSeen[s] = i
	}
	cur := starting[len(starting)-1]
	for i := len(starting) - 1; i < n-1; i++ {
		var nxt int
		ls, ok := lastSeen[cur]
		if ok {
			nxt = i - ls
		}
		lastSeen[cur] = i
		cur = nxt
	}
	return cur
}

func runPartOne() error {
	seq := []int{0, 20, 7, 16, 1, 18, 15}
	i := runSequence(seq, 2020)
	fmt.Println(i)
	return nil
}

func runPartTwo() error {
	seq := []int{0, 20, 7, 16, 1, 18, 15}
	i := runSequence(seq, 30000000)
	fmt.Println(i)
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
