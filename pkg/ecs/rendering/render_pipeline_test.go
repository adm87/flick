package rendering_test

import (
	"math/rand/v2"
	"strconv"
	"testing"

	"github.com/adm87/flick/pkg/ecs"
	"github.com/adm87/flick/pkg/ecs/components/renderable"
	"github.com/adm87/flick/pkg/ecs/rendering"
	"github.com/adm87/flick/pkg/geom"
	"github.com/adm87/flick/pkg/logger"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
)

var benchSizes = []int{0, 10, 100, 1000, 10000}

type testrenderer struct{}

func (t *testrenderer) Render(*ebiten.Image, *rendering.RenderingCandidate, geom.Rect, ebiten.GeoM) error {
	return nil
}

type testSceneView struct{}

func (t *testSceneView) GetView() (geom.Rect, ebiten.GeoM) {
	return geom.Rect{Width: 800, Height: 600}, ebiten.GeoM{}
}

func BenchmarkRenderPipeline(b *testing.B) {
	for i := range benchSizes {
		size := benchSizes[i]
		b.Run("RenderPipeline/size="+strconv.Itoa(size), func(b *testing.B) {
			ecs := ecs.NewECS()
			renderPipeline := rendering.NewECSRenderPipeline(ecs, &testSceneView{}, logger.NewLogger(b.Output()))
			renderer := renderPipeline.RegisterRenderer(&testrenderer{})
			createBenchmarkEntities(ecs.World(), size, renderer)

			b.ReportAllocs()
			b.ResetTimer()

			for b.Loop() {
				renderPipeline.Draw(nil)
			}
		})
	}
}

func createBenchmarkEntities(world donburi.World, count int, renderer uint64) {
	entities := world.CreateMany(count,
		renderable.RenderableComponent,
	)
	z := rand.Perm(count)
	for i := range entities {
		entry := world.Entry(entities[i])

		r := renderable.RenderableComponent.Get(entry)
		r.SetRenderer(renderer)
		r.SetZIndex(int32(z[i]))
	}
}
