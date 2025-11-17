package shapes

type Rectangle struct {
	xy   [2]float32
	size [2]float32
}

func NewRectangle() *Rectangle {
	return &Rectangle{
		size: [2]float32{0, 0},
	}
}

func (r *Rectangle) Size() (float32, float32) {
	return r.size[0], r.size[1]
}

func (r *Rectangle) SetSize(width, height float32) *Rectangle {
	r.size[0] = width
	r.size[1] = height
	return r
}

func (r *Rectangle) Position() (float32, float32) {
	return r.xy[0], r.xy[1]
}

func (r *Rectangle) SetPosition(x, y float32) Shape {
	r.xy[0] = x
	r.xy[1] = y
	return r
}

func (r *Rectangle) Bounds(x, y float32) [4]float32 {
	return [4]float32{
		x + r.xy[0],
		y + r.xy[1],
		x + r.xy[0] + r.size[0],
		y + r.xy[1] + r.size[1],
	}
}

func (r *Rectangle) Type() ShapeType {
	return ShapeTypeRectangle
}
