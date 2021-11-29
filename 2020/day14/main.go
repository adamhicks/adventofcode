package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
)

type program interface {
	setMask(string)
	setValue(int64, int64)
	sum() int64
}

type intMask struct {
	set, clear int64
}

func (m *intMask) shift(setBit, clearBit int64) {
	m.set <<= 1
	m.clear <<= 1
	m.set |= setBit
	m.clear |= clearBit
}

func (m intMask) apply(i int64) int64 {
	i = i | m.set
	i = i &^ m.clear
	return i
}

type programV1 struct {
	valMask intMask
	mem     map[int64]int64
}

func newProgramV1() *programV1 {
	return &programV1{mem: make(map[int64]int64)}
}

func (p programV1) sum() int64 {
	var s int64
	for _, v := range p.mem {
		s += v
	}
	return s
}

func (p *programV1) setMask(mask string) {
	var m intMask
	for _, c := range mask {
		var set, clear int64
		if c == '1' {
			set = 1
		} else if c == '0' {
			clear = 1
		}
		m.shift(set, clear)
	}
	p.valMask = m
}

func (p *programV1) setValue(addr, val int64) {
	p.mem[addr] = p.valMask.apply(val)
}

func run(p program, ins []instruction) {
	for _, i := range ins {
		switch c := i.(type) {
		case setMask:
			p.setMask(c.mask)
		case setValue:
			p.setValue(c.address, c.value)
		default:
			panic("unknown instruction")
		}
	}
}

type instruction interface{}

type setMask struct {
	mask string
}

type setValue struct {
	address, value int64
}

func parseInput(content string) ([]instruction, error) {
	re, err := regexp.Compile("(\\w+)(\\[(\\d+)\\])? = (.*)")
	if err != nil {
		return nil, err
	}

	matches := re.FindAllStringSubmatch(content, -1)
	ret := make([]instruction, 0, len(matches))

	for _, m := range matches {
		switch m[1] {
		case "mem":
			addr, err := strconv.Atoi(m[3])
			if err != nil {
				return nil, err
			}
			val, err := strconv.Atoi(m[4])
			if err != nil {
				return nil, err
			}
			ret = append(ret, setValue{address: int64(addr), value: int64(val)})
		case "mask":
			ret = append(ret, setMask{mask: m[4]})
		default:
			return nil, errors.New("invalid command")
		}
	}
	return ret, nil
}

func runPartOne() error {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		return err
	}
	ins, err := parseInput(string(content))
	if err != nil {
		return err
	}
	p := newProgramV1()
	run(p, ins)
	fmt.Println(p.sum())
	return nil
}

type programV2 struct {
	addrMasks []intMask
	mem       map[int64]int64
}

func newProgramV2() *programV2 {
	return &programV2{mem: make(map[int64]int64)}
}

func (p *programV2) setMask(mask string) {
	addrMasks := []intMask{{}}
	for _, m := range mask {
		switch m {
		case '0':
			for i := range addrMasks {
				addrMasks[i].shift(0, 0)
			}
		case '1':
			for i := range addrMasks {
				addrMasks[i].shift(1, 0)
			}
		case 'X':
			for i := range addrMasks {
				aCopy := addrMasks[i]
				addrMasks[i].shift(0, 1)
				aCopy.shift(1, 0)
				addrMasks = append(addrMasks, aCopy)
			}
		}
	}
	p.addrMasks = addrMasks
}

func (p *programV2) setValue(addr, val int64) {
	for _, am := range p.addrMasks {
		a := am.apply(addr)
		p.mem[a] = val
	}
}

func (p programV2) sum() int64 {
	var s int64
	for _, v := range p.mem {
		s += v
	}
	return s
}

func runPartTwo() error {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		return err
	}
	ins, err := parseInput(string(content))
	if err != nil {
		return err
	}
	p := newProgramV2()
	run(p, ins)
	fmt.Println(p.sum())
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
