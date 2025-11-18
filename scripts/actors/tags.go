package actors

import "github.com/yohamta/donburi"

var (
	WorldBounds = donburi.NewTag().SetName("WorldBounds")
	Camera      = donburi.NewTag().SetName("Camera")
	Debug       = donburi.NewTag().SetName("Debug")
	Player      = donburi.NewTag().SetName("Player")
	Slope       = donburi.NewTag().SetName("Slope")
	Solid       = donburi.NewTag().SetName("Solid")
)
