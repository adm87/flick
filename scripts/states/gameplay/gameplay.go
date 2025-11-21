package gameplay

import (
	"github.com/adm87/flick/data"
	"github.com/adm87/flick/scripts/assets"
	"github.com/adm87/flick/scripts/collision"
	"github.com/adm87/flick/scripts/game"
	"github.com/adm87/tiled"
	"github.com/adm87/tiled/tilemap"
)

const (
	GridCellSize = 18
)

var assetBundle = []assets.AssetHandle{
	data.TilemapExampleA,
	data.TilemapCharactersPacked,
	data.TilesetCharacters,
	data.TilesetTiles,
	data.TilemapPacked,
}

type state struct {
	world   *collision.World
	tilemap *tilemap.Map
}

func NewState() game.State {
	return &state{
		world:   collision.NewWorld(GridCellSize),
		tilemap: tilemap.NewMap(),
	}
}

func (s *state) Enter(g game.Game) error {
	if err := assets.Load(assetBundle...); err != nil {
		return err
	}

	s.tilemap.SetTmx(assets.MustGet[*tiled.Tmx](data.TilemapExampleA))

	if err := s.buildWorld(g); err != nil {
		return err
	}

	s.registerInput(g)
	s.registerSystems(g)

	return nil
}

func (s *state) Exit(g game.Game) error {
	s.tilemap.Flush()

	g.Input().RemoveAllBindings()

	g.ClearSystems()

	// TASK: Unload assets when exiting the state

	return nil
}
