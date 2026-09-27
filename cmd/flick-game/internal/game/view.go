package game

import (
	"github.com/adm87/flick/pkg/ecs/components/transform"
	"github.com/adm87/flick/pkg/geom"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
)

// View represents a camera view in the game world. It manages the camera's position and transformation,
// and provides methods to convert between world and screen coordinates.
type View struct {
	world  donburi.World
	entry  *donburi.Entry
	screen *Screen
}

func NewView(world donburi.World, screen *Screen) *View {
	return &View{
		world:  world,
		screen: screen,
	}
}

func (v *View) SetCamera(entry *donburi.Entry) {
	v.entry = entry
}

// GetView retrieves the viewport and view matrix for the current camera.
// It returns zero values if no camera is set or if the camera does not have a valid transform.
func (v *View) GetView() (viewport geom.Rect, viewmatrix ebiten.GeoM) {
	if v.entry == nil || !v.entry.Valid() {
		return
	}

	safeArea := v.screen.SafeArea()

	tr, _ := transform.GetTransform(v.entry)
	tr.SetOrigin(safeArea.Center())

	matrix := transform.GetTransformMatrix(v.entry, tr)

	minX, minY := safeArea.Min()
	maxX, maxY := safeArea.Max()

	ax, ay := matrix.Apply(minX, minY)
	bx, by := matrix.Apply(maxX, minY)
	cx, cy := matrix.Apply(minX, maxY)
	dx, dy := matrix.Apply(maxX, maxY)

	vMinX, vMaxX := min(ax, bx, cx, dx), max(ax, bx, cx, dx)
	vMinY, vMaxY := min(ay, by, cy, dy), max(ay, by, cy, dy)

	viewport = geom.Rect{
		X:      vMinX,
		Y:      vMinY,
		Width:  vMaxX - vMinX,
		Height: vMaxY - vMinY,
	}

	viewmatrix = matrix
	viewmatrix.Invert()
	return
}

// WorldToScreen converts world coordinates to screen coordinates using the current camera.
// It returns zero values if no camera is set or if the camera does not have a valid transform.
func (v *View) WorldToScreen(x, y float64) (cx, cy float64) {
	_, matrix := v.GetView()
	return matrix.Apply(x, y)
}

// ScreenToWorld converts screen coordinates to world coordinates using the current camera.
// It returns zero values if no camera is set or if the camera does not have a valid transform.
func (v *View) ScreenToWorld(x, y float64) (wx, wy float64) {
	if v.entry == nil || !v.entry.Valid() {
		return
	}
	matrix := transform.GetMatrix(v.entry)
	return matrix.Apply(x, y)
}
