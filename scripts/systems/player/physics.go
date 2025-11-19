package player

import (
	"math"

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

	inter, collided := nearestHit(ctx, bounds, horizontal, func(hitA, hitB collision.Hit) bool {
		return float32(math.Abs(float64(hitA.Delta[0]))) < float32(math.Abs(float64(hitB.Delta[0])))
	})

	if collided {
		vx = 0
		nx = nx + inter.Delta[0]
	}

	// =========== Vertical Movement and Collision ===========

	vy += player.Gravity() * float32(ctx.Time().FixedDeltaTime())
	ny := y + vy

	bounds = collider.Shape().Bounds(nx, ny)
	vertical := world.Check(ctx, bounds, collider.Layer(), models.SolidColliderType)

	player.SetOnGround(false)

	inter, collided = nearestHit(ctx, bounds, vertical, func(hitA, hitB collision.Hit) bool {
		return float32(math.Abs(float64(hitA.Delta[1]))) < float32(math.Abs(float64(hitB.Delta[1])))
	})

	if collided {
		if vy > 0 {
			player.SetOnGround(true)
		}
		vy = 0
		ny = ny + inter.Delta[1]
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

func nearestHit(ctx game.Context, bounds [4]float32, candidates []donburi.Entity, fn func(hitA, hitB collision.Hit) bool) (collision.Hit, bool) {
	var nearest collision.Hit
	var found bool

	for _, entity := range candidates {
		other := ctx.ECS().Entry(entity)

		shape := components.Collider.Get(other).Shape()
		x, y := components.Transform.Get(other).Position()

		if inter, collided := getInterection(ctx, bounds, x, y, shape); collided {
			if !found || fn(inter, nearest) {
				nearest = inter
				found = true
			}
		}
	}

	return nearest, found
}

func getInterection(ctx game.Context, bounds [4]float32, x, y float32, shape shapes.Shape) (collision.Hit, bool) {
	var intersection collision.Hit
	var collided bool

	switch shape.(type) {
	case *shapes.Rectangle:
		rect := shape.(*shapes.Rectangle)
		intersection, collided = collision.AABBvsAABB(bounds, rect.Bounds(x, y))
	case *shapes.Polygon:
		polygon := shape.(*shapes.Polygon)
		intersection, collided = collision.AABBvsPolygon(bounds, polygon, x, y)
	default:
		ctx.Log().Warn("unsupported shape type for collision detection")
	}

	return intersection, collided
}
