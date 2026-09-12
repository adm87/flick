package ecs

import (
	"context"

	"github.com/adm87/flick/pkg/game"
	"github.com/yohamta/donburi"
)

type UpdatePhase int8

const (
	UpdatePhaseUpdate UpdatePhase = iota
	UpdatePhaseFixedUpdate
	UpdatePhaseLateUpdate
)

type FrameUpdate func(ctx context.Context, world donburi.World, dt float64) error

type ECSFrameScheduler struct {
	ecs          *ECS
	updates      []FrameUpdate
	fixedUpdates []FrameUpdate
	lateUpdates  []FrameUpdate
}

func NewECSFrameScheduler(e *ECS) *ECSFrameScheduler {
	if e == nil {
		panic("ECS cannot be nil")
	}
	return &ECSFrameScheduler{
		ecs: e,
	}
}

func (s *ECSFrameScheduler) AddUpdate(phase UpdatePhase, updates ...FrameUpdate) {
	switch phase {
	case UpdatePhaseUpdate:
		s.updates = append(s.updates, updates...)
	case UpdatePhaseFixedUpdate:
		s.fixedUpdates = append(s.fixedUpdates, updates...)
	case UpdatePhaseLateUpdate:
		s.lateUpdates = append(s.lateUpdates, updates...)
	}
}

func (s *ECSFrameScheduler) Update(ctx context.Context, t game.Time) error {
	world := s.ecs.World()
	if err := runUpdates(ctx, world, t.DeltaTime(), s.updates); err != nil {
		return err
	}
	for range t.FixedSteps() {
		if err := runUpdates(ctx, world, t.FixedDeltaTime(), s.fixedUpdates); err != nil {
			return err
		}
	}
	if err := runUpdates(ctx, world, t.DeltaTime(), s.lateUpdates); err != nil {
		return err
	}
	return nil
}

func runUpdates(ctx context.Context, world donburi.World, dt float64, updates []FrameUpdate) error {
	for i := range updates {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			if err := updates[i](ctx, world, dt); err != nil {
				return err
			}
		}
	}
	return nil
}
