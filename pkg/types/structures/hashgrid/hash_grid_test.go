package hashgrid_test

import (
	"testing"

	"github.com/adm87/flick/pkg/types/geom"
	"github.com/adm87/flick/pkg/types/structures/hashgrid"
	"github.com/go-openapi/testify/v2/require"
)

func TestHashGrid_Initialization(t *testing.T) {
	t.Run("Should initialize a new hash grid and return correct cell size", func(t *testing.T) {
		hg := hashgrid.New[int](10)

		require.NotNil(t, hg)
		require.Equal(t, 0, hg.CellCount())
		require.Equal(t, 0, hg.ItemCount())
		require.Equal(t, 10, hg.CellSize())
	})
	t.Run("Should panic when initialized with non-positive cell size", func(t *testing.T) {
		require.Panics(t, func() {
			_ = hashgrid.New[int](0)
		})
		require.Panics(t, func() {
			_ = hashgrid.New[int](-1)
		})
	})
}

func TestHashGrid_ToGrid(t *testing.T) {
	t.Run("Unpadded", func(t *testing.T) {
		testToGrid(t, hashgrid.Padding{})
	})
	t.Run("Padded", func(t *testing.T) {
		t.Run("Uniform", func(t *testing.T) {
			testToGrid(t, hashgrid.Padding{
				Top:    1,
				Bottom: 1,
				Left:   1,
				Right:  1,
			})
		})
		t.Run("Non-Uniform", func(t *testing.T) {
			testToGrid(t, hashgrid.Padding{
				Top:    2,
				Bottom: 1,
				Left:   3,
				Right:  1,
			})
		})
	})
}

func TestHashGrid_Insert(t *testing.T) {
	t.Run("Should insert an item and increase item count and cell count", func(t *testing.T) {
		hg := hashgrid.New[int](10)
		require.Equal(t, 0, hg.ItemCount())

		hg.Insert(5, geom.NewRect(0, 0, 10, 10))
		require.Equal(t, 1, hg.ItemCount())
		require.Equal(t, 1, hg.CellCount())
	})
	t.Run("Should insert multiple items and increase item count and cell count accordingly", func(t *testing.T) {
		hg := hashgrid.New[int](10)
		require.Equal(t, 0, hg.ItemCount())

		hg.Insert(5, geom.NewRect(0, 0, 10, 10))
		hg.Insert(6, geom.NewRect(10, 10, 10, 10))
		require.Equal(t, 2, hg.ItemCount())
		require.Equal(t, 2, hg.CellCount())
	})
	t.Run("Should insert an item that spans multiple cells and increase item count and cell count accordingly", func(t *testing.T) {
		hg := hashgrid.New[int](10)
		require.Equal(t, 0, hg.ItemCount())

		hg.Insert(5, geom.NewRect(0, 0, 20, 20))
		require.Equal(t, 1, hg.ItemCount())
		require.Equal(t, 4, hg.CellCount())
	})
	t.Run("Should insert multiple items within the same cell and increase item count but not cell count", func(t *testing.T) {
		hg := hashgrid.New[int](10)
		require.Equal(t, 0, hg.ItemCount())

		hg.Insert(5, geom.NewRect(0, 0, 10, 10))
		hg.Insert(6, geom.NewRect(0, 0, 10, 10))
		require.Equal(t, 2, hg.ItemCount())
		require.Equal(t, 1, hg.CellCount())
	})
	t.Run("Should insert duplicate items and not increase cell count", func(t *testing.T) {
		hg := hashgrid.New[int](10)
		require.Equal(t, 0, hg.ItemCount())

		hg.Insert(5, geom.NewRect(0, 0, 10, 10))
		hg.Insert(5, geom.NewRect(0, 0, 10, 10))
		require.Equal(t, 2, hg.ItemCount())
		require.Equal(t, 1, hg.CellCount())
	})
}

