package images

import (
	"image/color"

	"github.com/adm87/flick/pkg/geom"
	"github.com/adm87/flick/pkg/resources"
	"github.com/yohamta/donburi"
)

type ImageModel struct {
	handle resources.ResourceHandle
	frame  int
	color  color.RGBA
	anchor geom.Vec2
}

var ImageComponent = donburi.NewComponentType[ImageModel](ImageModel{
	color: color.RGBA{R: 255, G: 255, B: 255, A: 255},
})

func GetImage(entry *donburi.Entry) (*ImageModel, bool) {
	if entry.HasComponent(ImageComponent) {
		return ImageComponent.Get(entry), true
	}
	return nil, false
}

func (i *ImageModel) Color() color.RGBA {
	return i.color
}

func (i *ImageModel) SetColor(color color.RGBA) {
	i.color = color
}

func (i *ImageModel) Frame() int {
	return i.frame
}

func (i *ImageModel) SetFrame(frame int) {
	i.frame = frame
}

func (i *ImageModel) Handle() resources.ResourceHandle {
	return i.handle
}

func (i *ImageModel) SetHandle(handle resources.ResourceHandle) {
	i.handle = handle
}

func (i *ImageModel) Anchor() geom.Vec2 {
	return i.anchor
}

func (i *ImageModel) SetAnchor(anchor geom.Vec2) {
	i.anchor = anchor
}
