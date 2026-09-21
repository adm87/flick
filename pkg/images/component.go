package images

import (
	"image/color"

	"github.com/yohamta/donburi"
)

type ImageModel struct {
	reference uint64
	frame     int
	color     color.RGBA
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

func (i *ImageModel) Reference() uint64 {
	return i.reference
}

func (i *ImageModel) SetReference(reference uint64) {
	i.reference = reference
}

func (i *ImageModel) Color() color.RGBA {
	return i.color
}

func (i *ImageModel) SetColor(color color.RGBA) {
	i.color = color
}
