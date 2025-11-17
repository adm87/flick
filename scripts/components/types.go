package components

import (
	"github.com/adm87/flick/scripts/components/models"
	"github.com/yohamta/donburi"
)

var (
	Collider  = donburi.NewComponentType[models.Collider](models.DefaultCollider)
	Debug     = donburi.NewComponentType[models.Debug](models.DefaultDebug)
	Transform = donburi.NewComponentType[models.Transform](models.DefaultTransform)
)
