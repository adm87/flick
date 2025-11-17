package camera

import (
	"github.com/adm87/flick/scripts/actors"
	"github.com/adm87/flick/scripts/components"
	"github.com/adm87/flick/scripts/game"
)

// Viewport returns the current camera viewport bounds in world space.
func Viewport(ctx game.Context) [4]float32 {
	cameraEntry := actors.Camera.MustFirst(ctx.ECS())
	return components.Rectangle.Get(cameraEntry).Bounds(components.Transform.Get(cameraEntry).Position())
}
