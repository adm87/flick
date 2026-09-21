package renderable

import "github.com/yohamta/donburi"

type RenderableModel struct {
	renderer uint64
	zIndex   int32
	visible  bool
	layer    int8
}

var RenderableComponent = donburi.NewComponentType[RenderableModel](RenderableModel{
	visible: true,
})

func GetRenderable(entry *donburi.Entry) (*RenderableModel, bool) {
	if entry.HasComponent(RenderableComponent) {
		return RenderableComponent.Get(entry), true
	}
	return nil, false
}

func (r *RenderableModel) IsVisible() bool {
	return r.visible
}

func (r *RenderableModel) SetVisible(visible bool) {
	r.visible = visible
}

func (r *RenderableModel) Layer() int8 {
	return r.layer
}

func (r *RenderableModel) Renderer() uint64 {
	return r.renderer
}

func (r *RenderableModel) SetLayer(layer int8) {
	r.layer = layer
}

func (r *RenderableModel) ZIndex() int32 {
	return r.zIndex
}

func (r *RenderableModel) SetZIndex(zIndex int32) {
	r.zIndex = zIndex
}

func (r *RenderableModel) SetRenderer(renderer uint64) {
	r.renderer = renderer
}
