package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testInput = `Tile 2311:
..##.#..#.
##..#.....
#...##..#.
####.#...#
##.##.###.
##...#.###
.#.#.#..##
..#....#..
###...#.#.
..###..###

Tile 1951:
#.##...##.
#.####...#
.....#..##
#...######
.##.#....#
.###.#####
###.##.##.
.###....#.
..#.#..#.#
#...##.#..

Tile 1171:
####...##.
#..##.#..#
##.#..#.#.
.###.####.
..###.####
.##....##.
.#...####.
#.##.####.
####..#...
.....##...

Tile 1427:
###.##.#..
.#..#.##..
.#.##.#..#
#.#.#.##.#
....#...##
...##..##.
...#.#####
.#.####.#.
..#..###.#
..##.#..#.

Tile 1489:
##.#.#....
..##...#..
.##..##...
..#...#...
#####...#.
#..#.#.#.#
...#.#.#..
##.#...##.
..##.##.##
###.##.#..

Tile 2473:
#....####.
#..#.##...
#.##..#...
######.#.#
.#...#.#.#
.#########
.###.#..#.
########.#
##...##.#.
..###.#.#.

Tile 2971:
..#.#....#
#...###...
#.#.###...
##.##..#..
.#####..##
.#..####.#
#..#.#..#.
..####.###
..#.#.###.
...#.#.#.#

Tile 2729:
...#.#.#.#
####.#....
..#.#.....
....#..#.#
.##..##.#.
.#.####...
####.#.#..
##.####...
##..#.##..
#.##...##.

Tile 3079:
#.#.#####.
.#..######
..#.......
######....
####.#..#.
.#...#.##.
#.#####.##
..#.###...
..#.......
..#.###...
`

func TestReverse(t *testing.T) {
	var e1, e2 edgeID
	s := "........#."
	for _, c := range s {
		if c == '#' {
			e1.Push(pixelBlack)
		} else {
			e1.Push(pixelWhite)
		}
	}
	assert.Equal(t, s, e1.String())
	s2 := ".#........"
	for _, c := range s2 {
		if c == '#' {
			e2.Push(pixelBlack)
		} else {
			e2.Push(pixelWhite)
		}
	}

	e1r := e1.Reverse()
	assert.Equal(t, 2, int(e1))
	assert.Equal(t, e2, e1r)
}

func TestRender(t *testing.T) {
	imgs, err := parseInput(testInput)
	require.NoError(t, err)

	full := solveImage(imgs)

	img := renderFullImage(full)
	img = transformImage(img, flip, rotate)

	nessy := compileNessy()

	res := searchAllOrientations(nessy, img)

	assert.Equal(t, 273, min(res))
}

func TestNessy(t *testing.T) {
	nessy := compileNessy()
	assert.Len(t, nessy, 15)
}

func TestFlip(t *testing.T) {
	imgs, err := parseInput(testInput)
	require.NoError(t, err)

	i := imgs[0]

	expIDs := make([]edgeID, 4)
	copy(expIDs, i.EdgeIDs)

	i.FlipVertically()
	i.FlipVertically()

	assert.Equal(t, expIDs, i.EdgeIDs)
}

func TestParseInput(t *testing.T) {
	imgs, err := parseInput(testInput)
	require.NoError(t, err)

	assert.Len(t, imgs, 9)
	expIDs := []int{2311, 1951, 1171, 1427, 1489, 2473, 2971, 2729, 3079}
	for i, id := range expIDs {
		assert.Equal(t, id, imgs[i].ID)
	}

	corn := findCorners(imgs)
	ids := make(map[int]bool)
	for _, c := range corn {
		ids[c.ID] = true
	}

	exp := map[int]bool{1951: true, 1171: true, 2971: true, 3079: true}
	assert.Equal(t, exp, ids)
}

func TestSolveImage(t *testing.T) {
	imgs, err := parseInput(testInput)
	require.NoError(t, err)

	full := solveImage(imgs)

	assert.Len(t, full, 3)
}
