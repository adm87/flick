package models

type Player struct {
	gravity  float32
	onGround bool
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

func (p *Player) SetOnGround(og bool) *Player {
	p.onGround = og
	return p
}

var DefaultPlayer = Player{
	gravity:  100.0,
	onGround: false,
}
