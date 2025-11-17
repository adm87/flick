package collision

type CollisionType uint8

const (
	NoCollisionType CollisionType = iota
	StaticCollisionType
	DynamicCollisionType
	TriggerCollisionType
)
