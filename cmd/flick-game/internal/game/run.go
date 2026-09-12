package game

import (
	"github.com/adm87/flick/pkg/ecs"
	"github.com/adm87/flick/pkg/game"
	"github.com/adm87/flick/pkg/logger"
)

func Run() error {
	logger := logger.NewLogger()

	e := ecs.New()

	frameScheduler := ecs.NewECSFrameScheduler(e)
	setupFrameSchedule(frameScheduler)

	renderPipeline := ecs.NewECSRenderPipeline(e)

	return game.Run(
		game.WithUpdate(frameScheduler),
		game.WithDraw(renderPipeline),
		game.WithLogger(logger),
	)
}
