package internal

import (
	"math/rand"
	"os"
	"path/filepath"

	"github.com/adm87/flick/cmd/flick-game/internal/game"
	"github.com/adm87/flick/content"
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
	"github.com/yohamta/donburi"
)

func Run() error {
	log := logger.NewLogger(os.Stdout)
	res := resources.NewResources(log)

	dataStore := data.NewDataStore()
	dataImporter := data.NewDataImporter(dataStore)
	res.RegisterImporter(dataImporter, data.DataTypes())

	imageStore := images.NewImageStore()
	imageImporter := images.NewImageImporter(imageStore)
	imageRenderer := images.NewImageRenderer(imageStore)
	res.RegisterImporter(imageImporter, images.ImageTypes())

	cfg, err := loadConfig(res, dataStore)
	assert.NoError(err)

	ecs := ecs.NewECS()
	screen := game.NewScreen(cfg.Window.Width, cfg.Window.Height, log)
	view := game.NewView(ecs.World(), screen)
	rp := rendering.NewECSRenderPipeline(ecs, view, log)
	draw := game.NewDrawer(screen, view, ecs.World(), rp)
	update := game.NewUpdater(ecs.World())

	gameModel := &game.Model{
		Config: cfg,
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
	setupTestScene(ecs.World(), res, gameModel, view, gameAssets)

	return engine.Run(
		engine.WithLogger(log),
		engine.WithWindowSize(
			cfg.Window.Width,
			cfg.Window.Height,
		),
		engine.WithFullscreen(cfg.Window.Fullscreen),
		engine.WithUpdate(update),
		engine.WithDraw(draw),
		engine.WithLayout(screen),
	)
}

func loadConfig(res *resources.Resources, store *data.DataStore) (*game.Config, error) {
	err := res.Load(content.EmbeddedFS(), content.ConfResourcePath)
	if err != nil {
		return nil, err
	}

	defer func() {
		err := res.Unload(content.ConfResourcePath)
		assert.NoError(err)
	}()

	handle, err := res.GetHandle(content.ConfResourcePath)
	if err != nil {
		return nil, err
	}

	raw, err := store.Get(handle)
	if err != nil {
		return nil, err
	}

	return game.NewConfig(raw)
}

func setupTestScene(world donburi.World, res *resources.Resources, gameModel *game.Model, view *game.View, assets *game.Assets) {
	const TestImage resources.ResourcePath = "tile_0105.png"

	path, err := filepath.Abs("../../content")
	assert.NoError(err)

	testFS := os.DirFS(path + "/resources")

	err = res.Load(testFS, TestImage)
	assert.NoError(err)

	handle, err := res.GetHandle(TestImage)
	assert.NoError(err)

	image, err := assets.Images.Get(handle)
	assert.NoError(err)

	entities := world.CreateMany(100,
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
			rand.Float64()*float64(gameModel.Config.Window.Width),
			rand.Float64()*float64(gameModel.Config.Window.Height),
		)

		b, _ := transform.GetBounds(entry)
		b.SetSize(
			float64(image.Bounds().Dx()),
			float64(image.Bounds().Dy()),
		)

		r, _ := renderable.GetRenderable(entry)
		r.SetRenderer(gameModel.Renderers.ImageRendererID)

		img, _ := images.GetImage(entry)
		img.SetHandle(handle)
	}

	camEntry := world.Entry(world.Create(
		camera.MainCamera,
		transform.TransformComponent,
		transform.MatrixComponent,
	))

	view.SetCamera(camEntry)
}
