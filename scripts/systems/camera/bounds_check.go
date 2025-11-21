package camera

import (
	"math"

	"github.com/adm87/flick/scripts/actors"
	"github.com/adm87/flick/scripts/components"
	"github.com/adm87/flick/scripts/game"
)

func BoundsCheck(ctx game.Context, bounds [4]float32) error {
	cameraEntry := actors.Camera.MustFirst(ctx.ECS())
	viewport := ViewportOf(cameraEntry)

	halfWidth := (viewport[2] - viewport[0]) / 2
	halfHeight := (viewport[3] - viewport[1]) / 2

	x, y := components.Transform.Get(cameraEntry).Position()

	if x-halfWidth < bounds[0] {
		x = bounds[0] + halfWidth
	} else if x+halfWidth > bounds[2] {
		x = bounds[2] - halfWidth
	}

	if y-halfHeight < bounds[1] {
		y = bounds[1] + halfHeight
	} else if y+halfHeight > bounds[3] {
		y = bounds[3] - halfHeight
	}

	x = float32(math.Round(float64(x))) + 0.5
	y = float32(math.Round(float64(y))) + 0.5

	components.Transform.Get(cameraEntry).SetPosition(x, y)
	return nil
}
