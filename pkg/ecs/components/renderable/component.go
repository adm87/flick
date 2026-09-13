package renderable

import "github.com/yohamta/donburi"

type RenderableData struct {
	Type    uint64
	Layer   int
	ZIndex  int
	Visible bool
}

// RenderableComponent represents the renderable properties of an entity.
var RenderableComponent = donburi.NewComponentType[RenderableData](RenderableData{
	Visible: true,
})

func GetComponent(entry *donburi.Entry) *RenderableData {
	return RenderableComponent.Get(entry)
}

func GetType(entry *donburi.Entry) uint64 {
	if renderable := GetComponent(entry); renderable != nil {
		return renderable.Type
	}
	return 0
}

func SetType(entry *donburi.Entry, renderableType uint64) {
	if renderable := GetComponent(entry); renderable != nil {
		renderable.Type = renderableType
	}
}

func GetLayer(entry *donburi.Entry) int {
	if renderable := GetComponent(entry); renderable != nil {
		return renderable.Layer
	}
	return 0
}

func SetLayer(entry *donburi.Entry, layer int) {
	if renderable := GetComponent(entry); renderable != nil {
		renderable.Layer = layer
	}
}

func GetZIndex(entry *donburi.Entry) int {
	if renderable := GetComponent(entry); renderable != nil {
		return renderable.ZIndex
	}
	return 0
}

func SetZIndex(entry *donburi.Entry, zIndex int) {
	if renderable := GetComponent(entry); renderable != nil {
		renderable.ZIndex = zIndex
	}
}

func IsVisible(entry *donburi.Entry) bool {
	if renderable := GetComponent(entry); renderable != nil {
		return renderable.Visible
	}
	return false
}

func SetVisible(entry *donburi.Entry, visible bool) {
	if renderable := GetComponent(entry); renderable != nil {
		renderable.Visible = visible
	}
}
