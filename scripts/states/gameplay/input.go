package gameplay

import (
	"github.com/adm87/flick/scripts/game"
	"github.com/adm87/flick/scripts/input"
	"github.com/adm87/flick/scripts/systems/player"
	"github.com/hajimehoshi/ebiten/v2"
)

func (s *state) registerInput(g game.Game) {
	// Bind Move Action
	g.Input().Bind(input.NewAxisBinding(
		[]input.Listener{
			input.NewKey(ebiten.KeyD),
			input.NewKey(ebiten.KeyRight),
		},
		[]input.Listener{
			input.NewKey(ebiten.KeyA),
			input.NewKey(ebiten.KeyLeft),
		},
		player.Move,
	))

	// Bind Jump Action
	g.Input().Bind(input.NewSimplePressBinding(
		[]input.Listener{
			input.NewKey(ebiten.KeyW),
			input.NewKey(ebiten.KeyUp),
			input.NewKey(ebiten.KeySpace),
		},
		player.Jump,
	))
}
