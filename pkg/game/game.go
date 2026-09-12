package game

import (
	"context"

	"github.com/hajimehoshi/ebiten/v2"
)

type GameUpdate interface {
	Update(ctx context.Context, t Time) error
}

type GameDraw interface {
	Draw(target *ebiten.Image)
}

type GameLayout interface {
	Layout(outsideWidth, outsideHeight int) (int, int)
}

type GameShutdown interface {
	Shutdown(ctx context.Context) error
}

type noopUpdate struct{}

func (n *noopUpdate) Update(ctx context.Context, t Time) error {
	return nil
}

type noopDraw struct{}

func (n *noopDraw) Draw(target *ebiten.Image) {
}

type noopLayout struct{}

func (n *noopLayout) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

type noopShutdown struct{}

func (n *noopShutdown) Shutdown(ctx context.Context) error {
	return nil
}
