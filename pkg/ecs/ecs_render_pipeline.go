package ecs

import "github.com/hajimehoshi/ebiten/v2"

type ECSRenderPipeline struct {
	ecs *ECS
}

func NewECSRenderPipeline(e *ECS) *ECSRenderPipeline {
	return &ECSRenderPipeline{
		ecs: e,
	}
}

func (p *ECSRenderPipeline) Draw(target *ebiten.Image) {

}
