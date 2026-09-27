package game

import (
	"context"
	"math"

	"github.com/adm87/flick/pkg/ecs/components/camera"
	"github.com/adm87/flick/pkg/engine"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/yohamta/donburi"
)

// This is temporary, will be replace with ECS frame scheduling
type Updater struct {
	world donburi.World
}

func NewUpdater(world donburi.World) *Updater {
	return &Updater{
		world: world,
	}
}

func (u *Updater) Update(ctx context.Context, t engine.Time) error {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return ebiten.Termination
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyF) {
		ebiten.SetFullscreen(!ebiten.IsFullscreen())
	}

	_, tr, exists := camera.GetMainCamera(u.world)
	if exists {
		sx, sy := tr.Scale()
		if ebiten.IsKeyPressed(ebiten.KeyUp) {
			sx *= 0.95
			sy *= 0.95
		}
		if ebiten.IsKeyPressed(ebiten.KeyDown) {
			sx *= 1.05
			sy *= 1.05
		}
		tr.SetScale(sx, sy)

		rot := tr.Rotation()
		if ebiten.IsKeyPressed(ebiten.KeyQ) {
			rot -= 0.05
		}
		if ebiten.IsKeyPressed(ebiten.KeyE) {
			rot += 0.05
		}
		tr.SetRotation(rot)

		x, y := tr.Position()

		var dx, dy float64
		if ebiten.IsKeyPressed(ebiten.KeyA) {
			dx -= 5
		}
		if ebiten.IsKeyPressed(ebiten.KeyD) {
			dx += 5
		}
		if ebiten.IsKeyPressed(ebiten.KeyW) {
			dy -= 5
		}
		if ebiten.IsKeyPressed(ebiten.KeyS) {
			dy += 5
		}

		if dx != 0 || dy != 0 {
			cos, sin := math.Cos(rot), math.Sin(rot)
			rdx := dx*cos - dy*sin
			rdy := dx*sin + dy*cos

			x += rdx * sx
			y += rdy * sy
		}

		tr.SetPosition(x, y)
	}
	return nil
}
