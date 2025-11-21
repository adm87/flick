package input

// Mapping manages a collection of input bindings.
type Mapping interface {
	Bind(binding Binding)             // Add or update a binding
	Unbind(action Action)             // Remove a binding by action
	RemoveAllBindings()               // Clear all bindings
	GetBinding(action Action) Binding // Retrieve a binding by action
	Update(dt float64)                // Update all bindings
	IsActive(action Action) bool      // Check if a binding for the action is active
}

type mapping struct {
	bindings map[Action]Binding
}

func NewMapping() Mapping {
	return &mapping{
		bindings: make(map[Action]Binding),
	}
}

func (m *mapping) Bind(binding Binding) {
	m.bindings[binding.Action()] = binding
}

func (m *mapping) Unbind(action Action) {
	delete(m.bindings, action)
}

func (m *mapping) RemoveAllBindings() {
	m.bindings = make(map[Action]Binding)
}

func (m *mapping) GetBinding(action Action) Binding {
	return m.bindings[action]
}

func (m *mapping) Update(dt float64) {
	for _, binding := range m.bindings {
		binding.Update(dt)
	}
}

func (m *mapping) IsActive(action Action) bool {
	if binding := m.GetBinding(action); binding != nil {
		return binding.IsActive()
	}
	return false
}
