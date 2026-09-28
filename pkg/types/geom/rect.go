package geom

import "github.com/hajimehoshi/ebiten/v2"

type Rect struct {
	X, Y          float64
	Width, Height float64
}

func NewRect(x, y, width, height float64) Rect {
	return Rect{X: x, Y: y, Width: width, Height: height}
}

func (r Rect) Min() (float64, float64) {
	return r.X, r.Y
}

func (r Rect) Max() (float64, float64) {
	return r.X + r.Width, r.Y + r.Height
}

func (r Rect) Center() (float64, float64) {
	return r.X + r.Width*0.5, r.Y + r.Height*0.5
}

func (r Rect) Area() float64 {
	return r.Width * r.Height
}

func (r Rect) Contains(x, y float64) bool {
	return x >= r.X && x <= r.X+r.Width && y >= r.Y && y <= r.Y+r.Height
}

func (r Rect) Intersects(other Rect) bool {
	return r.X < other.X+other.Width && r.X+r.Width > other.X && r.Y < other.Y+other.Height && r.Y+r.Height > other.Y
}

func (r Rect) Union(other Rect) Rect {
	x1 := min(r.X, other.X)
	y1 := min(r.Y, other.Y)
	x2 := max(r.X+r.Width, other.X+other.Width)
	y2 := max(r.Y+r.Height, other.Y+other.Height)
	return Rect{X: x1, Y: y1, Width: x2 - x1, Height: y2 - y1}
}

func (r Rect) Transform(m ebiten.GeoM) Rect {
	x1, y1 := m.Apply(r.X, r.Y)
	x2, y2 := m.Apply(r.X+r.Width, r.Y+r.Height)
	return Rect{
		X:      x1,
		Y:      y1,
		Width:  x2 - x1,
		Height: y2 - y1,
	}
}
