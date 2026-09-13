package engine

import (
	"context"

	"github.com/hajimehoshi/ebiten/v2"
)

type GameUpdate interface {
	Update(ctx context.Context, t Time) error
}

type GameDraw interface {
	Draw(target *ebiten.Image) error
}

type GameLayout interface {
	Layout(outsideWidth, outsideHeight int) (int, int, error)
}

type GameShutdown interface {
	Shutdown(ctx context.Context) error
}

type noopUpdate struct{}

func (n *noopUpdate) Update(ctx context.Context, t Time) error {
	return nil
}

type noopDraw struct{}

func (n *noopDraw) Draw(target *ebiten.Image) error {
	return nil
}

type noopLayout struct{}

func (n *noopLayout) Layout(outsideWidth, outsideHeight int) (int, int, error) {
	return outsideWidth, outsideHeight, nil
}

type noopShutdown struct{}

func (n *noopShutdown) Shutdown(ctx context.Context) error {
	return nil
}
