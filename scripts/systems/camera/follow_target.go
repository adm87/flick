package camera

import (
	"github.com/adm87/flick/scripts/actors"
	"github.com/adm87/flick/scripts/components"
	"github.com/adm87/flick/scripts/game"
	"github.com/adm87/flick/scripts/math"
	"github.com/yohamta/donburi"
)

func FollowTarget(ctx game.Context, target *donburi.Entry) error {
	cameraEntry := actors.Camera.MustFirst(ctx.ECS())

	curX, curY := components.Transform.Get(cameraEntry).Position()
	tarX, tarY := components.Transform.Get(target).Position()

	x := math.Lerp(curX, tarX, 0.3)
	y := math.Lerp(curY, tarY, 0.3)

	if math.Distance(curX, curY, x, y) < 0.01 {
		x = tarX
		y = tarY
	}

	if x == curX && y == curY {
		return nil
	}

	components.Transform.Get(cameraEntry).SetPosition(x, y)
	return nil
}
