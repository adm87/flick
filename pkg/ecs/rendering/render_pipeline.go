package rendering

import (
	"errors"
	"fmt"

	"github.com/adm87/flick/pkg/ecs"
	"github.com/adm87/flick/pkg/ecs/components/renderable"
	"github.com/adm87/flick/pkg/ecs/components/transform"
	"github.com/adm87/flick/pkg/geom"
	"github.com/adm87/flick/pkg/logger"
	"github.com/adm87/flick/pkg/structures/slotmap"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

var (
	ErrViewportNotValid = errors.New("viewport is not valid")
	ErrRenderingError   = errors.New("render pipeline error")
)

type View interface {
	GetView() (geom.Rect, ebiten.GeoM)
}

type ECSRenderPipeline struct {
	ecs            *ecs.ECS
	query          *donburi.Query
	renderers      *slotmap.SlotMap[Renderer]
	renderingQueue *RenderingQueue
	view           View
	logger         logger.Logger
}

func NewECSRenderPipeline(ecs *ecs.ECS, view View, log logger.Logger) *ECSRenderPipeline {
	rp := &ECSRenderPipeline{
		ecs: ecs,
		query: donburi.NewQuery(
			filter.Contains(
				renderable.RenderableComponent,
			),
		),
		renderers:      slotmap.New[Renderer](10),
		renderingQueue: NewRenderingQueue(),
		view:           view,
		logger:         log.With(logger.String("source", "RenderPipeline")),
	}
	return rp
}

func (rp *ECSRenderPipeline) RegisterRenderer(renderer Renderer) RendererK {
	k := rp.renderers.Insert(renderer)
	return k.Pack()
}

func (rp *ECSRenderPipeline) Draw(target *ebiten.Image) error {
	viewport, viewmatrix := rp.view.GetView()
	if !(viewport.Area() > 0) {
		return fmt.Errorf("%w: %v", ErrViewportNotValid, viewport)
	}

	rp.renderingQueue.Reset()

	// TODO: Replace with a spatial partitioning query for better performance
	rp.query.Each(rp.ecs.World(), func(entry *donburi.Entry) {
		r := renderable.RenderableComponent.Get(entry)
		if !r.IsVisible() {
			return
		}
		bounds, ok := transform.GetWorldBounds(entry)
		if ok && !bounds.Intersects(viewport) {
			return
		}
		rp.renderingQueue.Enqueue(entry, r)
	})

	candidates := rp.renderingQueue.Candidates()
	keys := rp.renderingQueue.sortedKeys()

	var errs []error

	// TODO: Replace with rendering jobs for batching
	for i := range keys {
		rc := candidates[keys[i].index()]
		k := slotmap.Unpack(rc.Renderable.Renderer())

		renderer, err := rp.renderers.Get(k)
		if err != nil {
			errs = append(errs, err)
			continue
		}

		if err := renderer.Render(target, rc, viewport, viewmatrix); err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		return fmt.Errorf("%w: %w", ErrRenderingError, errors.Join(errs...))
	}

	return nil
}
