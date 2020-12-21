package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"
)

const cellWidth = 10

type edgeID int

func (i *edgeID) Push(pixel pixelContent) {
	(*i) <<= 1
	if pixel == pixelBlack {
		(*i)++
	}
}

func (i edgeID) Reverse() edgeID {
	var rev edgeID
	for j := 0; j < cellWidth; j++ {
		if i%2 == 1 {
			rev.Push(pixelBlack)
		} else {
			rev.Push(pixelWhite)
		}
		i >>= 1
	}
	return rev
}

func (i edgeID) Symmetric() int {
	return int(i) + int(i.Reverse())
}

func (i edgeID) String() string {
	b := make([]byte, cellWidth)
	for j := 0; j < cellWidth; j++ {
		if i%2 == 0 {
			b[cellWidth-1-j] = '.'
		} else {
			b[cellWidth-1-j] = '#'
		}
		i >>= 1
	}
	return string(b)
}

const (
	edgeTop    = 0
	edgeRight  = 1
	edgeBottom = 2
	edgeLeft   = 3
)

type image struct {
	ID       int
	Data     [][]pixelContent
	EdgeIDs  []edgeID
	Rotation int
	Flipped  bool
}

func newImage(id int, data [][]pixelContent) *image {
	var top edgeID
	for _, p := range data[0] {
		top.Push(p)
	}
	var bottom edgeID
	for i := range data[cellWidth-1] {
		bottom.Push(data[cellWidth-1][cellWidth-1-i])
	}
	var left, right edgeID
	for i := range data {
		left.Push(data[cellWidth-1-i][0])
		right.Push(data[i][cellWidth-1])
	}
	return &image{
		ID:      id,
		Data:    data,
		EdgeIDs: []edgeID{top, right, bottom, left},
	}
}

func (i *image) UniqueEdges(eCount map[edgeID]int) int {
	var s int
	for _, ie := range i.EdgeIDs {
		if eCount[ie] == 1 {
			s++
		}
	}
	return s
}

func (i *image) Rotate() {
	i.Rotation++
	last := i.EdgeIDs[3]
	i.EdgeIDs = append([]edgeID{last}, i.EdgeIDs[0:3]...)
}

func (i *image) FlipVertically() {
	i.Flipped = !i.Flipped
	i.EdgeIDs[edgeTop], i.EdgeIDs[edgeBottom] = i.EdgeIDs[edgeBottom], i.EdgeIDs[edgeTop]
	for dir := range i.EdgeIDs {
		i.EdgeIDs[dir] = i.EdgeIDs[dir].Reverse()
	}
}

func (i *image) RenderInner(buffer [][]pixelContent, xOff, yOff int) {
	var t []coordTransform
	if i.Flipped {
		t = append(t, flip)
	}
	for r := 0; r < i.Rotation; r++ {
		t = append(t, rotate)
	}

	for y := 1; y < cellWidth-1; y++ {
		for x := 1; x < cellWidth-1; x++ {
			newX, newY := transformCoord(x-1, y-1, cellWidth-2, t...)
			buffer[yOff+newY][xOff+newX] = i.Data[y][x]
		}
	}
}

type pixelContent int

const (
	pixelWhite pixelContent = 0
	pixelBlack pixelContent = 1
)

type coordTransform func(x, y int, n int) (int, int)

func rotate(x, y int, n int) (int, int) {
	return n - 1 - y, x
}

func flip(x, y int, n int) (int, int) {
	return x, n - 1 - y
}

func transformCoord(x, y int, n int, trans ...coordTransform) (int, int) {
	for _, t := range trans {
		x, y = t(x, y, n)
	}
	return x, y
}

func transformImage(data [][]pixelContent, trans ...coordTransform) [][]pixelContent {
	n := len(data)
	newData := make([][]pixelContent, n)
	for i := range newData {
		newData[i] = make([]pixelContent, n)
	}

	for y := 0; y < n; y++ {
		for x := 0; x < n; x++ {
			newX, newY := transformCoord(x, y, n, trans...)
			newData[newY][newX] = data[y][x]
		}
	}
	return newData
}

