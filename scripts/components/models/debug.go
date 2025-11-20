package models

var DefaultDebug = Debug{
	drawColliders: true,
	drawPlayer:    true,
	drawTiles:     true,
}

type Debug struct {
	drawColliders     bool
	drawFPS           bool
	drawPlayer        bool
	drawCollisionGrid bool
	drawTiles         bool
}

func (d *Debug) ToggleColliders() {
	d.drawColliders = !d.drawColliders
}

func (d *Debug) ToggleFPS() {
	d.drawFPS = !d.drawFPS
}

func (d *Debug) TogglePlayer() {
	d.drawPlayer = !d.drawPlayer
}

func (d *Debug) ToggleCollisionGrid() {
	d.drawCollisionGrid = !d.drawCollisionGrid
}

func (d *Debug) ToggleTiles() {
	d.drawTiles = !d.drawTiles
}

func (d *Debug) ShowColliders() bool {
	return d.drawColliders
}

func (d *Debug) ShowFPS() bool {
	return d.drawFPS
}

func (d *Debug) ShowPlayer() bool {
	return d.drawPlayer
}

func (d *Debug) ShowStaticGrid() bool {
	return d.drawCollisionGrid
}

func (d *Debug) ShowTiles() bool {
	return d.drawTiles
}
