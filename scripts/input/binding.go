package input

type Action string

// Binding represents an input binding consisting of one or more listeners.
//
// Implement custom bindings for different input behaviours such as delayed activation,
// toggles, or combinations of inputs.
type Binding interface {
	Listeners() []Listener
	Update(dt float64)
	IsActive() bool
	Action() Action
	Value() float32
}

// ========== Simple Press Binding ===========

// SimplePressBinding activates when any of its listeners detect an action (e.g., a key press).
type SimplePressBinding struct {
	listeners []Listener
	action    Action
}

func NewSimplePressBinding(listeners []Listener, action Action) *SimplePressBinding {
	return &SimplePressBinding{
		listeners: listeners,
		action:    action,
	}
}

func (b *SimplePressBinding) Listeners() []Listener {
	return b.listeners
}

func (b *SimplePressBinding) Update(dt float64) {
	for _, listener := range b.listeners {
		listener.Update()
	}
}

func (b *SimplePressBinding) IsActive() bool {
	for _, listener := range b.listeners {
		if listener.IsActive() {
			return true
		}
	}
	return false
}

func (b *SimplePressBinding) Action() Action {
	return b.action
}

func (b *SimplePressBinding) Value() float32 {
	if b.IsActive() {
		return 1.0
	}
	return 0.0
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

func (b *AxisBinding) Value() float32 {
	var positive float32
	var negative float32

	for _, listener := range b.positiveListeners {
		if listener.IsActive() {
			positive += listener.Value()
			break
		}
	}
	for _, listener := range b.negativeListeners {
		if listener.IsActive() {
			negative += listener.Value()
			break
		}
	}

	return positive - negative
}

func (b *AxisBinding) Action() Action {
	return b.action
}