func TestHashGrid_Remove(t *testing.T) {
	t.Run("Should remove an item and decrease item count and cell count", func(t *testing.T) {
		hg := hashgrid.New[int](10)
		require.Equal(t, 0, hg.ItemCount())

		k := hg.Insert(5, geom.NewRect(0, 0, 10, 10))
		require.Equal(t, 1, hg.ItemCount())
		require.Equal(t, 1, hg.CellCount())

		hg.Remove(k)
		require.Equal(t, 0, hg.ItemCount())
		require.Equal(t, 0, hg.CellCount())
	})
	t.Run("Should decrease cell count when removing an item that spans multiple cells", func(t *testing.T) {
		hg := hashgrid.New[int](10)
		k := hg.Insert(5, geom.NewRect(0, 0, 20, 20))
		require.Equal(t, 1, hg.ItemCount())
		require.Equal(t, 4, hg.CellCount())

		hg.Remove(k)
		require.Equal(t, 0, hg.ItemCount())
		require.Equal(t, 0, hg.CellCount())
	})
	t.Run("Should only remove unshared cells and preserve shared cells when removing an item", func(t *testing.T) {
		hg := hashgrid.New[int](10)
		hg.Insert(5, geom.NewRect(0, 0, 10, 10))
		k := hg.Insert(6, geom.NewRect(0, 0, 20, 20))
		require.Equal(t, 2, hg.ItemCount())
		require.Equal(t, 4, hg.CellCount())

		hg.Remove(k)
		require.Equal(t, 1, hg.ItemCount())
		require.Equal(t, 1, hg.CellCount())
	})
	t.Run("Should not decrease cell count when removing a non-existent item", func(t *testing.T) {
		hg := hashgrid.New[int](10)
		k := hg.Insert(5, geom.NewRect(0, 0, 10, 10))
		require.Equal(t, 1, hg.ItemCount())
		require.Equal(t, 1, hg.CellCount())

		hg.Remove(k + 1) // Attempt to remove a non-existent item
		require.Equal(t, 1, hg.ItemCount())
		require.Equal(t, 1, hg.CellCount())
		require.True(t, hg.Contains(5)) // Ensure the existing item is still contained
	})
	t.Run("Should remove one associated duplicate item and not affect other items", func(t *testing.T) {
		hg := hashgrid.New[int](10)
		hg.Insert(5, geom.NewRect(0, 0, 10, 10))
		k := hg.Insert(5, geom.NewRect(0, 0, 10, 10))
		require.Equal(t, 2, hg.ItemCount())
		require.Equal(t, 1, hg.CellCount())

		hg.Remove(k)
		require.Equal(t, 1, hg.ItemCount())
		require.Equal(t, 1, hg.CellCount())
		require.True(t, hg.Contains(5))
	})
}

