package collision

import (
	"github.com/adm87/flick/scripts/components"
	"github.com/adm87/flick/scripts/components/models"
	"github.com/adm87/flick/scripts/game"
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

// Check runs a broad-phase check for the given bounds against any entities within regional grid.
// It returns a slice of entities that potentially collide with the given bounds based on the collider type and layer.
func (w *World) Check(ctx game.Context, bounds [4]float32, layer models.CollisionLayer, cTypes ...models.ColliderType) []donburi.Entity {
	var results []donburi.Entity

	for _, entity := range w.grid.Query(bounds) {
		other := ctx.ECS().Entry(entity)

		collider := components.Collider.Get(other)
		if !containsColliderType(cTypes, collider.Type()) {
			continue
		}

		if !ShouldCollide(collider.Layer(), layer) {
			continue
		}

		otherBounds := collider.Shape().Bounds(components.Transform.Get(other).Position())
		if !AABBOverlap(bounds, otherBounds) {
			continue
		}

		results = append(results, entity)
	}

	return results
}

func containsColliderType(types []models.ColliderType, cType models.ColliderType) bool {
	for _, t := range types {
		if t == cType {
			return true
		}
	}
	return false
}
