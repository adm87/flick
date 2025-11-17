package actors

import (
	"github.com/adm87/flick/scripts/components"
	"github.com/adm87/flick/scripts/game"
	"github.com/yohamta/donburi"
)

var (
	// WorldBoundsActor defines the archetype for the world bounds entity
	WorldBoundsActor = NewActorType(WorldBounds,
		components.Transform,
		components.Rectangle,
	)

	// CameraActor defines the archetype for the camera entity
	CameraActor = NewActorType(Camera,
		components.Transform,
		components.Rectangle,
	)

	// DebugActor defines the archetype for the debug entity
	DebugActor = NewActorType(Debug,
		components.Debug,
	)

	// PlayerActor defines the archetype for the player entity
	PlayerActor = NewActorType(Player,
		components.Collider,
		components.Movement,
		components.Player,
		components.Transform,
	)

	// SolidActor defines the archetype for solid entities
	SolidActor = NewActorType(Solid,
		components.Collider,
		components.Transform,
	)
)

// ActorType represents a collection of component types that define an actor's structure.
type ActorType []donburi.IComponentType

// NewActorType creates a new ActorType with the given tag and component types.
func NewActorType(tag *donburi.ComponentType[donburi.Tag], componentTypes ...donburi.IComponentType) ActorType {
	return ActorType(append([]donburi.IComponentType{tag}, componentTypes...))
}

// Spawn creates a new entity in the game world with the specified ActorType and returns its entry.
func Spawn(ctx game.Context, actorType ActorType) *donburi.Entry {
	return ctx.ECS().Entry(ctx.ECS().Create(actorType...))
}

// Despawn removes the specified entity from the game world.
func Despawn(ctx game.Context, entity donburi.Entity) {
	ctx.ECS().Remove(entity)
}
