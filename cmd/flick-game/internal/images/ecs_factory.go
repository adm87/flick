package images

import (
	"github.com/adm87/flick/pkg/ecs"
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

func (f *ECSImageFactory) CreateEntity() *donburi.Entry {
	return nil
}
