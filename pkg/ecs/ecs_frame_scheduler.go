package ecs

import (
	"context"
	"errors"

	"github.com/adm87/flick/pkg/engine"
	"github.com/adm87/flick/pkg/logger"
)

type UpdatePhase int8

const (
	UpdatePhaseUpdate UpdatePhase = iota
	UpdatePhaseFixedUpdate
	UpdatePhaseLateUpdate
)

type FrameUpdate func(ctx context.Context, e *ECS, dt float64) error

type ECSFrameScheduler struct {
	logger       logger.Logger
	ecs          *ECS
	updates      []FrameUpdate
	fixedUpdates []FrameUpdate
	lateUpdates  []FrameUpdate
}

func NewECSFrameScheduler(e *ECS, logger logger.Logger) (*ECSFrameScheduler, error) {
	if e == nil {
		logger.Error("ECS cannot be nil")
		return nil, errors.New("ECS cannot be nil")
	}
	return &ECSFrameScheduler{
		ecs:    e,
		logger: logger,
	}, nil
}

func (s *ECSFrameScheduler) AddUpdate(phase UpdatePhase, updates ...FrameUpdate) {
	switch phase {
	case UpdatePhaseUpdate:
		s.updates = append(s.updates, updates...)
	case UpdatePhaseFixedUpdate:
		s.fixedUpdates = append(s.fixedUpdates, updates...)
	case UpdatePhaseLateUpdate:
		s.lateUpdates = append(s.lateUpdates, updates...)
	default:
		s.logger.Error("unknown update phase", logger.Int("phase", int(phase)))
	}
}

func (s *ECSFrameScheduler) Update(ctx context.Context, t engine.Time) error {
	if err := runUpdates(ctx, s.ecs, t.DeltaTime(), s.updates); err != nil {
		return err
	}
	for range t.FixedSteps() {
		if err := runUpdates(ctx, s.ecs, t.FixedDeltaTime(), s.fixedUpdates); err != nil {
			return err
		}
	}
	if err := runUpdates(ctx, s.ecs, t.DeltaTime(), s.lateUpdates); err != nil {
		return err
	}
	return nil
}

func runUpdates(ctx context.Context, e *ECS, dt float64, updates []FrameUpdate) error {
	for i := range updates {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			if err := updates[i](ctx, e, dt); err != nil {
				return err
			}
		}
	}
	return nil
}
