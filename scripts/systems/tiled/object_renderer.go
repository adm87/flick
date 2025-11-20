package tiled

import (
	"image"

	"github.com/adm87/flick/scripts/assets"
	"github.com/adm87/flick/scripts/components/models"
	"github.com/adm87/flick/scripts/game"
	"github.com/hajimehoshi/ebiten/v2"
)

func RenderObject(ctx game.Context, screen *ebiten.Image, tile *models.Tile, view ebiten.GeoM, matrix ebiten.GeoM) error {
	source := assets.MustGet[*ebiten.Image](tile.Source())

	x, y := tile.Position()
	w, h := tile.Size()
	ax, ay := tile.Offset()

	matrix.Translate(float64(ax), float64(ay))
	matrix.Concat(view)

	screen.DrawImage(source.SubImage(image.Rect(x, y, x+w, y+h)).(*ebiten.Image), &ebiten.DrawImageOptions{
		GeoM: matrix,
	})
	return nil
}
