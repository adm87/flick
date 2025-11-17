package models

import (
	"github.com/hajimehoshi/ebiten/v2"
)

var DefaultTransform = Transform{
	xy:      [2]float32{0, 0},
	origin:  [2]float32{0, 0},
	scale:   [2]float32{1, 1},
	isDirty: true,
}

// Transform represents position, scale, and rotation of an entity.
type Transform struct {
	xy       [2]float32
	origin   [2]float32
	scale    [2]float32
	rotation float32

	matrix    ebiten.GeoM
	invMatrix ebiten.GeoM

	isDirty bool
}

func (t *Transform) Position() (float32, float32) {
	return t.xy[0], t.xy[1]
}

func (t *Transform) Origin() (float32, float32) {
	return t.origin[0], t.origin[1]
}

func (t *Transform) Scale() (float32, float32) {
	return t.scale[0], t.scale[1]
}

func (t *Transform) Rotation() float32 {
	return t.rotation
}

func (t *Transform) SetPosition(x, y float32) *Transform {
	t.xy[0] = x
	t.xy[1] = y
	t.isDirty = true
	return t
}

func (t *Transform) SetOrigin(ox, oy float32) *Transform {
	t.origin[0] = ox
	t.origin[1] = oy
	t.isDirty = true
	return t
}

func (t *Transform) SetScale(sx, sy float32) *Transform {
	t.scale[0] = sx
	t.scale[1] = sy
	t.isDirty = true
	return t
}

func (t *Transform) SetRotation(r float32) *Transform {
	t.rotation = r
	t.isDirty = true
	return t
}

func (t *Transform) Matrix() ebiten.GeoM {
	if t.isDirty {
		t.matrix.Reset()
		t.matrix.Translate(-float64(t.origin[0]), -float64(t.origin[1]))
		t.matrix.Scale(float64(t.scale[0]), float64(t.scale[1]))
		t.matrix.Rotate(float64(t.rotation))
		t.matrix.Translate(float64(t.xy[0]), float64(t.xy[1]))
		t.isDirty = false
	}
	return t.matrix
}

func (t *Transform) InvMatrix() ebiten.GeoM {
	if t.isDirty {
		t.invMatrix = t.Matrix()
		t.invMatrix.Invert()
	}
	return t.invMatrix
}
