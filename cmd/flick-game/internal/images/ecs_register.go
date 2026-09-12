package images

import (
	"github.com/adm87/flick/cmd/flick-game/internal/models"
	"github.com/adm87/flick/pkg/ecs"
	"github.com/adm87/flick/pkg/logger"
)

// ECSRegister registers the ECS image renderer and factory with the ECS world and updates the game model accordingly.
func ECSRegister(e *ecs.ECS, gameModel *models.GameModel, logger logger.Logger) error {
	gameModel.Renderer.Image = e.RegisterRenderer(NewECSImageRenderer(logger))
	gameModel.Factory.Image = e.RegisterFactory(NewECSImageFactory(gameModel.Renderer.Image, logger))
	return nil
}
