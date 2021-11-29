package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMove(t *testing.T) {
	var c coord

	c = move(c, dirEast)
	assert.NotEqual(t, coord{}, c)
	c = move(c, dirWest)
	assert.Equal(t, coord{}, c)

	c = move(c, dirNorthEast)
	assert.NotEqual(t, coord{}, c)
	c = move(c, dirSouthWest)
	assert.Equal(t, coord{}, c)

	c = move(c, dirNorthWest)
	assert.NotEqual(t, coord{}, c)
	c = move(c, dirSouthEast)
	assert.Equal(t, coord{}, c)
}

func TestMoveDiagonal(t *testing.T) {
	var a coord
	a = move(a, dirNorthEast)
	a = move(a, dirNorthEast)
	a = move(a, dirSouthEast)
	a = move(a, dirSouthEast)

	var b coord
	b = move(b, dirEast)
	b = move(b, dirEast)

	assert.Equal(t, a, b)
}

const testInput = `sesenwnenenewseeswwswswwnenewsewsw
neeenesenwnwwswnenewnwwsewnenwseswesw
seswneswswsenwwnwse
nwnwneseeswswnenewneswwnewseswneseene
swweswneswnenwsewnwneneseenw
eesenwseswswnenwswnwnwsewwnwsene
sewnenenenesenwsewnenwwwse
wenwwweseeeweswwwnwwe
wsweesenenewnwwnwsenewsenwwsesesenwne
neeswseenwwswnwswswnw
nenwswwsewswnenenewsenwsenwnesesenew
enewnwewneswsewnwswenweswnenwsenwsw
sweneswneswneneenwnewenewwneswswnese
swwesenesewenwneswnwwneseswwne
enesenwswwswneneswsenwnewswseenwsese
wnwnesenesenenwwnenwsewesewsesesew
nenewswnwewswnenesenwnesewesw
eneswnwswnwsenenwnwnwwseeswneewsenese
neswnwewnwnwseenwseesewsenwsweewe
wseweeenwnesenwwwswnew`

func TestExample(t *testing.T) {
	ins := parseInput(testInput)
	s := findBlackTiles(ins)
	assert.Equal(t, 10, len(s))
}

func TestPartTwo(t *testing.T) {
	ins := parseInput(testInput)
	s := findBlackTiles(ins)

	for i := 0; i < 100; i++ {
		s = tick(s)
	}
	assert.Equal(t, 2208, len(s))
}
