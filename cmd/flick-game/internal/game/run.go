package game

import (
	"github.com/adm87/flick/pkg/ecs"
	"github.com/adm87/flick/pkg/game"
)

func Run() error {
	e := ecs.New()

	frameSchedule := ecs.NewECSFrameScheduler(e)
	setupFrameSchedule(frameSchedule)

	renderPipeline := ecs.NewECSRenderPipeline(e)

	return game.Run(
		game.WithUpdate(frameSchedule),
		game.WithDraw(renderPipeline),
	)
}
