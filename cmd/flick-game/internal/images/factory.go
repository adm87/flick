package images

import (
	"github.com/adm87/flick/pkg/ecs"
	"github.com/adm87/flick/pkg/ecs/components/renderable"
	"github.com/adm87/flick/pkg/ecs/components/transform"
	"github.com/adm87/flick/pkg/logger"
	"github.com/yohamta/donburi"
)

type ECSImageFactory struct {
	logger       logger.Logger
	rendererType ecs.ECSRendererType
}

func NewECSImageFactory(rendererType ecs.ECSRendererType, logger logger.Logger) *ECSImageFactory {
	return &ECSImageFactory{
		rendererType: rendererType,
		logger:       logger,
	}
}

func (f *ECSImageFactory) CreateEntity(world donburi.World) *donburi.Entry {
	entry := world.Entry(world.Create(
		transform.TransformComponent,
		transform.MatrixComponent,
		renderable.RenderableComponent,
		ImageComponent,
	))
	renderable.SetType(entry, uint64(f.rendererType))
	return entry
}
