package player

import (
	"github.com/adm87/flick/scripts/actors"
	"github.com/adm87/flick/scripts/collision"
	"github.com/adm87/flick/scripts/components"
	"github.com/adm87/flick/scripts/components/models"
	"github.com/adm87/flick/scripts/game"
	"github.com/adm87/flick/scripts/shapes"
	"github.com/yohamta/donburi"
)

func UpdatePhysics(ctx game.Context, world *collision.World) error {
	playerEntry := actors.Player.MustFirst(ctx.ECS())

	collider := components.Collider.Get(playerEntry)
	transform := components.Transform.Get(playerEntry)
	movement := components.Movement.Get(playerEntry)
	player := components.Player.Get(playerEntry)

	x, y := transform.Position()
	vx, vy := movement.Velocity()

	// =========== Horizontal Movement and Collision ===========

	nx := x + vx

	bounds := collider.Shape().Bounds(nx, y)
	horizontal := world.Check(ctx, bounds, collider.Layer(), models.SolidColliderType)

	if inter, collided := nearestHorizontalCollision(ctx, bounds, horizontal); collided {
		vx = 0
		nx = nx + inter.OverlapX
	}

	// =========== Vertical Movement and Collision ===========

	vy += player.Gravity() * float32(ctx.Time().FixedDeltaTime())
	ny := y + vy

	bounds = collider.Shape().Bounds(nx, ny)
	vertical := world.Check(ctx, bounds, collider.Layer(), models.SolidColliderType)

	player.SetOnGround(false)

	if inter, collided := nearestVerticalCollision(ctx, bounds, vertical); collided {
		if vy > 0 {
			player.SetOnGround(true)
		}
		vy = 0
		ny = ny + inter.OverlapY
	}

	// =========== Apply Movement ===========

	transform.SetPosition(nx, ny)
	movement.SetVelocity(vx, vy)

	if x != nx || y != ny {
		// Update the player's position in the collision world if it has moved
		world.Reinsert(playerEntry.Entity(), collider.Shape().Bounds(nx, ny))
	}

	return nil
}

func nearestHorizontalCollision(ctx game.Context, bounds [4]float32, candidates []donburi.Entity) (*collision.Intersection, bool) {
	var nearest *collision.Intersection

	for _, entity := range candidates {
		other := ctx.ECS().Entry(entity)

		shape := components.Collider.Get(other).Shape()
		x, y := components.Transform.Get(other).Position()

		if inter, collided := getInterection(ctx, bounds, x, y, shape); collided {
			if nearest == nil || inter.OverlapX < nearest.OverlapX {
				nearest = inter
			}
		}
	}

	return nearest, nearest != nil
}

func nearestVerticalCollision(ctx game.Context, bounds [4]float32, candidates []donburi.Entity) (*collision.Intersection, bool) {
	var nearest *collision.Intersection

	for _, entity := range candidates {
		other := ctx.ECS().Entry(entity)

		shape := components.Collider.Get(other).Shape()
		x, y := components.Transform.Get(other).Position()

		if inter, collided := getInterection(ctx, bounds, x, y, shape); collided {
			if nearest == nil || inter.OverlapY < nearest.OverlapY {
				nearest = inter
			}
		}
	}

	return nearest, nearest != nil
}

func getInterection(ctx game.Context, bounds [4]float32, x, y float32, shape shapes.Shape) (*collision.Intersection, bool) {
	var intersection *collision.Intersection
	var collided bool

	switch shape.(type) {
	case *shapes.Rectangle:
		rect := shape.(*shapes.Rectangle)
		intersection, collided = collision.AABBvsAABB(bounds, rect.Bounds(x, y))
	case *shapes.Polygon:
		// polygon := shape.(*shapes.Polygon)
		// intersection, collided = collision.AABBvsPolygon(bounds, polygon, x, y)
	default:
		ctx.Log().Warn("unsupported shape type for collision detection")
	}

	return intersection, collided
}
