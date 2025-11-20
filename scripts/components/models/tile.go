package models

import "github.com/adm87/flick/scripts/assets"

var DefaultTile = Tile{
	size: [2]int{1, 1},
	gid:  0,
	src:  "",
}

type Tile struct {
	offset [2]float32
	size   [2]int
	xy     [2]int
	gid    uint32
	src    assets.AssetHandle
}

func (t *Tile) Offset() (float32, float32) {
	return t.offset[0], t.offset[1]
}

func (t *Tile) SetOffset(ax, ay float32) *Tile {
	t.offset[0] = ax
	t.offset[1] = ay
	return t
}

func (t *Tile) Size() (int, int) {
	return t.size[0], t.size[1]
}

func (t *Tile) GID() uint32 {
	return t.gid
}

func (t *Tile) Source() assets.AssetHandle {
	return t.src
}

func (t *Tile) Position() (int, int) {
	return t.xy[0], t.xy[1]
}

func (t *Tile) SetGID(gid uint32) *Tile {
	t.gid = gid
	return t
}

func (t *Tile) SetSource(src assets.AssetHandle) *Tile {
	t.src = src
	return t
}

func (t *Tile) SetSize(width, height int) *Tile {
	t.size[0] = width
	t.size[1] = height
	return t
}

func (t *Tile) SetPosition(x, y int) *Tile {
	t.xy[0] = x
	t.xy[1] = y
	return t
}
