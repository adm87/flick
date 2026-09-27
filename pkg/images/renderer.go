package images

import (
	"fmt"
	"image/color"

	"github.com/adm87/flick/pkg/ecs/components/transform"
	"github.com/adm87/flick/pkg/ecs/rendering"
	"github.com/adm87/flick/pkg/geom"
	"github.com/hajimehoshi/ebiten/v2"
)

var ErrImageRenderer = fmt.Errorf("image renderer error")

type ImageRenderer struct {
	store       *ImageStore
	drawOptions *ebiten.DrawImageOptions
	missing     *ebiten.Image
	identity    ebiten.GeoM
}

func NewImageRenderer(store *ImageStore) *ImageRenderer {
	missing := ebiten.NewImage(8, 8)
	missing.Fill(color.RGBA{R: 255, B: 255, A: 255})

	return &ImageRenderer{
		store:       store,
		missing:     missing,
		drawOptions: &ebiten.DrawImageOptions{},
		identity:    ebiten.GeoM{},
	}
}

func (ir *ImageRenderer) Render(target *ebiten.Image, candidate rendering.RenderingCandidate, viewport geom.Rect, viewmatrix ebiten.GeoM) error {
	var renderErr error

	m := transform.GetMatrix(candidate.Entry)

	ir.drawOptions.GeoM = m
	ir.drawOptions.GeoM.Concat(viewmatrix)

	ir.drawOptions.ColorScale.Reset()

	texture := ir.missing

	if img, ok := GetImage(candidate.Entry); ok {
		ir.drawOptions.ColorScale.ScaleWithColor(img.color)

		if tex, err := ir.store.GetFrame(img.handle, img.frame); err != nil {
			renderErr = fmt.Errorf("%w: handle: %v, frame: %v: %w", ErrImageRenderer, img.handle, img.frame, err)
		} else {
			texture = tex
		}
	}

	target.DrawImage(texture, ir.drawOptions)
	return renderErr
}