func getImage(data [][]pixelContent) string {
	var b strings.Builder
	for _, row := range data {
		for _, p := range row {
			if p == pixelWhite {
				fmt.Fprint(&b, ".")
			} else {
				fmt.Fprint(&b, "#")
			}
		}
		fmt.Fprint(&b, "\n")
	}
	return b.String()
}

func parseInput(content string) ([]*image, error) {
	idRe, err := regexp.Compile("Tile (\\d+):")
	if err != nil {
		return nil, err
	}

	var id int
	var data [][]pixelContent
	var ret []*image

	for _, l := range strings.Split(content, "\n") {
		if l == "" {
			continue
		}
		if id == 0 {
			m := idRe.FindAllStringSubmatch(l, -1)
			if len(m) != 1 {
				return nil, errors.New("missing image ID")
			}
			i, err := strconv.Atoi(m[0][1])
			if err != nil {
				return nil, err
			}
			id = i
			continue
		}
		line := make([]pixelContent, len(l))
		for i, c := range l {
			if c == '#' {
				line[i] = pixelBlack
			}
		}
		data = append(data, line)
		if len(data) == 10 {
			ret = append(ret, newImage(id, data))
			id = 0
			data = nil
		}
	}
	return ret, nil
}

func countEdgeIDs(imgs []*image) map[edgeID]int {
	count := make(map[edgeID]int)
	for _, i := range imgs {
		for _, id := range i.EdgeIDs {
			count[id]++
			count[id.Reverse()]++
		}
	}
	return count
}

func findCorners(imgs []*image) []*image {
	counts := countEdgeIDs(imgs)
	var corners []*image

	for _, i := range imgs {
		if i.UniqueEdges(counts) == 2 {
			corners = append(corners, i)
		}
	}
	return corners
}

func runPartOne() error {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		return err
	}
	imgs, err := parseInput(string(content))
	if err != nil {
		return err
	}
	i := 1
	for _, corn := range findCorners(imgs) {
		i *= corn.ID
	}
	fmt.Println(i)
	return nil
}

type imageConstructor struct {
	d               int
	remainingImages []*image
	edgeCount       map[edgeID]int
}

func (c *imageConstructor) isEdgeUnique(e edgeID) bool {
	return c.edgeCount[e] == 1
}

func (c *imageConstructor) popTopLeft() *image {
	for idx, i := range c.remainingImages {
		if i.UniqueEdges(c.edgeCount) == 2 {
			c.remainingImages = append(c.remainingImages[:idx], c.remainingImages[idx+1:]...)
			for j := 0; j < 4; j++ {
				if c.edgeCount[i.EdgeIDs[edgeTop]] == 1 && c.edgeCount[i.EdgeIDs[edgeLeft]] == 1 {
					return i
				}
				i.Rotate()
			}
			panic("rotated 4 times and still not top left")
		}
	}
	panic("no top left")
}

func (c *imageConstructor) popImageMatchingEdge(e edgeID, dir int) *image {
	er := e.Reverse()
	for idx, i := range c.remainingImages {
		for _, ie := range i.EdgeIDs {
			if ie != e && ie != er {
				continue
			}
			c.remainingImages = append(c.remainingImages[:idx], c.remainingImages[idx+1:]...)

			if ie == er {
				i.FlipVertically()
			}
			for j := 0; j < 4; j++ {
				if i.EdgeIDs[dir] == e {
					return i
				}
				i.Rotate()
			}
			panic("rotated 4 times but didn't match")
		}
	}
	return nil
}

