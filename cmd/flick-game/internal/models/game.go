package models

import "github.com/adm87/flick/pkg/ecs"

type GameModel struct {
	Factory  FactoryModel
	Renderer RendererModel
}

type FactoryModel struct {
	Image ecs.ECSFactoryType
}

type RendererModel struct {
	Image ecs.ECSRendererType
}

func NewGameModel() *GameModel {
	return &GameModel{
		Factory:  FactoryModel{},
		Renderer: RendererModel{},
	}
}
