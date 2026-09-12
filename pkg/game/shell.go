package game

import (
	"context"
	"errors"

	"github.com/hajimehoshi/ebiten/v2"
)

type shell struct {
	ctx      context.Context
	update   GameUpdate
	draw     GameDraw
	layout   GameLayout
	shutdown GameShutdown
	time     Time
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
		return errors.Join(s.ctx.Err(), ebiten.Termination)
	default:
		err := s.update.Update(s.ctx, s.time)
		if errors.Is(err, context.DeadlineExceeded) {
			return errors.Join(err, ebiten.Termination)
		}
		return err
	}
}

func (s *shell) Draw(target *ebiten.Image) {
	select {
	case <-s.ctx.Done():
		return
	default:
		s.draw.Draw(target)
	}
}

func (s *shell) Layout(outsideWidth, outsideHeight int) (int, int) {
	return s.layout.Layout(outsideWidth, outsideHeight)
}
