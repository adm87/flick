package gameplay

import (
	"errors"

	"github.com/adm87/flick/scripts/actors"
	"github.com/adm87/flick/scripts/assets"
	"github.com/adm87/flick/scripts/components"
	"github.com/adm87/flick/scripts/components/models"
	"github.com/adm87/flick/scripts/game"
	"github.com/adm87/flick/scripts/shapes"
	"github.com/adm87/tiled"
)

func (s *state) buildWorld(ctx game.Context) error {
	if err := s.buildSolidWorld(ctx, tiled.ObjectGroupByName(s.tilemap.Tmx, "Collision")); err != nil {
		return err
	}
	if err := s.spawnPlayer(ctx, s.tilemap.Tmx, tiled.ObjectGroupByName(s.tilemap.Tmx, "Player")); err != nil {
		return err
	}

	w, h := s.tilemap.Tmx.Width*s.tilemap.Tmx.TileWidth, s.tilemap.Tmx.Height*s.tilemap.Tmx.TileHeight

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
		SetSize(float32(w), float32(h))

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

func (s *state) spawnPlayer(ctx game.Context, tmx *tiled.Tmx, objects *tiled.ObjectGroup) error {
	if len(objects.Objects) > 1 {
		return errors.New("ambiguous player spawn: multiple spawn points found")
	}

	for _, obj := range objects.Objects {
		player := actors.Spawn(ctx, actors.PlayerActor)

		// Physical size of the player
		w := float32(tmx.TileWidth) * 0.5
		h := float32(tmx.TileHeight)

		set, gid, _ := tiled.TilesetByGID(tmx, obj.GID)
		tsx := assets.MustGet[*tiled.Tsx](assets.AssetHandle(set.Source))

		// Tile source position
		srcX := (int(gid) % int(tsx.Columns)) * int(tsx.TileWidth)
		srcY := (int(gid) / int(tsx.Columns)) * int(tsx.TileHeight)
		src := assets.AssetHandle(tsx.Image.Source)

		// center horizontally, align bottom
		toX := float32(tsx.TileOffset.X) - (float32(tsx.TileWidth)-w)/4
		toY := float32(tsx.TileOffset.Y) - (float32(tsx.TileHeight) - h)

		components.Tile.Get(player).
			SetPosition(srcX, srcY).
			SetSize(int(tsx.TileWidth), int(tsx.TileHeight)).
			SetOffset(toX, toY).
			SetGID(gid).
			SetSource(src)

		alignX, alignY := tiled.ObjectAlignmentAnchor(tsx.ObjectAlignment)

		x, y := obj.X, obj.Y

		orX := alignX * w
		orY := alignY * h

		components.Transform.Get(player).
			SetPosition(x, y).
			SetOrigin(orX, orY)

		rectangle := shapes.NewRectangle().
			SetSize(w, h).
			SetPosition(-orX, -orY)

		components.Collider.Get(player).
			SetType(models.DynamicColliderType).
			SetShape(rectangle)

		s.world.Insert(player.Entity(), rectangle.Bounds(x, y))
		return nil
	}

	return nil
}
