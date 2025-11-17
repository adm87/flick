package actors

import "github.com/yohamta/donburi"

var (
	WorldBounds = donburi.NewTag().SetName("WorldBounds")
	Camera      = donburi.NewTag().SetName("Camera")
	Debug       = donburi.NewTag().SetName("Debug")
	Player      = donburi.NewTag().SetName("Player")
	Solid       = donburi.NewTag().SetName("Solid")
)
