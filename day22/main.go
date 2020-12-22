package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func parseInput(content string) ([][]int, error) {
	var ret [][]int
	for _, player := range strings.Split(content, "\n\n") {
		var cards []int
		for _, l := range strings.Split(player, "\n")[1:] {
			if l == "" {
				continue
			}
			i, err := strconv.Atoi(l)
			if err != nil {
				return nil, err
			}
			cards = append(cards, i)
		}
		ret = append(ret, cards)
	}
	return ret, nil
}

func playGame(deckA, deckB []int) []int {
	for len(deckA) > 0 && len(deckB) > 0 {
		a, b := deckA[0], deckB[0]
		if a > b {
			deckA = append(deckA[1:], a, b)
			deckB = deckB[1:]
		} else {
			deckA = deckA[1:]
			deckB = append(deckB[1:], b, a)
		}
	}
	if len(deckA) > 0 {
		return deckA
	}
	return deckB
}

func score(deck []int) int {
	var sum int
	n := len(deck)
	for i, v := range deck {
		sum += (n - i) * v
	}
	return sum
}

func runPartOne() error {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		return err
	}
	decks, err := parseInput(string(content))
	if err != nil {
		return err
	}
	if len(decks) != 2 {
		return errors.New("wrong number of decks")
	}
	s := score(playGame(decks[0], decks[1]))
	fmt.Println(s)
	return nil
}

func playRecursiveGame(deckA, deckB []int) (int, []int) {
	seen := make(map[int]bool)

	for len(deckA) > 0 && len(deckB) > 0 {
		hash := (score(deckA) << 32) | score(deckB)
		if seen[hash] {
			return 1, nil
		}
		seen[hash] = true

		a, b := deckA[0], deckB[0]
		deckA = deckA[1:]
		deckB = deckB[1:]

		var winner int
		if a <= len(deckA) && b <= len(deckB) {
			copyA := make([]int, a)
			copy(copyA, deckA)
			copyB := make([]int, b)
			copy(copyB, deckB)
			winner, _ = playRecursiveGame(copyA, copyB)
		} else if a > b {
			winner = 1
		} else {
			winner = 2
		}

		if winner == 1 {
			deckA = append(deckA, a, b)
		} else {
			deckB = append(deckB, b, a)
		}
	}
	if len(deckA) > 0 {
		return 1, deckA
	}
	return 2, deckB
}

func runPartTwo() error {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		return err
	}
	decks, err := parseInput(string(content))
	if err != nil {
		return err
	}
	if len(decks) != 2 {
		return errors.New("wrong number of decks")
	}
	_, res := playRecursiveGame(decks[0], decks[1])
	fmt.Println(score(res))
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
