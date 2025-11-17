package collision

import (
	"github.com/adm87/utilities/hash"
	"github.com/yohamta/donburi"
)

type World struct {
	grid *hash.Grid[donburi.Entity]
}

func NewWorld(cellSize float32) *World {
	return &World{
		grid: hash.NewGrid[donburi.Entity](cellSize, cellSize),
	}
}

func (w *World) Reinsert(e donburi.Entity, bounds [4]float32) {
	w.Remove(e)
	w.Insert(e, bounds)
}

func (w *World) Insert(e donburi.Entity, region [4]float32) {
	w.grid.Insert(e, region, hash.NoGridPadding)
}

func (w *World) Remove(e donburi.Entity) {
	w.grid.Remove(e)
}

func (w *World) QueryCells(region [4]float32) []uint64 {
	return w.grid.QueryCells(region)
}

func (w *World) Query(region [4]float32) []donburi.Entity {
	return w.grid.Query(region)
}

// Check runs a broad-phase check and return potential entities colliding with the given region.
func (w *World) Check(bounds [4]float32, region [4]float32) []donburi.Entity {
	return w.grid.Query(region)
}
