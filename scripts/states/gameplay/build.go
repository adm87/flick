package gameplay

import (
	"errors"

	"github.com/adm87/flick/data"
	"github.com/adm87/flick/scripts/actors"
	"github.com/adm87/flick/scripts/assets"
	"github.com/adm87/flick/scripts/collision"
	"github.com/adm87/flick/scripts/components"
	"github.com/adm87/flick/scripts/game"
	"github.com/adm87/flick/scripts/shapes"
	"github.com/adm87/tiled"
)

func (s *state) buildWorld(ctx game.Context) error {
	tmx, err := assets.Get[*tiled.Tmx](data.GymCollision)
	if err != nil {
		return err
	}

	if err := s.buildSolidWorld(ctx, tiled.ObjectGroupByName(tmx, "Solid")); err != nil {
		return err
	}
	if err := s.spawnPlayer(ctx, tiled.ObjectGroupByName(tmx, "Player")); err != nil {
		return err
	}

	return nil
}

func (s *state) buildSolidWorld(ctx game.Context, objects *tiled.ObjectGroup) error {
	for _, obj := range objects.Objects {
		var shape shapes.Shape

		if len(obj.Polygon.Points) > 0 {
			polygon := shapes.NewPolygon()
			polygon.SetVertices(shapes.GroupVertices(obj.Polygon.Points))
			shape = polygon
		} else {
			rectangle := shapes.NewRectangle()
			rectangle.SetSize(obj.Width, obj.Height)
			shape = rectangle
		}

		solid := actors.Spawn(ctx, actors.SolidActor)

		components.Transform.Get(solid).
			SetPosition(obj.X, obj.Y)
		components.Collider.Get(solid).
			SetType(collision.StaticCollisionType).
			SetShape(shape)

		s.world.Insert(solid.Entity(), shape.Bounds(obj.X, obj.Y))
	}
	return nil
}

func (s *state) spawnPlayer(ctx game.Context, objects *tiled.ObjectGroup) error {
	if len(objects.Objects) > 1 {
		return errors.New("ambiguous player spawn: multiple spawn points found")
	}

	for _, obj := range objects.Objects {
		player := actors.Spawn(ctx, actors.PlayerActor)

		center, bottom := obj.Width/2, obj.Height

		rectangle := shapes.NewRectangle().
			SetSize(obj.Width, obj.Height).
			SetPosition(-center, -bottom)

		components.Transform.Get(player).
			SetPosition(obj.X, obj.Y).
			SetOrigin(center, bottom)
		components.Collider.Get(player).
			SetType(collision.DynamicCollisionType).
			SetShape(rectangle)

		s.world.Insert(player.Entity(), rectangle.Bounds(obj.X, obj.Y))
		return nil
	}

	return nil
}
