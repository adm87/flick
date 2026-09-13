package images

import (
	"github.com/adm87/flick/pkg/logger"
	"github.com/hajimehoshi/ebiten/v2"
)

type ECSImageRenderer struct {
	logger logger.Logger
}

func NewECSImageRenderer(logger logger.Logger) *ECSImageRenderer {
	return &ECSImageRenderer{
		logger: logger,
	}
}

func (r *ECSImageRenderer) Render(target *ebiten.Image) error {
	return nil
}
