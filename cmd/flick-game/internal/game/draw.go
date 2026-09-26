package game

import (
	"github.com/adm87/flick/pkg/engine"
	"github.com/hajimehoshi/ebiten/v2"
)

// Drawer wraps an engine's game drawer in game-specific screen scaling
type Drawer struct {
	screen *Screen
	drawer engine.GameDraw
}

func NewDrawer(screen *Screen, drawer engine.GameDraw) *Drawer {
	return &Drawer{
		screen: screen,
		drawer: drawer,
	}
}

func (d *Drawer) Draw(target *ebiten.Image) error {
	d.screen.buffer.Clear()

	err := d.drawer.Draw(d.screen.buffer)
	target.DrawImage(d.screen.buffer, d.screen.options)

	return err
}
