package tilemap

import (
	"image"
	"math"

	"github.com/adm87/flick/scripts/assets"
	"github.com/adm87/flick/scripts/game"
	"github.com/adm87/tiled"
	"github.com/adm87/tiled/tilemap"
	"github.com/hajimehoshi/ebiten/v2"
)

var op = &ebiten.DrawImageOptions{}

func RenderTilemap(ctx game.Context, screen *ebiten.Image, tmap *tilemap.Map, view ebiten.GeoM, viewport [4]float32) error {
	tmap.BufferFrame()

	itr := tmap.Itr()
	for tiles := itr.Next(); tiles != nil; tiles = itr.Next() {
		for _, tile := range tiles {
			drawTile(ctx, screen, tmap, &tile, view)
		}
	}

	return nil
}

func drawTile(ctx game.Context, screen *ebiten.Image, tmap *tilemap.Map, tile *tilemap.Data, view ebiten.GeoM) {
	tileset, err := tmap.GetTileset(tile.TsIdx)
	if err != nil {
		ctx.Log().Error(err.Error())
		return
	}

	tsx, err := assets.Get[*tiled.Tsx](assets.AssetHandle(tileset.Source))
	if err != nil {
		ctx.Log().Error("missing tsx: " + tileset.Source)
		return
	}

	img, err := assets.Get[*ebiten.Image](assets.AssetHandle(tsx.Image.Source))
	if err != nil {
		ctx.Log().Error("missing image: " + tsx.Image.Source)
		return
	}

	srcX := (int32(tile.TileID) % tsx.Columns) * tsx.TileWidth
	srcY := (int32(tile.TileID) / tsx.Columns) * tsx.TileHeight
	srcRect := image.Rect(int(srcX), int(srcY), int(srcX+tsx.TileWidth), int(srcY+tsx.TileHeight))

	distX := float64(tile.X) + float64(tsx.TileOffset.X)
	distY := float64(tile.Y) + float64(tsx.TileOffset.Y)
	distY -= float64(tsx.TileHeight) - float64(tmap.Tmx.TileHeight) // Align to bottom of tile

	op.GeoM.Reset()

	if tile.FlipFlag.Diagonal() {
		op.GeoM.Rotate(math.Pi * 0.5)
		op.GeoM.Scale(-1, 1)
		op.GeoM.Translate(float64(tsx.TileHeight-tsx.TileWidth), 0)
	}

	if tile.FlipFlag.Horizontal() {
		op.GeoM.Scale(-1, 1)
		op.GeoM.Translate(float64(tsx.TileWidth), 0)
	}

	if tile.FlipFlag.Vertical() {
		op.GeoM.Scale(1, -1)
		op.GeoM.Translate(0, float64(tsx.TileHeight))
	}

	op.GeoM.Translate(distX, distY)
	op.GeoM.Concat(view)

	screen.DrawImage(img.SubImage(srcRect).(*ebiten.Image), op)
}
