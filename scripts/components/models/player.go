package models

var DefaultPlayer = Player{
	coyoteTime: 0.1,
	onGround:   false,
}

type Player struct {
	gravity  float32
	onGround bool

	coyoteTime      float32
	coyoteRemaining float32
}

func (p *Player) OnGround() bool {
	return p.onGround
}

func (p *Player) SetOnGround(og bool) *Player {
	p.onGround = og
	return p
}

func (p *Player) UpdateCoyoteTime(dt float32) {
	if p.onGround {
		p.coyoteRemaining = p.coyoteTime
	} else if p.coyoteRemaining > 0 {
		p.coyoteRemaining -= dt
	}
}

func (p *Player) CanJump() bool {
	return p.onGround || p.coyoteRemaining > 0
}

func (p *Player) SetCoyoteTime(ct float32) *Player {
	p.coyoteTime = ct
	return p
}
