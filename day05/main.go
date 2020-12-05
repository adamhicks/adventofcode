package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func parseAddress(s string) int {
	var a int
	for _, r := range s {
		a <<= 1
		if r == 'B' || r == 'R' {
			a++
		}
	}
	return a
}

func parseSeat(s string) (int, int) {
	row := parseAddress(s[:7])
	col := parseAddress(s[7:])
	return row, col
}

func getSeatID(s string) int {
	row, col := parseSeat(s)
	return row*8 + col
}

func runPartOne() error {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		return err
	}
	var max int
	for _, l := range strings.Split(string(content), "\n") {
		if l == "" {
			continue
		}
		sID := getSeatID(l)
		if sID > max {
			max = sID
		}
	}
	fmt.Println(max)
	return nil
}

func runPartTwo() error {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		return err
	}

	lines := strings.Split(string(content), "\n")
	seatIDs := make([]int, 0, len(lines))
	for _, l := range lines {
		if l == "" {
			continue
		}
		sID := getSeatID(l)
		seatIDs = append(seatIDs, sID)
	}

	var minID int
	for _, id := range seatIDs {
		if minID == 0 || id < minID {
			minID = id
		}
	}

	rowCount := (len(content) / 8) + 1
	foundSeats := make([]bool, rowCount*8)
	for _, id := range seatIDs {
		foundSeats[id-minID] = true
	}

	for idx, found := range foundSeats {
		if !found {
			fmt.Println(minID + idx)
			return nil
		}
	}

	return errors.New("missing seat not found")
}

func main() {
	if err := runPartOne(); err != nil {
		log.Fatal(err)
	}
	if err := runPartTwo(); err != nil {
		log.Fatal(err)
	}
}
