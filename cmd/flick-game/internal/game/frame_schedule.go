package game

import (
	"context"

	"github.com/adm87/flick/cmd/flick-game/internal/models"
	"github.com/adm87/flick/pkg/ecs"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/yohamta/donburi"
)

func setupFrameSchedule(model *models.GameModel, frameScheduler *ecs.ECSFrameScheduler) {
	_ = model

	frameScheduler.AddUpdate(ecs.UpdatePhaseUpdate,
		func(ctx context.Context, world donburi.World, dt float64) error {
			if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
				return ebiten.Termination
			}
			// Add your update logic here
			return nil
		},
	)
	frameScheduler.AddUpdate(ecs.UpdatePhaseFixedUpdate,
		func(ctx context.Context, world donburi.World, dt float64) error {
			// Add your fixed update logic here
			return nil
		},
	)
	frameScheduler.AddUpdate(ecs.UpdatePhaseLateUpdate,
		func(ctx context.Context, world donburi.World, dt float64) error {
			// Add your late update logic here
			return nil
		},
	)
}
