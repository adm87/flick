package internal

import (
	"os"

	"github.com/adm87/flick/cmd/flick-game/internal/game"
	"github.com/adm87/flick/cmd/flick-game/internal/images"
	"github.com/adm87/flick/cmd/flick-game/internal/models"
	"github.com/adm87/flick/pkg/ecs"
	"github.com/adm87/flick/pkg/engine"
	"github.com/adm87/flick/pkg/logger"
)

func Run() error {
	model := models.GameModel{}
	log := logger.NewLogger(os.Stdout)

	e, err := ecs.New(log)
	if err != nil {
		return err
	}

	if err := register(e, &model, log); err != nil {
		return err
	}

	game.SetupFrameSchedule(&model, e.FrameScheduler(), log)
	return engine.Run(
		engine.WithUpdate(e.FrameScheduler()),
		engine.WithDraw(e.RenderPipeline()),
		engine.WithLogger(log),
	)
}

func register(e *ecs.ECS, model *models.GameModel, log logger.Logger) error {
	imageModel, err := images.ECSRegister(e, log)
	if err != nil {
		return err
	}
	model.Image = imageModel
	return nil
}
