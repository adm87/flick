package aseprite

import "github.com/yohamta/donburi"

type AnimatorModel struct {
	library string
	clip    string
	speed   float64
	loop    bool
}

var AnimatorComponent = donburi.NewComponentType[AnimatorModel](AnimatorModel{
	speed: 1.0,
})

func GetAnimator(entry *donburi.Entry) (*AnimatorModel, bool) {
	if entry.HasComponent(AnimatorComponent) {
		return AnimatorComponent.Get(entry), true
	}
	return nil, false
}

func (a *AnimatorModel) Library() string {
	return a.library
}

func (a *AnimatorModel) SetLibrary(library string) {
	a.library = library
}

func (a *AnimatorModel) Clip() string {
	return a.clip
}

func (a *AnimatorModel) SetClip(clip string) {
	a.clip = clip
}

func (a *AnimatorModel) Speed() float64 {
	return a.speed
}

func (a *AnimatorModel) SetSpeed(speed float64) {
	a.speed = speed
}

func (a *AnimatorModel) Loop() bool {
	return a.loop
}

func (a *AnimatorModel) SetLoop(loop bool) {
	a.loop = loop
}
