package debug

import (
	"fmt"

	"image/color"
	"iter"

	"github.com/adm87/flick/scripts/components"
	"github.com/adm87/flick/scripts/game"
	"github.com/adm87/flick/scripts/shapes"
	"github.com/adm87/utilities/hash"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/yohamta/donburi"
)

func DrawFPS(ctx game.Context, screen *ebiten.Image) error {
	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %0.2f", ebiten.ActualFPS()))
	return nil
}

func DrawRect(ctx game.Context, screen *ebiten.Image, pos [2]float32, size [2]float32, col color.Color) error {
	vector.StrokeRect(screen, pos[0], pos[1], size[0], size[1], 1, col, false)
	return nil
}

func DrawPolygon(ctx game.Context, screen *ebiten.Image, pos [2]float32, vertices [][2]float32, col color.Color) error {
	n := len(vertices)
	if n < 2 {
		return nil
	}
	for i := 0; i < n; i++ {
		x1 := pos[0] + vertices[i][0]
		y1 := pos[1] + vertices[i][1]
		x2 := pos[0] + vertices[(i+1)%n][0]
		y2 := pos[1] + vertices[(i+1)%n][1]
		vector.StrokeLine(screen, x1, y1, x2, y2, 1, col, false)
	}
	return nil
}

func DrawCollisionGrid(ctx game.Context, screen *ebiten.Image, view ebiten.GeoM, cells []uint64, cellSize float32, col color.Color) error {
	for _, cell := range cells {
		cellX, cellY := hash.DecodeGridKey(cell)
		x, y := view.Apply(float64(cellX*int32(cellSize)), float64(cellY*int32(cellSize)))
		if err := DrawRect(ctx, screen, [2]float32{float32(x), float32(y)}, [2]float32{cellSize, cellSize}, col); err != nil {
			return err
		}
	}
	return nil
}

func DrawEntityColliders(ctx game.Context, screen *ebiten.Image, view ebiten.GeoM, itr iter.Seq[*donburi.Entry], col color.Color) error {
	for e := range itr {
		matrix := components.Transform.Get(e).Matrix()
		matrix.Concat(view)

		shape := components.Collider.Get(e).Shape()
		switch shape.(type) {
		case *shapes.Rectangle:
			bounds := shape.Bounds(0, 0)

			x, y := matrix.Apply(0, 0)
			w, h := bounds[2]-bounds[0], bounds[3]-bounds[1]

			if err := DrawRect(ctx, screen, [2]float32{float32(x), float32(y)}, [2]float32{w, h}, col); err != nil {
				return err
			}

		case *shapes.Polygon:
			polygon := shape.(*shapes.Polygon)

			x, y := matrix.Apply(0, 0)
			if err := DrawPolygon(ctx, screen, [2]float32{float32(x), float32(y)}, polygon.Vertices(), col); err != nil {
				return err
			}
		}
	}
	return nil
}
