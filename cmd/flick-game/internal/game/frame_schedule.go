package game

import (
	"context"

	"github.com/adm87/flick/pkg/ecs"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/yohamta/donburi"
)

func setupFrameSchedule(frameSchedule *ecs.ECSFrameScheduler) {
	frameSchedule.AddUpdate(ecs.UpdatePhaseUpdate,
		func(ctx context.Context, world donburi.World, dt float64) error {
			// Add your update logic here

			if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
				return ebiten.Termination
			}
			return nil
		},
	)
	frameSchedule.AddUpdate(ecs.UpdatePhaseFixedUpdate,
		func(ctx context.Context, world donburi.World, dt float64) error {
			// Add your fixed update logic here
			return nil
		},
	)
	frameSchedule.AddUpdate(ecs.UpdatePhaseLateUpdate,
		func(ctx context.Context, world donburi.World, dt float64) error {
			// Add your late update logic here
			return nil
		},
	)
}
