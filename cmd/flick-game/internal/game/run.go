package game

import (
	"github.com/adm87/flick/cmd/flick-game/internal/images"
	"github.com/adm87/flick/cmd/flick-game/internal/models"
	"github.com/adm87/flick/pkg/ecs"
	"github.com/adm87/flick/pkg/game"
	"github.com/adm87/flick/pkg/logger"
)

func Run() error {
	model := models.NewGameModel()
	logger := logger.NewLogger()

	e, err := ecs.New(logger)
	if err != nil {
		return err
	}

	if err := register(e, model, logger); err != nil {
		return err
	}

	setupFrameSchedule(model, e.FrameScheduler())
	return game.Run(
		game.WithUpdate(e.FrameScheduler()),
		game.WithDraw(e.RenderPipeline()),
		game.WithLogger(logger),
	)
}

func register(e *ecs.ECS, model *models.GameModel, logger logger.Logger) error {
	if err := images.ECSRegister(e, model, logger); err != nil {
		return err
	}
	return nil
}
