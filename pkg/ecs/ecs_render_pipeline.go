package ecs

import (
	"errors"

	"github.com/adm87/flick/pkg/logger"
	"github.com/adm87/flick/pkg/structures/slotmap"
	"github.com/hajimehoshi/ebiten/v2"
)

type ECSRendererType = slotmap.K

type ECSRenderer interface {
}

type ECSRenderPipeline struct {
	ecs *ECS

	logger logger.Logger
}

func NewECSRenderPipeline(e *ECS, logger logger.Logger) (*ECSRenderPipeline, error) {
	if e == nil {
		logger.Error("ECS cannot be nil")
		return nil, errors.New("ECS cannot be nil")
	}
	return &ECSRenderPipeline{
		ecs:    e,
		logger: logger,
	}, nil
}

func (p *ECSRenderPipeline) Draw(target *ebiten.Image) error {
	return nil
}