func TestHashGrid_Reinsert(t *testing.T) {
	t.Run("Should reinsert an existing item correctly and not affect the item count or cell count", func(t *testing.T) {
		hg := hashgrid.New[int](10)
		k := hg.Insert(5, geom.NewRect(0, 0, 10, 10))
		require.Equal(t, 1, hg.ItemCount())
		require.Equal(t, 1, hg.CellCount())

		hg.Reinsert(k, geom.NewRect(0, 0, 10, 10))
		require.Equal(t, 1, hg.ItemCount())
		require.Equal(t, 1, hg.CellCount())
		require.True(t, hg.Contains(5))
	})
	t.Run("Should reinsert an existing item and update its position correctly", func(t *testing.T) {
		hg := hashgrid.New[int](10)
		k := hg.Insert(5, geom.NewRect(0, 0, 10, 10))
		require.Equal(t, 1, hg.ItemCount())
		require.Equal(t, 1, hg.CellCount())

		hg.Reinsert(k, geom.NewRect(10, 10, 10, 10))
		require.Equal(t, 1, hg.ItemCount())
		require.Equal(t, 1, hg.CellCount())
		require.True(t, hg.Contains(5))
	})
	t.Run("Should reinsert an existing item to a larger region", func(t *testing.T) {
		hg := hashgrid.New[int](10)
		k := hg.Insert(5, geom.NewRect(0, 0, 10, 10))
		require.Equal(t, 1, hg.ItemCount())
		require.Equal(t, 1, hg.CellCount())

		hg.Reinsert(k, geom.NewRect(10, 10, 20, 20))
		require.Equal(t, 1, hg.ItemCount())
		require.Equal(t, 4, hg.CellCount())
		require.True(t, hg.Contains(5))
	})
	t.Run("Should reinsert an existing item to a smaller region", func(t *testing.T) {
		hg := hashgrid.New[int](10)
		k := hg.Insert(5, geom.NewRect(0, 0, 20, 20))
		require.Equal(t, 1, hg.ItemCount())
		require.Equal(t, 4, hg.CellCount())

		hg.Reinsert(k, geom.NewRect(0, 0, 10, 10))
		require.Equal(t, 1, hg.ItemCount())
		require.Equal(t, 1, hg.CellCount())
		require.True(t, hg.Contains(5))
	})
	t.Run("Should reinsert to a shared cell correctly", func(t *testing.T) {
		hg := hashgrid.New[int](10)
		k1 := hg.Insert(5, geom.NewRect(0, 0, 10, 10))
		_ = hg.Insert(6, geom.NewRect(0, 0, 10, 10))
		require.Equal(t, 2, hg.ItemCount())
		require.Equal(t, 1, hg.CellCount())

		hg.Reinsert(k1, geom.NewRect(0, 0, 10, 10))
		require.Equal(t, 2, hg.ItemCount())
		require.Equal(t, 1, hg.CellCount())
		require.True(t, hg.Contains(5))
		require.True(t, hg.Contains(6))
	})
	t.Run("Should handle reinserting multiple items correctly", func(t *testing.T) {
		hg := hashgrid.New[int](10)
		k1 := hg.Insert(5, geom.NewRect(0, 0, 10, 10))
		k2 := hg.Insert(6, geom.NewRect(0, 0, 10, 10))
		require.Equal(t, 2, hg.ItemCount())
		require.Equal(t, 1, hg.CellCount())

		hg.Reinsert(k1, geom.NewRect(10, 10, 10, 10))
		hg.Reinsert(k2, geom.NewRect(20, 20, 10, 10))
		require.Equal(t, 2, hg.ItemCount())
		require.Equal(t, 2, hg.CellCount())
		require.True(t, hg.Contains(5))
		require.True(t, hg.Contains(6))
	})
	t.Run("Should not reinsert an item that has been removed", func(t *testing.T) {
		hg := hashgrid.New[int](10)
		k := hg.Insert(5, geom.NewRect(0, 0, 10, 10))
		hg.Remove(k)
		hg.Reinsert(k, geom.NewRect(0, 0, 10, 10)) // Attempt to reinsert a removed item
		require.Equal(t, 0, hg.ItemCount())
		require.Equal(t, 0, hg.CellCount())
	})
	t.Run("Should not reinsert a non-existent item", func(t *testing.T) {
		hg := hashgrid.New[int](10)
		hg.Reinsert(999, geom.NewRect(0, 0, 10, 10)) // Attempt to reinsert a non-existent item
		require.Equal(t, 0, hg.ItemCount())
		require.Equal(t, 0, hg.CellCount())
	})
	t.Run("Should reinsert and remove old footprint", func(t *testing.T) {
		hg := hashgrid.New[int](10)
		k := hg.Insert(5, geom.NewRect(0, 0, 20, 20))
		require.Equal(t, 1, hg.ItemCount())
		require.Equal(t, 4, hg.CellCount())

		// Capture the current set of occupied cells before reinserting the item.
		cells := hg.Cells()

		hg.Reinsert(k, geom.NewRect(20, 20, 20, 20))
		require.Equal(t, 1, hg.ItemCount())
		require.Equal(t, 4, hg.CellCount())
		require.True(t, hg.Contains(5))

		// Verify that the old cells are no longer occupied after reinserting the item.
		for _, cell := range cells {
			require.False(t, hg.HasCell(cell))
		}
	})
}

