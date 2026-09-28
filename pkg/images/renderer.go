package images

import (
	"fmt"
	"image/color"

	"github.com/adm87/flick/pkg/ecs/components/transform"
	"github.com/adm87/flick/pkg/ecs/rendering"
	"github.com/adm87/flick/pkg/types/geom"
	"github.com/hajimehoshi/ebiten/v2"
)

var ErrImageRenderer = fmt.Errorf("image renderer error")

type ImageRenderer struct {
	store       *ImageStore
	drawOptions *ebiten.DrawImageOptions
	missing     *ebiten.Image
}

func NewImageRenderer(store *ImageStore) *ImageRenderer {
	missing := ebiten.NewImage(8, 8)
	missing.Fill(color.RGBA{R: 255, B: 255, A: 255})

	return &ImageRenderer{
		store:       store,
		missing:     missing,
		drawOptions: &ebiten.DrawImageOptions{},
	}
}

func (ir *ImageRenderer) Render(target *ebiten.Image, candidate rendering.RenderingCandidate, _ geom.Rect, viewmatrix ebiten.GeoM) error {
	var renderErr error
	var anchor geom.Vec2

	ir.drawOptions.ColorScale.Reset()

	texture := ir.missing

	if img, ok := GetImage(candidate.Entry); ok {
		ir.drawOptions.ColorScale.ScaleWithColor(img.color)
		anchor = img.anchor

		if tex, err := ir.store.GetFrame(img.handle, img.frame); err != nil {
			renderErr = fmt.Errorf("%w: handle: %v, frame: %v: %w", ErrImageRenderer, img.handle, img.frame, err)
		} else {
			texture = tex
		}
	}

	m := transform.GetMatrix(candidate.Entry)

	ir.drawOptions.GeoM.Reset()
	if anchor != (geom.Vec2{}) {
		ir.drawOptions.GeoM.Translate(
			-anchor.X*float64(texture.Bounds().Dx()),
			-anchor.Y*float64(texture.Bounds().Dy()),
		)
	}
	ir.drawOptions.GeoM.Concat(m)
	ir.drawOptions.GeoM.Concat(viewmatrix)

	target.DrawImage(texture, ir.drawOptions)
	return renderErr
}
