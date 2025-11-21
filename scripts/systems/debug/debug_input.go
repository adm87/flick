package debug

import (
	"github.com/adm87/flick/scripts/game"

	"github.com/adm87/flick/scripts/actors"
	"github.com/adm87/flick/scripts/components"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func PollDebugInput(ctx game.Context) error {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		return ebiten.Termination
	}

	debugEntry := actors.Debug.MustFirst(ctx.ECS())
	debugModel := components.Debug.Get(debugEntry)

	if inpututil.IsKeyJustPressed(ebiten.KeyF9) {
		debugModel.ToggleColliders()
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyF5) {
		debugModel.TogglePlayer()
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyF6) {
		debugModel.ToggleTiles()
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyF7) {
		debugModel.ToggleTilemap()
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyF10) {
		debugModel.ToggleCollisionGrid()
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyF11) {
		ebiten.SetFullscreen(!ebiten.IsFullscreen())
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyF12) {
		debugModel.ToggleFPS()
	}

	return nil
}
