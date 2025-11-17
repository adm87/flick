package shapes

import "errors"

type ShapeType uint8

const (
	ShapeTypeRectangle ShapeType = iota
	ShapeTypePolygon
)

type Shape interface {
	Position() (float32, float32)
	SetPosition(x, y float32) Shape

	// Bounds returns the minimum and maximum bounding box of the shape.
	// The returned array is in the format [minX, minY, maxX, maxY].
	// The x and y parameters specify the position to calculate the bounds at.
	Bounds(x, y float32) [4]float32

	// Type returns the type of the shape.
	Type() ShapeType
}

func Get[T Shape](s Shape) (T, error) {
	casted, ok := s.(T)
	if !ok {
		var zero T
		return zero, errors.New("failed to cast shape to target type")
	}
	return casted, nil
}
