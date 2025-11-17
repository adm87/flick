package collision

var collisionMatrix [32][32]bool

func init() {
	for i := range collisionMatrix {
		collisionMatrix[DefaultCollisionLayer][i] = true
		collisionMatrix[i][DefaultCollisionLayer] = true
	}
}

// =========== Collision Layer ==========

type CollisionLayer uint8

const (
	MaxCollisionLayers int = 32

	NoCollisionLayer      CollisionLayer = 0
	DefaultCollisionLayer CollisionLayer = iota
)

var nameByLayer = map[CollisionLayer]string{
	DefaultCollisionLayer: "Default",
}

func NewLayer(name string) CollisionLayer {
	if len(nameByLayer) >= MaxCollisionLayers {
		panic("maximum number of collision layers exceeded")
	}

	layer := CollisionLayer(len(nameByLayer))
	nameByLayer[layer] = name

	return layer
}

func (l CollisionLayer) String() string {
	if name, ok := nameByLayer[l]; ok {
		return name
	}
	return "unknown"
}

func (l CollisionLayer) IsValid() bool {
	_, ok := nameByLayer[l]
	return ok
}

func NameByLayer(layer CollisionLayer) (string, bool) {
	name, ok := nameByLayer[layer]
	return name, ok
}

// EnableCollision enables collision detection between the two specified layers.
func EnableCollision(layerA, layerB CollisionLayer) {
	collisionMatrix[layerA][layerB] = true
	collisionMatrix[layerB][layerA] = true
}

// DisableCollision disables collision detection between the two specified layers.
func DisableCollision(layerA, layerB CollisionLayer) {
	collisionMatrix[layerA][layerB] = false
	collisionMatrix[layerB][layerA] = false
}

// ShouldCollide returns true if collision detection is enabled between the two specified layers.
func ShouldCollide(layerA, layerB CollisionLayer) bool {
	return collisionMatrix[layerA][layerB]
}
