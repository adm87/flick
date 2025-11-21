package input

import "github.com/hajimehoshi/ebiten/v2"

// A listener manages listening for when devices input occurs.
type Listener interface {
	Update()            // Update the listener state
	IsActive() bool     // Check if the listener is currently active
	JustActive() bool   // Check if the listener was just activated
	JustInactive() bool // Check if the listener was just deactivated
	Value() float32     // Get the value of the listener (useful for analog inputs)
}

// ========== Key Listener ===========

// Key represents a keyboard key listener.
type Key struct {
	key       ebiten.Key
	wasActive bool
	active    bool
}

func NewKey(key ebiten.Key) *Key {
	return &Key{
		key: key,
	}
}

func (k *Key) Update() {
	k.wasActive = k.active
	k.active = ebiten.IsKeyPressed(k.key)
	// Note: could use inpututils.IsKeyJustPressed for this,
	// however we'll manage the state ourselves to keep it consistent with other input types.
}

func (k *Key) IsActive() bool {
	return k.active
}

func (k *Key) JustActive() bool {
	return k.active && !k.wasActive
}

func (k *Key) JustInactive() bool {
	return !k.active && k.wasActive
}

func (k *Key) Value() float32 {
	if k.active {
		return 1.0
	}
	return 0.0
}
