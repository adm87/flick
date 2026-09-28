package internal

import (
	"os"

	"github.com/adm87/flick/cmd/flick-game/internal/game"
	"github.com/adm87/flick/content"
	"github.com/adm87/flick/pkg/aseprite"
	"github.com/adm87/flick/pkg/assert"
	"github.com/adm87/flick/pkg/data"
	"github.com/adm87/flick/pkg/ecs"
	"github.com/adm87/flick/pkg/ecs/components/camera"
	"github.com/adm87/flick/pkg/ecs/components/renderable"
	"github.com/adm87/flick/pkg/ecs/components/transform"
	"github.com/adm87/flick/pkg/ecs/rendering"
	"github.com/adm87/flick/pkg/engine"
	"github.com/adm87/flick/pkg/images"
	"github.com/adm87/flick/pkg/logger"
	"github.com/adm87/flick/pkg/resources"
	"github.com/adm87/flick/pkg/types/geom"
	"github.com/alexflint/go-arg"
	"github.com/yohamta/donburi"

	asepritegen "github.com/adm87/flick/generated/aseprite"
)

// TODO: Revisit for production
type cmdArgs struct {
	WorkingDir string `arg:"--working-dir" help:"directory to operate in"`
}

func Run() error {
	var v cmdArgs
	arg.MustParse(&v)
	return run(v)
}

func run(v cmdArgs) error {
	log := logger.NewLogger(os.Stdout)
	res := resources.NewResources(log)

	gameConfig, err := game.NewConfig(content.GameConfig)
	assert.NoError(err)

	asepriteConfig, err := aseprite.LoadConfig(content.AsepriteConfig)
	assert.NoError(err)

	dataStore := data.NewDataStore()
	dataImporter := data.NewDataImporter(dataStore)
	res.RegisterImporter(dataImporter, data.DataTypes())

	imageStore := images.NewImageStore()
	imageImporter := images.NewImageImporter(imageStore)
	imageRenderer := images.NewImageRenderer(imageStore)
	res.RegisterImporter(imageImporter, images.ImageTypes())

	ecs := ecs.NewECS()
	screen := game.NewScreen(
		gameConfig.Window.Width,
		gameConfig.Window.Height,
	)
	view := game.NewView(screen)
	rp := rendering.NewECSRenderPipeline(ecs, view, log)
	draw := game.NewDrawer(screen, view, ecs.World(), rp)
	update := game.NewUpdater(ecs.World())
	ase := aseprite.New(
		asepriteConfig,
		dataStore,
		imageStore,
		v.WorkingDir,
	)

	gameModel := &game.Model{
		Config: gameConfig,
		Renderers: &game.Renderers{
			ImageRendererID: rp.RegisterRenderer(imageRenderer),
		},
	}

	gameAssets := &game.Assets{
		Resources: res,
		Data:      dataStore,
		Images:    imageStore,
	}

	// Temporary test scene setup
	setupTestScene(ecs.World(), res, view, gameModel, gameAssets, ase)

	return engine.Run(
		engine.WithLogger(log),
		engine.WithWindowSize(
			gameConfig.Window.Width,
			gameConfig.Window.Height,
		),
		engine.WithFullscreen(gameConfig.Window.Fullscreen),
		engine.WithCursorMode(gameConfig.CursorMode),
		engine.WithUpdate(update),
		engine.WithDraw(draw),
		engine.WithLayout(screen),
	)
}

func setupTestScene(world donburi.World, res *resources.Resources, view *game.View, gameModel *game.Model, gameAssets *game.Assets, ase *aseprite.Aseprite) {
	err := res.Load(ase.ContentFS(), asepritegen.CaptainImagePath, asepritegen.CaptainJsonPath)
	assert.NoError(err)

	err = ase.BuildLibrary(res, asepritegen.CaptainLibrary)
	assert.NoError(err)

	imgHandle, err := res.GetHandle(asepritegen.CaptainImagePath)
	assert.NoError(err)

	entities := world.CreateMany(1,
		transform.BoundsComponent,
		transform.MatrixComponent,
		transform.TransformComponent,
		renderable.RenderableComponent,
		images.ImageComponent,
	)

	frame, err := gameAssets.Images.GetFrame(imgHandle, 0)
	assert.NoError(err)

	for i := range entities {
		entry := world.Entry(entities[i])

		r, _ := renderable.GetRenderable(entry)
		r.SetRenderer(gameModel.Renderers.ImageRendererID)

		ax, ay := 0.5, 1.0

		img, _ := images.GetImage(entry)
		img.SetAnchor(geom.Vec2{X: ax, Y: ay})
		img.SetHandle(imgHandle)

		w, h := float64(frame.Bounds().Dx()), float64(frame.Bounds().Dy())

		b, _ := transform.GetBounds(entry)
		b.SetPosition(-w*ax, -h*ay)
		b.SetSize(w, h)
	}

	camEntry := world.Entry(world.Create(
		camera.MainCamera,
		transform.TransformComponent,
		transform.MatrixComponent,
	))

	view.SetCamera(camEntry)
}
