package player

import (
	"github.com/adm87/flick/scripts/actors"
	"github.com/adm87/flick/scripts/components"
	"github.com/adm87/flick/scripts/game"
	"github.com/hajimehoshi/ebiten/v2"
)

const playerWalkSpeed = 1.0
const playerJumpStrength = 3.0

func PollPlayerInput(ctx game.Context) error {
	playerEntry := actors.Player.MustFirst(ctx.ECS())

	movement := components.Movement.Get(playerEntry)
	vx, vy := movement.Velocity()

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		vx = -playerWalkSpeed
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		vx = playerWalkSpeed
	} else {
		vx = 0
	}

	if ebiten.IsKeyPressed(ebiten.KeyUp) && components.Player.Get(playerEntry).OnGround() {
		vy = -playerJumpStrength
	}

	movement.SetVelocity(vx, vy)

	return nil
}
