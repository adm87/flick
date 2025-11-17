package scripts

import (
	"context"

	"github.com/adm87/flick/scripts/game"
	"github.com/adm87/flick/scripts/states/gameplay"
)

// NewGame creates and initializes a new game context.
func NewGame(ctx context.Context, version string) game.Game {
	g := game.NewGame(ctx)

	// For now, we directly set the initial state to gameplay.
	g.OnStart(func(g game.Game) error {
		if err := g.SetState(gameplay.NewState()); err != nil {
			return err
		}
		return nil
	})

	return g
}
