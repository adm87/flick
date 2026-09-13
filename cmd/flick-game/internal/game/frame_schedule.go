package game

import (
	"context"

	"github.com/adm87/flick/cmd/flick-game/internal/models"
	"github.com/adm87/flick/pkg/ecs"
	"github.com/adm87/flick/pkg/logger"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func SetupFrameSchedule(model *models.GameModel, frameScheduler *ecs.ECSFrameScheduler, log logger.Logger) {
	frameScheduler.AddUpdate(ecs.UpdatePhaseUpdate,
		func(ctx context.Context, e *ecs.ECS, dt float64) error {
			if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
				return ebiten.Termination
			}
			return nil
		},
	)
	frameScheduler.AddUpdate(ecs.UpdatePhaseFixedUpdate,
		func(ctx context.Context, e *ecs.ECS, dt float64) error {
			return nil
		},
	)
	frameScheduler.AddUpdate(ecs.UpdatePhaseLateUpdate,
		func(ctx context.Context, e *ecs.ECS, dt float64) error {
			return nil
		},
	)
}
