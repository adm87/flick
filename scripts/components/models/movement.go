package models

var DefaultMovement = Movement{
	velocity: [2]float32{0, 0},
}

type Movement struct {
	velocity [2]float32
}

func (m *Movement) Velocity() (float32, float32) {
	return m.velocity[0], m.velocity[1]
}

func (m *Movement) SetVelocity(vx, vy float32) *Movement {
	m.velocity[0] = vx
	m.velocity[1] = vy
	return m
}
