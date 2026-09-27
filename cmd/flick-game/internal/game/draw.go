package game

import (
	"github.com/adm87/flick/pkg/engine"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/yohamta/donburi"
)

// Drawer wraps an engine's game drawer in game-specific screen scaling
type Drawer struct {
	screen *Screen
	view   *View
	world  donburi.World
	drawer engine.GameDraw
}

func NewDrawer(screen *Screen, view *View, world donburi.World, drawer engine.GameDraw) *Drawer {
	return &Drawer{
		screen: screen,
		view:   view,
		world:  world,
		drawer: drawer,
	}
}

func (d *Drawer) Draw(target *ebiten.Image) error {
	d.screen.buffer.Clear()

	err := d.drawer.Draw(d.screen.buffer)

	x, y := d.view.WorldToScreen(0, 0)
	ebitenutil.DebugPrintAt(d.screen.buffer, "0, 0", int(x), int(y))

	target.DrawImage(d.screen.buffer, d.screen.options)

	return err
}
