package internal

import (
	"math/rand"
	"os"

	"github.com/adm87/flick/cmd/flick-game/internal/game"
	"github.com/adm87/flick/cmd/flick-game/internal/models"
	"github.com/adm87/flick/pkg/ecs"
	"github.com/adm87/flick/pkg/ecs/components/camera"
	"github.com/adm87/flick/pkg/ecs/components/renderable"
	"github.com/adm87/flick/pkg/ecs/components/transform"
	"github.com/adm87/flick/pkg/ecs/rendering"
	"github.com/adm87/flick/pkg/engine"
	"github.com/adm87/flick/pkg/images"
	"github.com/adm87/flick/pkg/logger"
)

const (
	count = 100

	screenWidth  = 800
	screenHeight = 600
)

func Run() error {
	ecs := ecs.NewECS()
	world := ecs.World()
	view := game.NewView(world)
	log := logger.NewLogger(os.Stdout)
	renderPipeline := rendering.NewRenderPipeline(ecs, view, log)

	gameModel := &models.GameModel{
		Renderers: models.Renderers{
			ImageRenderer: renderPipeline.RegisterRenderer(images.NewImageRenderer()),
		},
	}

	entities := world.CreateMany(count,
		transform.BoundsComponent,
		transform.MatrixComponent,
		transform.TransformComponent,
		renderable.RenderableComponent,
		images.ImageComponent,
	)

	for i := range entities {
		entry := world.Entry(entities[i])

		tr, _ := transform.GetTransform(entry)
		tr.SetPosition(
			rand.Float64()*800,
			rand.Float64()*600,
		)

		b, _ := transform.GetBounds(entry)
		b.SetSize(20, 20)

		r, _ := renderable.GetRenderable(entry)
		r.SetRenderer(gameModel.Renderers.ImageRenderer)
	}

	camEntry := world.Entry(world.Create(
		camera.CameraComponent,
		camera.MainCamera,
		transform.MatrixComponent,
		transform.BoundsComponent,
	))

	b, _ := transform.GetBounds(camEntry)
	b.SetSize(screenWidth, screenHeight)

	view.SetCamera(camEntry)

	return engine.Run(
		engine.WithLogger(log),
		engine.WithWindowSize(screenWidth, screenHeight),
		engine.WithDraw(renderPipeline),
	)
}
