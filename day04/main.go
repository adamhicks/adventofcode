package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type document map[string]string

func parseInput(s string) ([]document, error) {
	lineRe, err := regexp.Compile("(\\w+):([\\w\\d#]+)")
	if err != nil {
		return nil, err
	}

	var docs []document
	cur := make(document)
	for _, line := range strings.Split(s, "\n") {
		if line == "" {
			docs = append(docs, cur)
			cur = make(document)
		}
		matches := lineRe.FindAllSubmatch([]byte(line), -1)
		for _, match := range matches {
			key, val := string(match[1]), string(match[2])
			cur[key] = val
		}
	}
	return docs, nil
}

const (
	identBirthYear      = "byr"
	identIssueYear      = "iyr"
	identExpirationYear = "eyr"
	identHeight         = "hgt"
	identHairColour     = "hcl"
	identEyeColor       = "ecl"
	identPassportID     = "pid"
	identCountryID      = "cid"
)

func isDocValid(d document) bool {
	if len(d) == 8 {
		return true
	}
	if len(d) < 7 {
		return false
	}
	_, ok := d[identCountryID]
	return !ok
}

func countValid(docs []document) int {
	var sum int
	for _, d := range docs {
		if isDocValid(d) {
			sum++
		}
	}
	return sum
}

func runPartOne() error {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		return err
	}
	docs, err := parseInput(string(content))
	if err != nil {
		return err
	}
	fmt.Println(countValid(docs))
	return nil
}

func numberInRange(min, max int) func(string) bool {
	return func(s string) bool {
		i, err := strconv.Atoi(s)
		if err != nil {
			return false
		}
		return i >= min && i <= max
	}
}

func regexMatch(reStr string) func(string) bool {
	re, err := regexp.Compile(reStr)
	if err != nil {
		panic(err)
	}
	return re.MatchString
}

func oneOf(check []string) func(string) bool {
	return func(s string) bool {
		for _, t := range check {
			if s == t {
				return true
			}
		}
		return false
	}
}

func checkHeight(s string) bool {
	if strings.HasSuffix(s, "cm") {
		i, err := strconv.Atoi(strings.TrimSuffix(s, "cm"))
		if err != nil {
			panic(err)
		}
		return i >= 150 && i <= 193
	}
	if strings.HasSuffix(s, "in") {
		i, err := strconv.Atoi(strings.TrimSuffix(s, "in"))
		if err != nil {
			panic(err)
		}
		return i >= 59 && i <= 76
	}
	return false
}

var fieldValidation = map[string]func(string) bool{
	identBirthYear:      numberInRange(1920, 2002),
	identIssueYear:      numberInRange(2010, 2020),
	identExpirationYear: numberInRange(2020, 2030),
	identHeight:         checkHeight,
	identHairColour:     regexMatch("^#[0-9a-f]{6}$"),
	identEyeColor:       oneOf([]string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}),
	identPassportID:     regexMatch("^\\d{9}$"),
}

func isDocValid2(d document) (bool, error) {
	for f, check := range fieldValidation {
		val, ok := d[f]
		if !ok || !check(val) {
			return false, fmt.Errorf("invalid %s = %s", f, val)
		}
	}
	return true, nil
}

func countValid2(docs []document) int {
	var sum int
	for _, d := range docs {
		ok, _ := isDocValid2(d)
		if ok {
			sum++
		}
	}
	return sum
}

func runPartTwo() error {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		return err
	}
	docs, err := parseInput(string(content))
	if err != nil {
		return err
	}
	fmt.Println(countValid2(docs))
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
