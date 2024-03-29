package aoc

var (
	North = Vec2{Y: -1}
	East  = Vec2{X: 1}
	South = Vec2{Y: 1}
	West  = Vec2{X: -1}
)

var Cardinals = []Vec2{North, East, South, West}

type Vec2 struct {
	X, Y int
}

func (v Vec2) Sub(b Vec2) Vec2 {
	return Vec2{v.X - b.X, v.Y - b.Y}
}

func (v Vec2) Add(b Vec2) Vec2 {
	return Vec2{v.X + b.X, v.Y + b.Y}
}

func (v Vec2) Distance(b Vec2) int {
	d := b.Sub(v)
	return Abs(d.X) + Abs(d.Y)
}

func (v Vec2) Neighbours() []Vec2 {
	return []Vec2{
		{X: v.X - 1, Y: v.Y - 1},
		{X: v.X, Y: v.Y - 1},
		{X: v.X + 1, Y: v.Y - 1},
		{X: v.X - 1, Y: v.Y},
		{X: v.X + 1, Y: v.Y},
		{X: v.X - 1, Y: v.Y + 1},
		{X: v.X, Y: v.Y + 1},
		{X: v.X + 1, Y: v.Y + 1},
	}
}

func (v Vec2) Orthogonal() []Vec2 {
	return []Vec2{
		{X: v.X, Y: v.Y - 1},
		{X: v.X - 1, Y: v.Y},
		{X: v.X + 1, Y: v.Y},
		{X: v.X, Y: v.Y + 1},
	}
}

func (v Vec2) InRange(maxim Vec2) bool {
	return v.X >= 0 && v.X < maxim.X && v.Y >= 0 && v.Y < maxim.Y
}

func (v Vec2) TurnLeft() Vec2 {
	return Vec2{X: v.Y, Y: -v.X}
}

func (v Vec2) TurnRight() Vec2 {
	return Vec2{X: -v.Y, Y: v.X}
}

func (v Vec2) Mul(i int) Vec2 {
	return Vec2{X: v.X * i, Y: v.Y * i}
}
