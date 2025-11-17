package player

import (
	"github.com/adm87/flick/scripts/actors"
	"github.com/adm87/flick/scripts/collision"
	"github.com/adm87/flick/scripts/components"
	"github.com/adm87/flick/scripts/components/models"
	"github.com/adm87/flick/scripts/game"
	"github.com/yohamta/donburi"
)

func UpdatePhysics(ctx game.Context, world *collision.World) error {
	playerEntry := actors.Player.MustFirst(ctx.ECS())

	transform := components.Transform.Get(playerEntry)
	collider := components.Collider.Get(playerEntry)

	// Broad-phase collision detection
	bounds := collider.Shape().Bounds(transform.Position())
	solids := world.Check(ctx, bounds, models.SolidColliderType, collider.Layer())

	// Narrow-phase collision resolution
	if err := resolveSolidCollision(ctx, playerEntry, solids); err != nil {
		return err
	}

	return nil
}

func resolveSolidCollision(ctx game.Context, playerEntry *donburi.Entry, solids []donburi.Entity) error {
	for _, other := range solids {
		if other == playerEntry.Entity() {
			continue // Skip self
		}
		// TASK: Implement narrow-phase collision resolution between the player and solid entities.
	}
	return nil
}
