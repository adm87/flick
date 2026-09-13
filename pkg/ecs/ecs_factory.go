package ecs

import (
	"github.com/adm87/flick/pkg/structures/slotmap"
	"github.com/yohamta/donburi"
)

type ECSFactoryType = slotmap.K

type ECSFactory interface {
	CreateEntity(world donburi.World) *donburi.Entry
}
