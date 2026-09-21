package game

import (
	"github.com/adm87/flick/pkg/ecs/components/transform"
	"github.com/adm87/flick/pkg/geom"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
)

type View struct {
	world donburi.World
	entry *donburi.Entry
}

func NewView(world donburi.World) *View {
	return &View{
		world: world,
	}
}

func (v *View) SetCamera(entry *donburi.Entry) {
	v.entry = entry
}

// GetView retrieves the viewport and view matrix for the current camera.
// It returns zero values if no camera is set or if the camera does not have a valid transform.
func (v *View) GetView() (viewport geom.Rect, viewmatrix ebiten.GeoM) {
	if v.entry == nil {
		return
	}

	bounds, ok := transform.GetBounds(v.entry)
	if !ok {
		return
	}
	matrix := transform.GetMatrix(v.entry)

	x0, y0 := bounds.Rect.X, bounds.Rect.Y
	x1, y1 := x0+bounds.Rect.Width, y0+bounds.Rect.Height

	ax, ay := matrix.Apply(x0, y0)
	bx, by := matrix.Apply(x1, y0)
	cx, cy := matrix.Apply(x0, y1)
	dx, dy := matrix.Apply(x1, y1)

	minX, maxX := min(ax, bx, cx, dx), max(ax, bx, cx, dx)
	minY, maxY := min(ay, by, cy, dy), max(ay, by, cy, dy)
	viewport = geom.Rect{X: minX, Y: minY, Width: maxX - minX, Height: maxY - minY}

	viewmatrix = matrix
	viewmatrix.Invert()
	return
}

// WorldToScreen converts world coordinates to screen coordinates using the current camera.
// It returns zero values if no camera is set or if the camera does not have a valid transform.
func (v *View) WorldToScreen(x, y float64) (cx, cy float64) {
	if v.entry == nil {
		return
	}
	_, viewmatrix := v.GetView()
	viewmatrix.Invert()
	return viewmatrix.Apply(x, y)
}

// ScreenToWorld converts screen coordinates to world coordinates using the current camera.
// It returns zero values if no camera is set or if the camera does not have a valid transform.
func (v *View) ScreenToWorld(x, y float64) (wx, wy float64) {
	if v.entry == nil {
		return
	}
	_, viewmatrix := v.GetView()
	return viewmatrix.Apply(x, y)
}
