package actors

import "github.com/yohamta/donburi"

var (
	Debug  = donburi.NewTag().SetName("Debug")
	Camera = donburi.NewTag().SetName("Camera")
	Player = donburi.NewTag().SetName("Player")
	Solid  = donburi.NewTag().SetName("Solid")
)
