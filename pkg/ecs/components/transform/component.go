package transform

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
)

type TransformData struct {
	X, Y             float64
	Rotation         float64
	ScaleX, ScaleY   float64
	OriginX, OriginY float64
	IsDirty          bool
}

type MatrixData struct {
	m ebiten.GeoM
}

// TransformComponent represents the position, rotation, scale, and origin of an entity.
var TransformComponent = donburi.NewComponentType[TransformData](TransformData{
	ScaleX: 1,
	ScaleY: 1,
})

// MatrixComponent represents the transformation matrix of an entity.
var MatrixComponent = donburi.NewComponentType[MatrixData](MatrixData{})

func GetComponent(entry *donburi.Entry) *TransformData {
	return TransformComponent.Get(entry)
}

func GetPosition(entry *donburi.Entry) (float64, float64) {
	if transform := GetComponent(entry); transform != nil {
		return transform.X, transform.Y
	}
	return 0, 0
}

func SetPosition(entry *donburi.Entry, x, y float64) {
	if transform := GetComponent(entry); transform != nil && (transform.X != x || transform.Y != y) {
		transform.X = x
		transform.Y = y
		transform.IsDirty = true
	}
}

func GetRotation(entry *donburi.Entry) float64 {
	if transform := GetComponent(entry); transform != nil {
		return transform.Rotation
	}
	return 0
}

func SetRotation(entry *donburi.Entry, rotation float64) {
	if transform := GetComponent(entry); transform != nil && transform.Rotation != rotation {
		transform.Rotation = rotation
		transform.IsDirty = true
	}
}

func GetScale(entry *donburi.Entry) (float64, float64) {
	if transform := GetComponent(entry); transform != nil {
		return transform.ScaleX, transform.ScaleY
	}
	return 1, 1
}

func SetScale(entry *donburi.Entry, scaleX, scaleY float64) {
	if transform := GetComponent(entry); transform != nil && (transform.ScaleX != scaleX || transform.ScaleY != scaleY) {
		transform.ScaleX = scaleX
		transform.ScaleY = scaleY
		transform.IsDirty = true
	}
}

func GetOrigin(entry *donburi.Entry) (float64, float64) {
	if transform := GetComponent(entry); transform != nil {
		return transform.OriginX, transform.OriginY
	}
	return 0, 0
}

func SetOrigin(entry *donburi.Entry, originX, originY float64) {
	if transform := GetComponent(entry); transform != nil && (transform.OriginX != originX || transform.OriginY != originY) {
		transform.OriginX = originX
		transform.OriginY = originY
		transform.IsDirty = true
	}
}

func IsDirty(entry *donburi.Entry) bool {
	if transform := GetComponent(entry); transform != nil {
		return transform.IsDirty
	}
	return false
}

func SetDirty(entry *donburi.Entry, isDirty bool) {
	if transform := GetComponent(entry); transform != nil {
		transform.IsDirty = isDirty
	}
}

func GetMatrix(entry *donburi.Entry) ebiten.GeoM {
	if matrix := MatrixComponent.Get(entry); matrix != nil {
		if transform := GetComponent(entry); transform != nil && transform.IsDirty {
			matrix.m.Reset()
			matrix.m.Translate(-transform.OriginX, -transform.OriginY)
			matrix.m.Scale(transform.ScaleX, transform.ScaleY)
			matrix.m.Rotate(transform.Rotation)
			matrix.m.Translate(transform.X, transform.Y)
			transform.IsDirty = false
		}
		// Return a copy of the matrix to avoid external mutation
		return matrix.m
	}
	return ebiten.GeoM{}
}
