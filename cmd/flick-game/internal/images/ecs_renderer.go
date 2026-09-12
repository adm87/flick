package images

import "github.com/adm87/flick/pkg/logger"

type ECSImageRenderer struct {
	logger logger.Logger
}

func NewECSImageRenderer(logger logger.Logger) *ECSImageRenderer {
	return &ECSImageRenderer{
		logger: logger,
	}
}
