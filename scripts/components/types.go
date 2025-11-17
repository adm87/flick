package components

import (
	"github.com/adm87/flick/scripts/components/models"
	"github.com/adm87/flick/scripts/shapes"
	"github.com/yohamta/donburi"
)

var (
	Collider  = donburi.NewComponentType[models.Collider](models.DefaultCollider)
	Debug     = donburi.NewComponentType[models.Debug](models.DefaultDebug)
	Rectangle = donburi.NewComponentType[shapes.Rectangle](*shapes.NewRectangle())
	Movement  = donburi.NewComponentType[models.Movement](models.DefaultMovement)
	Player    = donburi.NewComponentType[models.Player](models.DefaultPlayer)
	Transform = donburi.NewComponentType[models.Transform](models.DefaultTransform)
)
