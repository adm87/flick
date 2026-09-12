package game

import (
	"context"
	"errors"

	"github.com/adm87/flick/pkg/logger"
	"github.com/hajimehoshi/ebiten/v2"
)

type shell struct {
	ctx      context.Context
	update   GameUpdate
	draw     GameDraw
	layout   GameLayout
	shutdown GameShutdown
	logger   logger.Logger

	time *gametime
}

func (s *shell) Shutdown(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		return s.shutdown.Shutdown(ctx)
	}
}

func (s *shell) Update() error {
	select {
	case <-s.ctx.Done():
		s.logger.Info("context done, terminating game.")
		return errors.Join(s.ctx.Err(), ebiten.Termination)
	default:
		s.time.tick()
		if err := s.update.Update(s.ctx, s.time); err != nil {
			switch {
			case errors.Is(err, context.Canceled), errors.Is(err, context.DeadlineExceeded):
				s.logger.Info("context canceled or deadline exceeded, terminating game.")
				return errors.Join(err, ebiten.Termination)
			case errors.Is(err, ebiten.Termination):
				s.logger.Info("termination signal received, terminating game.")
				return errors.Join(err, ebiten.Termination)
			default:
				s.logger.Error("error during update: " + err.Error())
				return errors.Join(err, ebiten.Termination)
			}
		}
		return nil
	}
}

func (s *shell) Draw(target *ebiten.Image) {
	select {
	case <-s.ctx.Done():
		return
	default:
		if err := s.draw.Draw(target); err != nil {
			s.logger.Error("error during draw: " + err.Error())
		}
	}
}

func (s *shell) Layout(outsideWidth, outsideHeight int) (int, int) {
	w, h, err := s.layout.Layout(outsideWidth, outsideHeight)
	if err != nil {
		s.logger.Error("error during layout: " + err.Error())
		return outsideWidth, outsideHeight
	}
	return w, h
}