func TestHashGrid_Contains(t *testing.T) {
	t.Run("Should return true for an item that has been inserted", func(t *testing.T) {
		hg := hashgrid.New[int](10)
		hg.Insert(5, geom.NewRect(0, 0, 10, 10))
		require.True(t, hg.Contains(5))
	})
	t.Run("Should return false for an item that has not been inserted", func(t *testing.T) {
		hg := hashgrid.New[int](10)
		require.False(t, hg.Contains(5))
	})
}

func testToGrid(t *testing.T, padding hashgrid.Padding) {
	t.Helper()
	t.Run("Should return the correct number of cells for a large region", func(t *testing.T) {
		const testSize = 10
		baseWidth := 2 * testSize

		paddedWidth := baseWidth + int(padding.Left) + int(padding.Right)
		paddedHeight := baseWidth + int(padding.Top) + int(padding.Bottom)

		hg := hashgrid.New[int](testSize).WithPadding(padding)

		maxCoord := float64(testSize * hg.CellSize())
		cells := hg.ToGrid(-maxCoord, -maxCoord, maxCoord, maxCoord)

		require.NotNil(t, cells)
		require.Equal(t, paddedWidth*paddedHeight, len(cells))

		minCellX := -testSize - padding.Left
		minCellY := -testSize - padding.Bottom
		rowWidth := paddedWidth

		for i := range cells {
			x, y := hashgrid.UnpackCell(cells[i])

			expectedX := minCellX + int32(i%rowWidth)
			expectedY := minCellY + int32(i/rowWidth)

			if x != expectedX || y != expectedY {
				t.Errorf("Cell mismatch at index %d: got (%d, %d), expected (%d, %d)", i, x, y, expectedX, expectedY)
			}
		}
	})
	t.Run("Should return the correct number of cells for small regions", func(t *testing.T) {
		const testSize = 10
		halfSize := float64(testSize) / 2

		hg := hashgrid.New[int](testSize).WithPadding(padding)

		cells := hg.ToGrid(0, 0, halfSize, halfSize)

		expectedWidth := 1 + int(padding.Left) + int(padding.Right)
		expectedHeight := 1 + int(padding.Top) + int(padding.Bottom)
		expectedCount := expectedWidth * expectedHeight

		require.NotNil(t, cells)
		require.Equal(t, expectedCount, len(cells))

		minX := -padding.Left
		minY := -padding.Bottom
		maxX := padding.Right
		maxY := padding.Top

		for y := minY; y <= maxY; y++ {
			for x := minX; x <= maxX; x++ {
				require.Contains(t, cells, hashgrid.PackCell(int32(x), int32(y)))
			}
		}
	})
	t.Run("Should return the correct number of cells for a region that matches exactly one cell", func(t *testing.T) {
		const testSize = 10
		hg := hashgrid.New[int](testSize).WithPadding(padding)

		cells := hg.ToGrid(0, 0, float64(testSize), float64(testSize))

		expectedWidth := 1 + int(padding.Left) + int(padding.Right)
		expectedHeight := 1 + int(padding.Top) + int(padding.Bottom)
		expectedCount := expectedWidth * expectedHeight

		require.NotNil(t, cells)
		require.Equal(t, expectedCount, len(cells))

		minX := -padding.Left
		minY := -padding.Bottom
		maxX := padding.Right
		maxY := padding.Top

		for y := minY; y <= maxY; y++ {
			for x := minX; x <= maxX; x++ {
				require.Contains(t, cells, hashgrid.PackCell(int32(x), int32(y)))
			}
		}
	})
}
