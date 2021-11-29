package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
)

type InputLine struct {
	Min, Max int
	Char     rune
	Text     string
}

func ReadAll(s []byte) ([]InputLine, error) {
	re, err := regexp.Compile("(\\d+)-(\\d+) (\\w): (\\w+)")
	if err != nil {
		return nil, err
	}
	matches := re.FindAllSubmatch(s, -1)
	var ret []InputLine
	for _, m := range matches {
		min, err := strconv.Atoi(string(m[1]))
		if err != nil {
			return nil, err
		}
		max, err := strconv.Atoi(string(m[2]))
		if err != nil {
			return nil, err
		}
		ret = append(ret, InputLine{
			Min:  min,
			Max:  max,
			Char: rune(m[3][0]),
			Text: string(m[4]),
		})
	}
	return ret, nil
}

func PartOneValid(in InputLine) bool {
	var count int
	for _, c := range in.Text {
		if c == in.Char {
			count++
			if count > in.Max {
				return false
			}
		}
	}
	return count >= in.Min
}

func runPartOne() error {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		return err
	}
	lines, err := ReadAll(content)
	if err != nil {
		return err
	}
	var sum int
	for _, line := range lines {
		if PartOneValid(line) {
			sum++
		}
	}
	fmt.Println(sum)
	return nil
}

func PartTwoValid(in InputLine) bool {
	var c int
	if in.Text[in.Min-1] == byte(in.Char) {
		c++
	}
	if in.Text[in.Max-1] == byte(in.Char) {
		c++
	}
	return c == 1
}

func runPartTwo() error {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		return err
	}
	lines, err := ReadAll(content)
	if err != nil {
		return err
	}
	var sum int
	for _, line := range lines {
		if PartTwoValid(line) {
			sum++
		}
	}
	fmt.Println(sum)
	return nil
}

func main() {
	err := runPartTwo()
	if err != nil {
		log.Fatal(err)
	}
}
