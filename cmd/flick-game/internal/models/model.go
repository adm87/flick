package models

import (
	"github.com/adm87/flick/pkg/ecs"
)

type GameModel struct {
	Image ImageModel
}

type ImageModel struct {
	FactoryType  ecs.ECSFactoryType
	RendererType ecs.ECSRendererType
}
