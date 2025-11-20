package gameplay

import (
	"errors"

	"github.com/adm87/flick/data"
	"github.com/adm87/flick/scripts/actors"
	"github.com/adm87/flick/scripts/assets"
	"github.com/adm87/flick/scripts/components"
	"github.com/adm87/flick/scripts/components/models"
	"github.com/adm87/flick/scripts/game"
	"github.com/adm87/flick/scripts/shapes"
	"github.com/adm87/tiled"
)

func (s *state) buildWorld(ctx game.Context) error {
	tmx, err := assets.Get[*tiled.Tmx](data.TilemapExampleA)
	if err != nil {
		return err
	}

	if err := s.buildSolidWorld(ctx, tiled.ObjectGroupByName(tmx, "Collision")); err != nil {
		return err
	}
	if err := s.spawnPlayer(ctx, tiled.ObjectGroupByName(tmx, "Player")); err != nil {
		return err
	}

	actors.Spawn(ctx, actors.DebugActor)
	actors.Spawn(ctx, actors.WorldBoundsActor)

	cameraEntry := actors.Spawn(ctx, actors.CameraActor)

	halfWidth := float32(ctx.Screen().Width) / 2
	halfHeight := float32(ctx.Screen().Height) / 2

	// Center camera origin
	// This will center the world around the camera position
	components.Transform.Get(cameraEntry).
		SetOrigin(halfWidth, halfHeight)

	// Set Camera viewport and local position
	// We need to shift so that the camera's position represents the center of the screen
	components.Rectangle.Get(cameraEntry).
		SetSize(ctx.Screen().Width, ctx.Screen().Height).
		SetPosition(-halfWidth, -halfHeight)

	// Set world bounds size
	components.Rectangle.Get(actors.WorldBounds.MustFirst(ctx.ECS())).
		SetSize(float32(tmx.Width*tmx.TileWidth), float32(tmx.Height*tmx.TileHeight))

	return nil
}

func (s *state) buildSolidWorld(ctx game.Context, objects *tiled.ObjectGroup) error {
	for _, obj := range objects.Objects {
		var shape shapes.Shape

		colliderType := models.SolidColliderType

		if len(obj.Polygon.Points) > 0 {
			polygon := shapes.NewPolygon()
			polygon.SetVertices(shapes.GroupVertices(obj.Polygon.Points))
			shape = polygon
			colliderType = models.SlopeColliderType
		} else {
			rectangle := shapes.NewRectangle()
			rectangle.SetSize(obj.Width, obj.Height)
			shape = rectangle
		}

		solid := actors.Spawn(ctx, actors.SolidActor)

		components.Transform.Get(solid).
			SetPosition(obj.X, obj.Y)
		components.Collider.Get(solid).
			SetType(colliderType).
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

		components.Transform.Get(player).
			SetPosition(obj.X, obj.Y).
			SetOrigin(center, bottom)

		rectangle := shapes.NewRectangle().
			SetSize(obj.Width, obj.Height).
			SetPosition(-center, -bottom)

		components.Collider.Get(player).
			SetType(models.DynamicColliderType).
			SetShape(rectangle)

		s.world.Insert(player.Entity(), rectangle.Bounds(obj.X, obj.Y))
		return nil
	}

	return nil
}
