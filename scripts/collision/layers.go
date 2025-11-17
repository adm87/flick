package collision

import "github.com/adm87/flick/scripts/components/models"

var collisionMatrix [32][32]bool

func init() {
	for i := range collisionMatrix {
		collisionMatrix[models.DefaultCollisionLayer][i] = true
		collisionMatrix[i][models.DefaultCollisionLayer] = true
	}
}

// EnableCollision enables collision detection between the two specified layers.
func EnableCollision(layerA, layerB models.CollisionLayer) {
	collisionMatrix[layerA][layerB] = true
	collisionMatrix[layerB][layerA] = true
}

// DisableCollision disables collision detection between the two specified layers.
func DisableCollision(layerA, layerB models.CollisionLayer) {
	collisionMatrix[layerA][layerB] = false
	collisionMatrix[layerB][layerA] = false
}

// ShouldCollide returns true if collision detection is enabled between the two specified layers.
func ShouldCollide(layerA, layerB models.CollisionLayer) bool {
	return collisionMatrix[layerA][layerB]
}
