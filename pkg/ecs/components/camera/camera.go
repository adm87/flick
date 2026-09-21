package camera

import (
	"github.com/adm87/flick/pkg/ecs/components/transform"
	"github.com/yohamta/donburi"
)

type CameraModel struct {
	transform.TransformModel
}

var CameraComponent = donburi.NewComponentType[CameraModel](CameraModel{
	TransformModel: transform.DefaultTransform(),
})

var MainCamera = donburi.NewTag("MainCamera")

// GetMainCamera retrieves the main camera from the world.
// It returns the entry, the CameraModel, and a boolean indicating if it was found.
//
// If no main camera is found, or it doesn't have a CameraComponent, it returns nil, nil, and false.
func GetMainCamera(world donburi.World) (*donburi.Entry, *CameraModel, bool) {
	entry, ok := MainCamera.First(world)
	if !ok {
		return nil, nil, false
	}
	cam, ok := GetCamera(entry)
	return entry, cam, ok
}

func GetCamera(entry *donburi.Entry) (*CameraModel, bool) {
	if !entry.HasComponent(CameraComponent) {
		return nil, false
	}
	return CameraComponent.Get(entry), true
}

func (c *CameraModel) Zoom() float64 {
	sx, sy := c.Scale()
	return (sx + sy) * 0.5
}

func (c *CameraModel) SetZoom(zoom float64) {
	c.SetScale(zoom, zoom)
}
