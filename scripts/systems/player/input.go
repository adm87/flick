package player

import (
	"github.com/adm87/flick/scripts/actors"
	"github.com/adm87/flick/scripts/components"
	"github.com/adm87/flick/scripts/game"
	"github.com/adm87/flick/scripts/input"
)

const playerWalkSpeed = 1.0
const playerJumpStrength = 3.0

var (
	Move = input.Action("move")
	Jump = input.Action("jump")
)

func UpdatePlayerInput(ctx game.Context) error {
	playerEntry := actors.Player.MustFirst(ctx.ECS())
	player := components.Player.Get(playerEntry)

	movement := components.Movement.Get(playerEntry)
	vx, vy := movement.Velocity()

	// Handle Movement Input
	if move := ctx.Input().GetBinding(Move); move != nil {
		vx = move.Value() * playerWalkSpeed
	}

	// Handle Jump Input
	if player.OnGround() {
		if jump := ctx.Input().GetBinding(Jump); jump != nil && jump.IsActive() {
			vy = -playerJumpStrength
		}
	}

	movement.SetVelocity(vx, vy)

	return nil
}
