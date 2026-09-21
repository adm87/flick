package ecs

import "github.com/yohamta/donburi"

type ECS struct {
	world donburi.World
}

func NewECS() *ECS {
	return &ECS{
		world: donburi.NewWorld(),
	}
}

func (e *ECS) World() donburi.World {
	return e.world
}
