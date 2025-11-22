package player

import (
	"github.com/adm87/flick/scripts/actors"
	"github.com/adm87/flick/scripts/components"
	"github.com/adm87/flick/scripts/game"
	"github.com/adm87/flick/scripts/input"
	"github.com/adm87/flick/scripts/math"
)

const (
	WalkSpeed          = 1.0
	JumpStrength       = 2.9
	JumpCutMultiplier  = 0.3
	GroundAcceleration = 5.0
	AirAcceleration    = 3.0
	GroundDeceleration = 15.0
	AirDeceleration    = 1.0
	MinSpeedThreshold  = 0.01
)

var (
	Move = input.Action("move")
	Jump = input.Action("jump")
)

func UpdatePlayerInput(ctx game.Context) error {
	playerEntry := actors.Player.MustFirst(ctx.ECS())
	player := components.Player.Get(playerEntry)
	movement := components.Movement.Get(playerEntry)
	vx, vy := movement.Velocity()

	targetVx := float32(0)
	if move := ctx.Input().GetBinding(Move); move != nil {
		targetVx = move.Value() * WalkSpeed
	}

	onGround := player.OnGround()
	accelVx := GroundAcceleration
	if targetVx == 0 {
		accelVx = GroundDeceleration
		if !onGround {
			accelVx = AirDeceleration
		}
	} else if !onGround {
		accelVx = AirAcceleration
	}

	delta := float32(accelVx * ctx.Time().DeltaTime())
	vx = accelerateTowards(vx, targetVx, delta)
	if math.Abs(float64(vx)) < MinSpeedThreshold && targetVx == 0 {
		vx = 0
	}

	if jump := ctx.Input().GetBinding(Jump); jump != nil {
		if jump.JustActive() && player.CanJump() {
			vy = -JumpStrength
		} else if jump.JustInactive() && vy < 0 {
			vy *= JumpCutMultiplier
		}
	}

	movement.SetVelocity(vx, vy)
	return nil
}

func accelerateTowards(current, target, delta float32) float32 {
	diff := target - current
	if math.Abs(float64(diff)) <= float64(delta) {
		return target
	}
	return current + delta*math.Sign(diff)
}
