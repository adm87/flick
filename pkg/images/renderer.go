package images

import (
	"image/color"

	"github.com/adm87/flick/pkg/ecs/components/transform"
	"github.com/adm87/flick/pkg/ecs/rendering"
	"github.com/adm87/flick/pkg/geom"
	"github.com/hajimehoshi/ebiten/v2"
)

type ImageRenderer struct {
	drawOptions *ebiten.DrawImageOptions
	missing     *ebiten.Image
	identity    ebiten.GeoM
}

func NewImageRenderer() *ImageRenderer {
	missing := ebiten.NewImage(8, 8)
	missing.Fill(color.RGBA{R: 255, B: 255, A: 255})
	return &ImageRenderer{
		drawOptions: &ebiten.DrawImageOptions{},
		missing:     missing,
		identity:    ebiten.GeoM{},
	}
}

func (ir *ImageRenderer) Render(target *ebiten.Image, candidate *rendering.RenderingCandidate, viewport geom.Rect, viewmatrix ebiten.GeoM) {
	m := transform.GetMatrix(candidate.Entry)

	ir.drawOptions.GeoM = m
	ir.drawOptions.GeoM.Concat(viewmatrix)

	ir.drawOptions.ColorScale.Reset()

	texture := ir.missing

	if img, ok := GetImage(candidate.Entry); ok {
		ir.drawOptions.ColorScale.ScaleWithColor(img.color)

		// TODO: Pull the image referenced in the img component from a resource store
	}

	target.DrawImage(texture, ir.drawOptions)
}