func solveImage(imgs []*image) [][]*image {
	con := imageConstructor{
		d:               int(math.Sqrt(float64(len(imgs)))),
		remainingImages: imgs,
		edgeCount:       countEdgeIDs(imgs),
	}
	topLeft := con.popTopLeft()
	rightEdge := topLeft.EdgeIDs[edgeRight]

	row := []*image{topLeft}
	for i := 1; i < con.d; i++ {
		nextImg := con.popImageMatchingEdge(rightEdge.Reverse(), edgeLeft)
		rightEdge = nextImg.EdgeIDs[edgeRight]
		row = append(row, nextImg)
	}

	fullImage := [][]*image{row}
	prevRow := row

	for i := 1; i < con.d; i++ {
		var thisRow []*image
		for j := 0; j < con.d; j++ {
			above := prevRow[j].EdgeIDs[edgeBottom]
			img := con.popImageMatchingEdge(above.Reverse(), edgeTop)
			thisRow = append(thisRow, img)
		}
		fullImage = append(fullImage, thisRow)
		prevRow = thisRow
	}
	return fullImage
}

func renderFullImage(full [][]*image) [][]pixelContent {
	d := len(full)
	width := cellWidth - 2

	img := make([][]pixelContent, width*d)
	for i := range img {
		img[i] = make([]pixelContent, width*d)
	}

	for gY := 0; gY < d; gY++ {
		for gX := 0; gX < d; gX++ {
			i := full[gY][gX]
			i.RenderInner(img, gX*width, gY*width)
		}
	}
	return img
}

const nessyStr = `                  #
#    ##    ##    ###
 #  #  #  #  #  #   `

type coord struct {
	x, y int
}

func compileNessy() map[coord]bool {
	nessy := make(map[coord]bool)
	for y, l := range strings.Split(nessyStr, "\n") {
		for x, c := range l {
			if c == '#' {
				nessy[coord{x: x, y: y}] = true
			}
		}
	}
	return nessy
}

func searchForPattern(p map[coord]bool, img [][]pixelContent) int {
	n := len(img)
	allBlacks := make([][]pixelContent, n)
	for i := range allBlacks {
		allBlacks[i] = make([]pixelContent, n)
		copy(allBlacks[i], img[i])
	}

	var maxPx, maxPy int
	for c := range p {
		if c.x > maxPx {
			maxPx = c.x
		}
		if c.y > maxPy {
			maxPy = c.y
		}
	}

	for y := 0; y < n-maxPy; y++ {
		for x := 0; x < n-maxPx; x++ {
			var failMatch bool
			for c := range p {
				if img[y+c.y][x+c.x] != pixelBlack {
					failMatch = true
					break
				}
			}
			if !failMatch {
				for c := range p {
					allBlacks[y+c.y][x+c.x] = pixelWhite
				}
			}
		}
	}

	var remBlacks int
	for _, row := range allBlacks {
		for _, p := range row {
			if p == pixelBlack {
				remBlacks++
			}
		}
	}
	return remBlacks
}

func searchAllOrientations(p map[coord]bool, orig [][]pixelContent) []int {
	var res []int
	var t []coordTransform
	for i := 0; i < 4; i++ {
		img := transformImage(orig, t...)
		rb := searchForPattern(p, img)
		res = append(res, rb)
		t = append(t, rotate)
	}
	t = []coordTransform{flip}
	for i := 0; i < 4; i++ {
		img := transformImage(orig, t...)
		rb := searchForPattern(p, img)
		res = append(res, rb)
		t = append(t, rotate)
	}
	return res
}

func min(v []int) int {
	if len(v) == 0 {
		return 0
	}
	min := v[0]
	for _, i := range v {
		if i < min {
			min = i
		}
	}
	return min
}

func runPartTwo() error {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		return err
	}
	imgs, err := parseInput(string(content))
	if err != nil {
		return err
	}
	full := solveImage(imgs)
	img := renderFullImage(full)
	nessy := compileNessy()
	res := searchAllOrientations(nessy, img)
	fmt.Println(min(res))
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
