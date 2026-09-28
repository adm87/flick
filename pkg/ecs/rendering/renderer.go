package rendering

import (
	"github.com/adm87/flick/pkg/types/geom"
	"github.com/hajimehoshi/ebiten/v2"
)

type RendererK = uint64

type Renderer interface {
	Render(target *ebiten.Image, candidate RenderingCandidate, viewport geom.Rect, viewmatrix ebiten.GeoM) error
}
