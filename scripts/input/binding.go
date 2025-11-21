package input

type Action string

// Binding represents an input binding consisting of one or more listeners.
//
// Implement custom bindings for different input behaviours such as delayed activation,
// toggles, or combinations of inputs.
type Binding interface {
	Listeners() []Listener // Get all listeners associated with the binding
	Update(dt float64)     // Update the binding state
	IsActive() bool        // Check if the binding is currently active
	JustActive() bool      // Check if the binding was just activated
	JustInactive() bool    // Check if the binding was just deactivated
	Action() Action        // Get the action associated with the binding
	Value() float32        // Get the value of the binding (useful for axis bindings)
}

// ========== Button Bindings ===========

// ButtonBinding represents a binding for a button input (e.g., jump, shoot).
type ButtonBinding struct {
	listeners []Listener
	action    Action
}

func NewButtonBinding(listeners []Listener, action Action) *ButtonBinding {
	return &ButtonBinding{
		listeners: listeners,
		action:    action,
	}
}

func (b *ButtonBinding) Listeners() []Listener {
	return b.listeners
}

func (b *ButtonBinding) Update(dt float64) {
	for _, listener := range b.listeners {
		listener.Update()
	}
}

func (b *ButtonBinding) Action() Action {
	return b.action
}

func (b *ButtonBinding) Value() float32 {
	if b.IsActive() {
		return 1.0
	}
	return 0.0
}

func (b *ButtonBinding) IsActive() bool {
	for _, listener := range b.listeners {
		if listener.IsActive() {
			return true
		}
	}
	return false
}

func (b *ButtonBinding) JustActive() bool {
	for _, listener := range b.listeners {
		if listener.JustActive() {
			return true
		}
	}
	return false
}

func (b *ButtonBinding) JustInactive() bool {
	for _, listener := range b.listeners {
		if listener.JustInactive() {
			return true
		}
	}
	return false
}

// ========== Axis Binding ===========

// AxisBinding represents a binding for an axis input (e.g., horizontal or vertical movement).
type AxisBinding struct {
	positiveListeners []Listener
	negativeListeners []Listener
	action            Action
}

func NewAxisBinding(positiveListeners, negativeListeners []Listener, action Action) *AxisBinding {
	return &AxisBinding{
		positiveListeners: positiveListeners,
		negativeListeners: negativeListeners,
		action:            action,
	}
}
func (b *AxisBinding) Listeners() []Listener {
	return append(b.positiveListeners, b.negativeListeners...)
}

func (b *AxisBinding) Update(dt float64) {
	for _, listener := range b.positiveListeners {
		listener.Update()
	}
	for _, listener := range b.negativeListeners {
		listener.Update()
	}
}

func (b *AxisBinding) IsActive() bool {
	return b.Value() != 0
}

func (b *AxisBinding) JustActive() bool {
	var positiveActive bool
	var negativeActive bool

	for _, listener := range b.positiveListeners {
		if listener.JustActive() {
			positiveActive = true
			break
		}
	}
	for _, listener := range b.negativeListeners {
		if listener.JustActive() {
			negativeActive = true
			break
		}
	}

	return positiveActive || negativeActive
}

func (b *AxisBinding) JustInactive() bool {
	var positiveInactive bool
	var negativeInactive bool

	for _, listener := range b.positiveListeners {
		if listener.JustInactive() {
			positiveInactive = true
			break
		}
	}
	for _, listener := range b.negativeListeners {
		if listener.JustInactive() {
			negativeInactive = true
			break
		}
	}

	return positiveInactive || negativeInactive
}

func (b *AxisBinding) Value() float32 {
	var positive float32
	var negative float32

	for _, listener := range b.positiveListeners {
		if listener.IsActive() {
			positive = listener.Value()
			break
		}
	}
	for _, listener := range b.negativeListeners {
		if listener.IsActive() {
			negative = listener.Value()
			break
		}
	}

	return positive - negative
}

func (b *AxisBinding) Action() Action {
	return b.action
}
