package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type interval struct {
	low, high int
}

func (i interval) Intersects(o interval) bool {
	return o.low <= i.high && o.high >= i.low
}

func (i interval) Union(o interval) interval {
	min, max := i.low, i.high
	if o.low < min {
		min = o.low
	}
	if o.high > max {
		max = o.high
	}
	return interval{low: min, high: max}
}

func (i interval) HasValue(v int) bool {
	return v >= i.low && v < i.high
}

func simplifyIntervals(ivs []interval) []interval {
	var ret []interval
	for _, i := range ivs {
		var merged bool
		for idx := range ret {
			if i.Intersects(ret[idx]) {
				ret[idx] = ret[idx].Union(i)
				merged = true
				break
			}
		}
		if !merged {
			ret = append(ret, i)
		}
	}
	return ret
}

type field struct {
	name      string
	intervals []interval
}

type ticketInfo struct {
	fields        []field
	yourTicket    []int
	nearbyTickets [][]int
}

func parseTicket(s string) ([]int, error) {
	var ret []int
	for _, val := range strings.Split(s, ",") {
		i, err := strconv.Atoi(val)
		if err != nil {
			return nil, err
		}
		ret = append(ret, i)
	}
	return ret, nil
}

func parseInput(content string) (ticketInfo, error) {
	parts := strings.Split(content, "\n\n")
	var ti ticketInfo

	fieldRe, err := regexp.Compile("(\\d+)-(\\d+)")
	if err != nil {
		return ti, err
	}
	for _, l := range strings.Split(parts[0], "\n") {
		if l == "" {
			continue
		}
		parts := strings.Split(l, ":")
		field := field{name: parts[0]}
		for _, match := range fieldRe.FindAllStringSubmatch(parts[1], -1) {
			low, err := strconv.Atoi(match[1])
			if err != nil {
				return ti, err
			}
			high, err := strconv.Atoi(match[2])
			if err != nil {
				return ti, err
			}
			field.intervals = append(field.intervals, interval{low: low, high: high + 1})
		}
		ti.fields = append(ti.fields, field)
	}
	ti.yourTicket, err = parseTicket(strings.Split(parts[1], "\n")[1])
	if err != nil {
		return ti, err
	}
	for _, l := range strings.Split(parts[2], "\n")[1:] {
		if l == "" {
			continue
		}
		tick, err := parseTicket(l)
		if err != nil {
			return ti, err
		}
		ti.nearbyTickets = append(ti.nearbyTickets, tick)
	}

	return ti, nil
}

func simplifyFields(fields []field) []interval {
	var flatten []interval
	for _, f := range fields {
		flatten = append(flatten, f.intervals...)
	}
	return simplifyIntervals(flatten)
}

func checkTicketSimple(ticket []int, ints []interval) []int {
	var ret []int
	for _, i := range ticket {
		var isValid bool
		for _, iv := range ints {
			if iv.HasValue(i) {
				isValid = true
				break
			}
		}
		if !isValid {
			ret = append(ret, i)
		}
	}
	return ret
}

func runPartOne() error {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		return err
	}
	ti, err := parseInput(string(content))
	if err != nil {
		return err
	}

	ints := simplifyFields(ti.fields)
	var s int
	for _, tick := range ti.nearbyTickets {
		for _, i := range checkTicketSimple(tick, ints) {
			s += i
		}
	}
	fmt.Println(s)
	return nil
}

func pivot(in [][]int) [][]int {
	ret := make([][]int, len(in[0]))
	for i := range ret {
		for _, j := range in {
			ret[i] = append(ret[i], j[i])
		}
	}
	return ret
}

func findValidFields(vals []int, fields []field) []int {
	var ret []int
	for i, f := range fields {
		fieldValid := true

		for _, v := range vals {
			var valid bool
			for _, iv := range f.intervals {
				if iv.HasValue(v) {
					valid = true
					break
				}
			}
			if !valid {
				fieldValid = false
				break
			}
		}
		if fieldValid {
			ret = append(ret, i)
		}
	}
	return ret
}

type columnOptions struct {
	index        int
	fieldOptions []int
}

func runPartTwo() error {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		return err
	}
	ti, err := parseInput(string(content))
	if err != nil {
		return err
	}

	ints := simplifyFields(ti.fields)
	var tickets [][]int
	for _, tick := range ti.nearbyTickets {
		if len(checkTicketSimple(tick, ints)) == 0 {
			tickets = append(tickets, tick)
		}
	}

	cols := pivot(tickets)
	var colOpts []columnOptions
	for i := range cols {
		colOpts = append(colOpts, columnOptions{
			index:        i,
			fieldOptions: findValidFields(cols[i], ti.fields),
		})
	}
	sort.Slice(colOpts, func(i, j int) bool {
		return len(colOpts[i].fieldOptions) < len(colOpts[j].fieldOptions)
	})

	fieldMap := make(map[int]int)
	for _, p := range colOpts {
		for _, o := range p.fieldOptions {
			if _, ok := fieldMap[o]; !ok {
				fieldMap[o] = p.index
				break
			}
		}
	}

	mult := 1
	for fIdx, colIdx := range fieldMap {
		field := ti.fields[fIdx]
		if !strings.HasPrefix(field.name, "departure") {
			continue
		}
		mult *= ti.yourTicket[colIdx]
	}

	fmt.Println(mult)
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
