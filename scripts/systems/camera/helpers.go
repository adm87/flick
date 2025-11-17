package camera

import (
	"github.com/adm87/flick/scripts/actors"
	"github.com/adm87/flick/scripts/components"
	"github.com/adm87/flick/scripts/game"
	"github.com/yohamta/donburi"
)

// Viewport returns the current camera viewport bounds in world space.
func Viewport(ctx game.Context) [4]float32 {
	return ViewportOf(actors.Camera.MustFirst(ctx.ECS()))
}

// ViewportOf returns the current camera viewport bounds in world space for the given camera entry.
func ViewportOf(cameraEntry *donburi.Entry) [4]float32 {
	return components.Rectangle.Get(cameraEntry).Bounds(components.Transform.Get(cameraEntry).Position())
}
