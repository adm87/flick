package models

import (
	"github.com/adm87/flick/scripts/collision"
	"github.com/adm87/flick/scripts/shapes"
)

var DefaultCollider = Collider{
	layer: collision.DefaultCollisionLayer,
	cType: collision.StaticCollisionType,
	shape: shapes.NewRectangle(),
}

type Collider struct {
	layer collision.CollisionLayer
	cType collision.CollisionType
	shape shapes.Shape
}

func (c *Collider) Layer() collision.CollisionLayer {
	return c.layer
}

func (c *Collider) Type() collision.CollisionType {
	return c.cType
}

func (c *Collider) Shape() shapes.Shape {
	return c.shape
}

func (c *Collider) SetLayer(layer collision.CollisionLayer) *Collider {
	c.layer = layer
	return c
}

func (c *Collider) SetType(cType collision.CollisionType) *Collider {
	c.cType = cType
	return c
}

func (c *Collider) SetShape(shape shapes.Shape) *Collider {
	c.shape = shape
	return c
}
