package images

import (
	"github.com/adm87/flick/cmd/flick-game/internal/models"
	"github.com/adm87/flick/pkg/ecs"
	"github.com/adm87/flick/pkg/logger"
)

// ECSRegister registers the ECS image renderer and factory with the ECS world and returns the corresponding model.
func ECSRegister(e *ecs.ECS, logger logger.Logger) (models.ImageModel, error) {
	rendererType := e.RegisterRenderer(NewECSImageRenderer(logger))
	factoryType := e.RegisterFactory(NewECSImageFactory(rendererType, logger))
	return models.ImageModel{
		FactoryType:  factoryType,
		RendererType: rendererType,
	}, nil
}
