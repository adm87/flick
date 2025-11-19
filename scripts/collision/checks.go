package collision

import "github.com/adm87/flick/scripts/shapes"

type Hit struct {
	Delta [2]float32
}

func AABBOverlap(boundsA, boundsB [4]float32) bool {
	return boundsA[0] < boundsB[2] && boundsA[2] > boundsB[0] && boundsA[1] < boundsB[3] && boundsA[3] > boundsB[1]
}

func AABBvsAABB(boundsA, boundsB [4]float32) (Hit, bool) {
	var intersection Hit

	if !AABBOverlap(boundsA, boundsB) {
		return intersection, false
	}

	overlapX1 := boundsB[2] - boundsA[0]
	overlapX2 := boundsA[2] - boundsB[0]
	overlapY1 := boundsB[3] - boundsA[1]
	overlapY2 := boundsA[3] - boundsB[1]

	if overlapX1 < overlapX2 {
		intersection.Delta[0] = overlapX1

	} else {
		intersection.Delta[0] = -overlapX2
	}

	if overlapY1 < overlapY2 {
		intersection.Delta[1] = overlapY1
	} else {
		intersection.Delta[1] = -overlapY2
	}

	return intersection, true
}

func AABBvsPolygon(bounds [4]float32, polygon *shapes.Polygon, polyX, polyY float32) (Hit, bool) {
	var intersection Hit
	var collided bool

	return intersection, collided
}

func absf(x float32) float32 {
	if x < 0 {
		return -x
	}
	return x
}
