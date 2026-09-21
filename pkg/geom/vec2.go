package geom

type Vec2 struct {
	X, Y float64
}

func NewVec2(x, y float64) Vec2 {
	return Vec2{X: x, Y: y}
}

func (v Vec2) Length() float64 {
	return v.X*v.X + v.Y*v.Y
}

func (v Vec2) Normalize() Vec2 {
	length := v.Length()
	if length == 0 {
		return v
	}
	return Vec2{X: v.X / length, Y: v.Y / length}
}

func (v Vec2) Dot(other Vec2) float64 {
	return v.X*other.X + v.Y*other.Y
}

func (v Vec2) Add(other Vec2) Vec2 {
	return Vec2{X: v.X + other.X, Y: v.Y + other.Y}
}

func (v Vec2) Sub(other Vec2) Vec2 {
	return Vec2{X: v.X - other.X, Y: v.Y - other.Y}
}

func (v Vec2) Mul(scalar float64) Vec2 {
	return Vec2{X: v.X * scalar, Y: v.Y * scalar}
}

func (v Vec2) Div(scalar float64) Vec2 {
	return Vec2{X: v.X / scalar, Y: v.Y / scalar}
}
