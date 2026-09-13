package ecs

import (
	"github.com/adm87/flick/pkg/logger"
	"github.com/adm87/flick/pkg/structures/slotmap"
	"github.com/yohamta/donburi"
)

type ECS struct {
	world donburi.World

	frameScheduler *ECSFrameScheduler
	renderPipeline *ECSRenderPipeline

	factories *slotmap.SlotMap[ECSFactory]
	renderers *slotmap.SlotMap[ECSRenderer]

	logger logger.Logger
}

func New(logger logger.Logger) (*ECS, error) {
	e := &ECS{
		world:     donburi.NewWorld(),
		factories: slotmap.New[ECSFactory](16),
		renderers: slotmap.New[ECSRenderer](16),
		logger:    logger,
	}

	frameScheduler, err := NewECSFrameScheduler(e, logger)
	if err != nil {
		return nil, err
	}

	renderPipeline, err := NewECSRenderPipeline(e, logger)
	if err != nil {
		return nil, err
	}

	e.frameScheduler = frameScheduler
	e.renderPipeline = renderPipeline
	return e, nil
}

func (e *ECS) World() donburi.World {
	return e.world
}

func (e *ECS) FrameScheduler() *ECSFrameScheduler {
	return e.frameScheduler
}

func (e *ECS) RenderPipeline() *ECSRenderPipeline {
	return e.renderPipeline
}

// CreateEntity creates a new entity using the specified factory type. Returns nil if the factory is not registered.
func (e *ECS) CreateEntity(factoryType ECSFactoryType) *donburi.Entry {
	factory := e.GetFactory(factoryType)
	if factory == nil {
		return nil
	}
	return factory.CreateEntity(e.world)
}

// RegisterFactory registers a new factory and returns its unique ID.
func (e *ECS) RegisterFactory(factory ECSFactory) ECSFactoryType {
	return e.factories.Insert(factory)
}

// RegisterRenderer registers a new renderer and returns its unique ID.
func (e *ECS) RegisterRenderer(renderer ECSRenderer) ECSRendererType {
	return e.renderers.Insert(renderer)
}

// GetFactory retrieves a factory by its unique ID. Returns nil if the factory is not found.
func (e *ECS) GetFactory(factoryType ECSFactoryType) ECSFactory {
	factory, ok := e.factories.Get(slotmap.K(factoryType))
	if !ok {
		return nil
	}
	return factory
}

// GetRenderer retrieves a renderer by its unique ID. Returns nil if the renderer is not found.
func (e *ECS) GetRenderer(rendererType ECSRendererType) ECSRenderer {
	renderer, ok := e.renderers.Get(slotmap.K(rendererType))
	if !ok {
		return nil
	}
	return renderer
}
