package models

var DefaultPlayer = Player{
	gravity:  9.81,
	onGround: false,
}

type Player struct {
	gravity  float32
	onGround bool
	onSlope  bool
}

func (p *Player) Gravity() float32 {
	return p.gravity
}

func (p *Player) SetGravity(g float32) *Player {
	p.gravity = g
	return p
}

func (p *Player) OnGround() bool {
	return p.onGround
}

func (p *Player) OnSlope() bool {
	return p.onSlope
}

func (p *Player) SetOnGround(og bool) *Player {
	p.onGround = og
	return p
}

func (p *Player) SetOnSlope(os bool) *Player {
	p.onSlope = os
	return p
}
