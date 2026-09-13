package images

import (
	"image/color"

	"github.com/yohamta/donburi"
)

type ImageData struct {
	Color color.RGBA
	Ref   uint64
	Frame int
}

// ImageComponent represents the image properties of an entity.
var ImageComponent = donburi.NewComponentType[ImageData](ImageData{
	Color: color.RGBA{R: 255, G: 255, B: 255, A: 255},
})

func GetComponent(entry *donburi.Entry) *ImageData {
	return ImageComponent.Get(entry)
}

func GetColor(entry *donburi.Entry) color.RGBA {
	if image := GetComponent(entry); image != nil {
		return image.Color
	}
	return color.RGBA{}
}

func SetColor(entry *donburi.Entry, color color.RGBA) {
	if image := GetComponent(entry); image != nil {
		image.Color = color
	}
}

func GetRef(entry *donburi.Entry) uint64 {
	if image := GetComponent(entry); image != nil {
		return image.Ref
	}
	return 0
}

func SetRef(entry *donburi.Entry, ref uint64) {
	if image := GetComponent(entry); image != nil {
		image.Ref = ref
	}
}

func GetFrame(entry *donburi.Entry) int {
	if image := GetComponent(entry); image != nil {
		return image.Frame
	}
	return 0
}

func SetFrame(entry *donburi.Entry, frame int) {
	if image := GetComponent(entry); image != nil {
		image.Frame = frame
	}
}
