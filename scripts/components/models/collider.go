package models

import (
	"github.com/adm87/flick/scripts/shapes"
)

// =========== Collision Layer ==========

type CollisionLayer uint8

const (
	MaxCollisionLayers int = 32

	NoCollisionLayer      CollisionLayer = 0
	DefaultCollisionLayer CollisionLayer = iota
)

var nameByLayer = map[CollisionLayer]string{
	DefaultCollisionLayer: "Default",
}

func NewLayer(name string) CollisionLayer {
	if len(nameByLayer) >= MaxCollisionLayers {
		panic("maximum number of collision layers exceeded")
	}

	layer := CollisionLayer(len(nameByLayer))
	nameByLayer[layer] = name

	return layer
}

func (l CollisionLayer) String() string {
	if name, ok := nameByLayer[l]; ok {
		return name
	}
	return "unknown"
}

func (l CollisionLayer) IsValid() bool {
	_, ok := nameByLayer[l]
	return ok
}

func NameByLayer(layer CollisionLayer) (string, bool) {
	name, ok := nameByLayer[layer]
	return name, ok
}

type ColliderType uint8

const (
	NoColliderType ColliderType = iota
	SolidColliderType
	SlopeColliderType
	DynamicColliderType
)

var DefaultCollider = Collider{
	layer: DefaultCollisionLayer,
	cType: SolidColliderType,
	shape: shapes.NewRectangle(),
}

type Collider struct {
	layer CollisionLayer
	cType ColliderType
	shape shapes.Shape
}

func (c *Collider) Layer() CollisionLayer {
	return c.layer
}

func (c *Collider) Type() ColliderType {
	return c.cType
}

func (c *Collider) Shape() shapes.Shape {
	return c.shape
}

func (c *Collider) SetLayer(layer CollisionLayer) *Collider {
	c.layer = layer
	return c
}

func (c *Collider) SetType(cType ColliderType) *Collider {
	c.cType = cType
	return c
}

func (c *Collider) SetShape(shape shapes.Shape) *Collider {
	c.shape = shape
	return c
}
