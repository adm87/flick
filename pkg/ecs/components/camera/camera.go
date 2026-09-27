package camera

import (
	"github.com/adm87/flick/pkg/ecs/components/transform"
	"github.com/yohamta/donburi"
)

var MainCamera = donburi.NewTag("MainCamera")

// GetMainCamera retrieves the main camera from the world.
// It returns the entry, the TransformModel of the camera, and a boolean indicating if it was found.
func GetMainCamera(world donburi.World) (*donburi.Entry, *transform.TransformModel, bool) {
	entry, ok := MainCamera.First(world)
	if !ok {
		return nil, nil, false
	}
	cam, ok := transform.GetTransform(entry)
	return entry, cam, ok
}
