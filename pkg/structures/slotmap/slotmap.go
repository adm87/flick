package slotmap

// K is a versioned key that uniquely identifies a slot in the SlotMap.
// It combines a 32-bit index and 32-bit version in a single 64-bit value.
// The version is incremented each time a slot is reused, preventing ABA problems.
type K uint64

// pack combines an index and version into a key K.
func pack(index, version uint32) K {
	return K(uint64(version)<<32 | uint64(uint32(index)))
}

// unpack extracts the index and version from a key K.
func unpack(k K) (index, version uint32) {
	index = uint32(k)
	version = uint32(k >> 32)
	return
}

// checkCapacity validates that the requested capacity is within acceptable bounds.
// It panics if capacity is negative, exceeds uint32 address space, or is too large for the system.
func checkCapacity(newLen int) {
	if newLen < 0 {
		panic("slotmap: capacity is negative")
	}
	if uint64(newLen) > uint64(^uint32(0)) {
		panic("slotmap: capacity exceeds uint32 index space")
	}
	maxInt := int(^uint(0) >> 1)
	if newLen >= maxInt {
		panic("slotmap: capacity too large")
	}
}

// slotData stores the actual value and a link to the next free slot.
type slotData[T any] struct {
	value T
	next  uint32
}

// slot represents a single storage location in the SlotMap.
type slot[T any] struct {
	data    slotData[T]
	version uint32
}

// IsOccupied returns true if this slot currently contains a valid value.
// The occupation state is encoded in the least significant bit of the version.
func (s slot[T]) IsOccupied() bool {
	return s.version&1 == 1
}

// SlotMap is a dense, cache-friendly data structure that stores values with versioned keys.
// It maintains a free-list for O(1) insertions and deletions, and supports efficient iteration
// over occupied slots. The type parameter T specifies the value type to store.
type SlotMap[T any] struct {
	slots []slot[T]
	head  uint32
	len   int
}

// New creates a new SlotMap with an initial capacity.
// The capacity is the maximum number of values the SlotMap can hold without calling Grow.
// Panics if cap is negative or excessively large.
func New[T any](cap int) *SlotMap[T] {
	cap = max(cap, 0)
	checkCapacity(cap + 1)

	slotMap := &SlotMap[T]{
		slots: make([]slot[T], cap+1),
		head:  0,
	}

	if cap == 0 {
		return slotMap
	}

	for i := 1; i < len(slotMap.slots); i++ {
		slotMap.slots[i].data.next = uint32(i + 1)
	}
	slotMap.slots[len(slotMap.slots)-1].data.next = 0
	slotMap.head = 1

	return slotMap
}

// Insert adds a new value to the SlotMap and returns its versioned key.
// If the SlotMap is at capacity, it automatically grows.
func (s *SlotMap[T]) Insert(value T) K {
	if s.head == 0 {
		s.Grow(max(1, len(s.slots)))
	}

	index := s.head

	slot := &s.slots[index]
	s.head = slot.data.next

	slot.version++
	slot.data.value = value

	s.len++
	return pack(index, slot.version)
}

// Get retrieves the value associated with key k.
// It returns the value and true if the key is valid and refers to an occupied slot,
// or a zero value and false if the key is invalid or the slot has been deleted.
func (s *SlotMap[T]) Get(k K) (T, bool) {
	var zero T

	index, version := unpack(k)
	if index == 0 || index >= uint32(len(s.slots)) {
		return zero, false
	}

	slot := &s.slots[index]
	if slot.version != version || !slot.IsOccupied() {
		return zero, false
	}

	return slot.data.value, true
}

// Set updates the value at key k.
// It returns the previous value and true if the update succeeded,
// or a zero value and false if the key is invalid or the slot has been deleted.
func (s *SlotMap[T]) Set(k K, value T) (oldValue T, ok bool) {
	var zero T

	index, version := unpack(k)
	if index == 0 || index >= uint32(len(s.slots)) {
		return zero, false
	}

	slot := &s.slots[index]
	if slot.version != version || !slot.IsOccupied() {
		return zero, false
	}

	oldValue = slot.data.value
	slot.data.value = value
	return oldValue, true
}

// Delete removes the value at key k and returns it.
// It returns the deleted value and true if the deletion succeeded,
// or a zero value and false if the key is invalid or already deleted.
// After deletion, the key becomes permanently invalid due to version increment.
func (s *SlotMap[T]) Delete(k K) (T, bool) {
	var zero T

	index, version := unpack(k)
	if index == 0 || index >= uint32(len(s.slots)) {
		return zero, false
	}

	slot := &s.slots[index]
	if slot.version != version || !slot.IsOccupied() {
		return zero, false
	}

	oldValue := slot.data.value
	slot.data.value = zero

	slot.version++
	slot.data.next = s.head

	s.head = index
	s.len--
	return oldValue, true
}

// Range iterates over all occupied slots in the SlotMap, calling fn for each key-value pair.
// If fn returns false, iteration stops early.
// The order of iteration is not guaranteed to be stable across insertions and deletions.
func (s *SlotMap[T]) Range(fn func(k K, value T) bool) {
	for i := 1; i < len(s.slots); i++ {
		slot := &s.slots[i]
		if slot.IsOccupied() {
			k := pack(uint32(i), slot.version)
			if !fn(k, slot.data.value) {
				return
			}
		}
	}

}

// Grow increases the SlotMap's capacity by n additional slots.
// If n is zero or negative, Grow does nothing.
// Panics if the resulting capacity would exceed limits.
func (s *SlotMap[T]) Grow(n int) {
	if n <= 0 {
		return
	}

	newLen := len(s.slots) + n
	checkCapacity(newLen)

	oldLen := uint32(len(s.slots))
	s.slots = append(s.slots, make([]slot[T], n)...)
	newTail := uint32(len(s.slots) - 1)

	for i := oldLen; i < newTail; i++ {
		s.slots[i].data.next = i + 1
	}

	// Chain the new free-list segment onto the existing head.
	s.slots[newTail].data.next = s.head
	s.head = oldLen
}

// Len returns the number of currently occupied slots in the SlotMap.
func (s *SlotMap[T]) Len() int {
	return s.len
}

// Cap returns the maximum number of values the SlotMap can hold without growing.
func (s *SlotMap[T]) Cap() int {
	return len(s.slots) - 1 // subtract 1 for the sentinel slot at index 0
}
