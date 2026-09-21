package transform

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
)

type TransformModel struct {
	x, y    float64
	ox, oy  float64
	sx, sy  float64
	rot     float64
	isDirty bool
}

type MatrixModel struct {
	GeoM ebiten.GeoM
}

var (
	TransformComponent = donburi.NewComponentType[TransformModel](DefaultTransform())
	MatrixComponent    = donburi.NewComponentType[MatrixModel]()
)

func DefaultTransform() TransformModel {
	return TransformModel{
		sx:      1,
		sy:      1,
		isDirty: true,
	}
}

func GetTransform(entry *donburi.Entry) (*TransformModel, bool) {
	if entry.HasComponent(TransformComponent) {
		return TransformComponent.Get(entry), true
	}
	return nil, false
}

func GetMatrix(entry *donburi.Entry) ebiten.GeoM {
	t, ok := GetTransform(entry)
	if !ok {
		return ebiten.GeoM{}
	}
	return GetTransformMatrix(entry, t)
}

func GetTransformMatrix(entry *donburi.Entry, t *TransformModel) ebiten.GeoM {
	if !entry.HasComponent(MatrixComponent) {
		var geom ebiten.GeoM
		geom.Translate(-t.ox, -t.oy)
		geom.Scale(t.sx, t.sy)
		geom.Rotate(t.rot)
		geom.Translate(t.x, t.y)
		return geom
	}

	m := MatrixComponent.Get(entry)

	if t.isDirty {
		m.GeoM.Reset()
		m.GeoM.Translate(-t.ox, -t.oy)
		m.GeoM.Scale(t.sx, t.sy)
		m.GeoM.Rotate(t.rot)
		m.GeoM.Translate(t.x, t.y)
		t.isDirty = false
	}

	return m.GeoM
}

func (t *TransformModel) Position() (float64, float64) {
	return t.x, t.y
}

func (t *TransformModel) SetPosition(x, y float64) {
	if t.x != x || t.y != y {
		t.x, t.y = x, y
		t.isDirty = true
	}
}

func (t *TransformModel) Origin() (float64, float64) {
	return t.ox, t.oy
}

func (t *TransformModel) SetOrigin(originX, originY float64) {
	if t.ox != originX || t.oy != originY {
		t.ox, t.oy = originX, originY
		t.isDirty = true
	}
}

func (t *TransformModel) Scale() (float64, float64) {
	return t.sx, t.sy
}

func (t *TransformModel) SetScale(scaleX, scaleY float64) {
	if t.sx != scaleX || t.sy != scaleY {
		t.sx, t.sy = scaleX, scaleY
		t.isDirty = true
	}
}

func (t *TransformModel) Rotation() float64 {
	return t.rot
}

func (t *TransformModel) SetRotation(rotation float64) {
	if t.rot != rotation {
		t.rot = rotation
		t.isDirty = true
	}
}

func (t *TransformModel) IsDirty() bool {
	return t.isDirty
}

func (t *TransformModel) SetDirty(dirty bool) {
	t.isDirty = dirty
}
