package transform

import (
	"github.com/adm87/flick/pkg/geom"
	"github.com/yohamta/donburi"
)

type BoundsModel struct {
	Rect geom.Rect
}

var BoundsComponent = donburi.NewComponentType[BoundsModel]()

func GetBounds(entry *donburi.Entry) (*BoundsModel, bool) {
	if entry.HasComponent(BoundsComponent) {
		return BoundsComponent.Get(entry), true
	}
	return nil, false
}

func GetWorldBounds(entry *donburi.Entry) (*geom.Rect, bool) {
	tr, ok := GetTransform(entry)
	if !ok {
		return nil, false
	}
	bounds, ok := GetBounds(entry)
	if !ok {
		return nil, false
	}
	worldRect := bounds.Rect.Transform(GetTransformMatrix(entry, tr))
	return &worldRect, true
}

func (b *BoundsModel) GetPosition() (x, y float64) {
	return b.Rect.X, b.Rect.Y
}

func (b *BoundsModel) SetPosition(x, y float64) {
	b.Rect.X = x
	b.Rect.Y = y
}

func (b *BoundsModel) GetSize() (width, height float64) {
	return b.Rect.Width, b.Rect.Height
}

func (b *BoundsModel) SetSize(width, height float64) {
	b.Rect.Width = width
	b.Rect.Height = height
}
