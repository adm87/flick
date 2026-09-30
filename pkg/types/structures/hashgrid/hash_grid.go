package hashgrid

import (
	"maps"
	"math"
	"slices"

	"github.com/adm87/flick/pkg/types/geom"
	"github.com/adm87/flick/pkg/types/structures/slotmap"
)

type Cell uint64

func PackCell(x, y int32) Cell {
	return Cell((uint64(uint32(x)) << 32) | uint64(uint32(y)))
}

func UnpackCell(c Cell) (x, y int32) {
	x = int32(uint64(c) >> 32)
	y = int32(uint64(c))
	return
}

type occupied map[uint64]map[Cell]struct{}

type grid map[Cell][]uint64

// Padding represents the padding configuration for the hash grid,
// specifying the number of extra cells to include around the top, bottom, left, and right edges of the grid.
type Padding struct {
	Top    int32
	Bottom int32
	Left   int32
	Right  int32
}

// HashGrid represents a spatial hash grid data structure that organizes items into cells based on their positions.
// It supports efficient insertion, removal, and querying of items based on their spatial location.
type HashGrid[T comparable] struct {
	grid     grid
	occupied occupied
	items    *slotmap.SlotMap[T]
	padding  Padding
	cellSize int
}

func New[T comparable](cellSize int) *HashGrid[T] {
	if cellSize <= 0 {
		panic("cellSize must be greater than 0")
	}
	return &HashGrid[T]{
		grid:     make(grid),
		occupied: make(occupied),
		items:    slotmap.New[T](256),
		padding:  Padding{},
		cellSize: cellSize,
	}
}

func (hg *HashGrid[T]) HasCell(cell Cell) bool {
	_, exists := hg.grid[cell]
	return exists
}

// Cells returns a slice of all occupied cells in the hash grid.
func (hg *HashGrid[T]) Cells() []Cell {
	return slices.Collect(maps.Keys(hg.grid))
}

// Padding returns the current padding configuration of the hash grid.
func (hg *HashGrid[T]) Padding() Padding {
	return hg.padding
}

// WithPadding sets the padding configuration of the hash grid and returns the updated grid.
func (hg *HashGrid[T]) WithPadding(padding Padding) *HashGrid[T] {
	hg.padding = padding
	return hg
}

// CellSize returns the size of each cell in the hash grid.
func (hg *HashGrid[T]) CellSize() int {
	return hg.cellSize
}

// CellCount returns the number of occupied cells in the hash grid.
func (hg *HashGrid[T]) CellCount() int {
	return len(hg.grid)
}

// ItemCount returns the number of items stored in the hash grid.
func (hg *HashGrid[T]) ItemCount() int {
	return hg.items.Len()
}

// Contains checks if the specified item exists in the hash grid.
func (hg *HashGrid[T]) Contains(item T) bool {
	found := false

	hg.items.Range(func(k slotmap.K, value T) bool {
		if item == value {
			found = true
		}
		return !found
	})

	return found
}

// GetCells returns a slice of occupied cells that intersect with the specified rectangular region.
func (hg *HashGrid[T]) GetCells(minX, minY, maxX, maxY float64) []Cell {
	minCellX, minCellY, maxCellX, maxCellY := hg.SnapRegion(minX, minY, maxX, maxY)

	size := (maxCellX - minCellX) * (maxCellY - minCellY)
	cells := make([]Cell, 0, size)

	for y := minCellY; y < maxCellY; y++ {
		for x := minCellX; x < maxCellX; x++ {
			if _, exists := hg.grid[PackCell(x, y)]; exists {
				cells = append(cells, PackCell(x, y))
			}
		}
	}

	return cells
}

// ToGrid returns a slice of cells that intersect with the specified rectangular region.
func (hg *HashGrid[T]) ToGrid(minX, minY, maxX, maxY float64) []Cell {
	minCellX, minCellY, maxCellX, maxCellY := hg.SnapRegion(minX, minY, maxX, maxY)

	size := (maxCellX - minCellX) * (maxCellY - minCellY)
	cells := make([]Cell, 0, size)

	for y := minCellY; y < maxCellY; y++ {
		for x := minCellX; x < maxCellX; x++ {
			cells = append(cells, PackCell(x, y))
		}
	}

	return cells
}

// SnapRegion calculates the grid cell coordinates that encompass the specified rectangular region,
// taking into account the hash grid's cell size and padding.
// It returns the minimum and maximum cell coordinates (minCellX, minCellY, maxCellX, maxCellY).
func (hg *HashGrid[T]) SnapRegion(minX, minY, maxX, maxY float64) (int32, int32, int32, int32) {
	cellSize := float64(hg.cellSize)

	x1 := int32(math.Floor(minX/cellSize)) - hg.padding.Left
	y1 := int32(math.Floor(minY/cellSize)) - hg.padding.Bottom
	x2 := int32(math.Ceil(maxX/cellSize)) + hg.padding.Right
	y2 := int32(math.Ceil(maxY/cellSize)) + hg.padding.Top

	// Ensure min/max ordering for the cell coordinates.
	minCellX, maxCellX := minmax(x1, x2)
	minCellY, maxCellY := minmax(y1, y2)

	return minCellX, minCellY, maxCellX, maxCellY
}

// Insert adds an item to the hash grid, associating it with the specified region and returning its key.
//
// Insert does not check whether the item already exists in the grid. Inserting the same item
// multiple times will result in duplicate entries. If uniqueness is required, the caller can
// use Contains() to check for existence before inserting.
func (hg *HashGrid[T]) Insert(item T, region geom.Rect) uint64 {
	k := hg.items.Insert(item).Pack()
	return hg.insert(k, region)
}

// Remove deletes the specified item from the hash grid, disassociating it from all cells.
func (hg *HashGrid[T]) Remove(key uint64) {
	hg.remove(key)
	hg.items.Delete(slotmap.Unpack(key))
}

// Reinsert updates the position of an existing item in the hash grid by removing it from its current cells and inserting it into the new region.
//
// If the item does not exist in the grid, Reinsert does nothing.
func (hg *HashGrid[T]) Reinsert(key uint64, region geom.Rect) {
	if _, ok := hg.occupied[key]; !ok {
		return
	}
	hg.remove(key)
	hg.insert(key, region)
}

func (hg *HashGrid[T]) insert(k uint64, region geom.Rect) uint64 {
	minX, minY := region.Min()
	maxX, maxY := region.Max()

	cells := hg.ToGrid(minX, minY, maxX, maxY)

	for _, cell := range cells {
		hg.grid[cell] = append(hg.grid[cell], k)

		if _, ok := hg.occupied[k]; !ok {
			hg.occupied[k] = make(map[Cell]struct{})
		}

		hg.occupied[k][cell] = struct{}{}
	}

	return k
}

func (hg *HashGrid[T]) remove(key uint64) {
	cells, ok := hg.occupied[key]

	if !ok {
		return
	}

	for cellID := range cells {
		occupants := hg.grid[cellID]

		for j, k := range occupants {
			if k == key {
				// Note: Unordered removal, copies the last element to the matched index, then truncates the slice.
				occupants[j] = occupants[len(occupants)-1]
				occupants = occupants[:len(occupants)-1]
				break
			}
		}

		// Remove the cell from the grid if it has no more occupants.
		if len(occupants) == 0 {
			delete(hg.grid, cellID)
		} else {
			hg.grid[cellID] = occupants
		}
	}

	delete(hg.occupied, key)
}

func minmax(a, b int32) (int32, int32) {
	if a < b {
		return a, b
	}
	return b, a
}
