package aoc

type Rect struct {
	From, To Vec2
}

func (r Rect) Overlap(b Rect) Rect {
	from := Vec2{X: max(r.From.X, b.From.X), Y: max(r.From.Y, b.From.Y)}
	to := Vec2{X: min(r.To.X, b.To.X), Y: min(r.To.Y, b.To.Y)}

	if from.X >= to.X || from.Y >= to.Y {
		return Rect{}
	}
	return Rect{From: from, To: to}
}

var zero Rect

func (r Rect) IsZero() bool {
	return r == zero
}
