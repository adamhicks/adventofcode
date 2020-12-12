package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

type instruction struct {
	Command   string
	Magnitude int
}

func parseInput(content string) ([]instruction, error) {
	var ret []instruction
	for _, l := range strings.Split(content, "\n") {
		if l == "" {
			continue
		}
		i, err := strconv.Atoi(l[1:])
		if err != nil {
			return nil, err
		}
		ret = append(ret, instruction{
			Command:   l[0:1],
			Magnitude: i,
		})
	}
	return ret, nil
}

var (
	dirNorth = position{dNorth: 1}
	dirSouth = position{dNorth: -1}
	dirEast  = position{dEast: 1}
	dirWest  = position{dEast: -1}
)

type position struct {
	dNorth, dEast int
}

func (p position) distance() int {
	return abs(p.dNorth) + abs(p.dEast)
}

func abs(i int) int {
	if i > 0 {
		return i
	}
	return -i
}

func rotate(vec position, deg int) position {
	r := float64(deg) * math.Pi / 180

	iSin := int(math.Sin(r))
	iCos := int(math.Cos(r))

	return position{
		dNorth: iCos*vec.dNorth - iSin*vec.dEast,
		dEast:  iCos*vec.dEast + iSin*vec.dNorth,
	}
}

func move(pos, vec position, mag int) position {
	pos.dEast += vec.dEast * mag
	pos.dNorth += vec.dNorth * mag
	return pos
}

func followInstructions(ins []instruction, pos, vec position) position {
	for _, i := range ins {
		switch i.Command {
		case "N":
			pos = move(pos, dirNorth, i.Magnitude)
		case "S":
			pos = move(pos, dirSouth, i.Magnitude)
		case "E":
			pos = move(pos, dirEast, i.Magnitude)
		case "W":
			pos = move(pos, dirWest, i.Magnitude)
		case "L":
			vec = rotate(vec, -i.Magnitude)
		case "R":
			vec = rotate(vec, i.Magnitude)
		case "F":
			pos = move(pos, vec, i.Magnitude)
		default:
			panic(i)
		}
	}
	return pos
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
	pos := followInstructions(ins, position{}, dirEast)
	fmt.Println(pos.distance())
	return nil
}

func followInstructions2(ins []instruction, pos, vec position) position {
	for _, i := range ins {
		switch i.Command {
		case "N":
			vec = move(vec, dirNorth, i.Magnitude)
		case "S":
			vec = move(vec, dirSouth, i.Magnitude)
		case "E":
			vec = move(vec, dirEast, i.Magnitude)
		case "W":
			vec = move(vec, dirWest, i.Magnitude)
		case "L":
			vec = rotate(vec, -i.Magnitude)
		case "R":
			vec = rotate(vec, i.Magnitude)
		case "F":
			pos = move(pos, vec, i.Magnitude)
		default:
			panic(i)
		}
	}
	return pos
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
	pos := followInstructions2(ins, position{}, position{dNorth: 1, dEast: 10})
	fmt.Println(pos.distance())
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
