package gameplay

import (
	"github.com/adm87/flick/data"
	"github.com/adm87/flick/scripts/assets"
	"github.com/adm87/flick/scripts/collision"
	"github.com/adm87/flick/scripts/game"
)

const (
	GridCellSize = 8
)

var assetBundle = []assets.AssetHandle{
	data.GymCollision,
	data.SampleSheet,
	data.TilemapPacked,
}

type state struct {
	world *collision.World
}

func NewState() game.State {
	return &state{
		world: collision.NewWorld(GridCellSize),
	}
}

func (s *state) Enter(g game.Game) error {
	s.registerSystems(g)

	if err := assets.Load(assetBundle...); err != nil {
		return err
	}

	if err := s.buildWorld(g); err != nil {
		return err
	}

	return nil
}

func (s *state) Exit(g game.Game) error {
	g.ClearSystems()

	// TASK: Unload assets when exiting the state

	return nil
}
