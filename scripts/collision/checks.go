package collision

type Intersection struct {
	OverlapX float32
	OverlapY float32
}

func AABBOverlap(boundsA, boundsB [4]float32) bool {
	return boundsA[0] < boundsB[2] && boundsA[2] > boundsB[0] && boundsA[1] < boundsB[3] && boundsA[3] > boundsB[1]
}

func AABBvsAABB(boundsA, boundsB [4]float32) (*Intersection, bool) {
	if !AABBOverlap(boundsA, boundsB) {
		return nil, false
	}

	overlapX1 := boundsB[2] - boundsA[0]
	overlapX2 := boundsA[2] - boundsB[0]
	overlapY1 := boundsB[3] - boundsA[1]
	overlapY2 := boundsA[3] - boundsB[1]

	var overlapX, overlapY float32

	if overlapX1 < overlapX2 {
		overlapX = overlapX1
	} else {
		overlapX = -overlapX2
	}

	if overlapY1 < overlapY2 {
		overlapY = overlapY1
	} else {
		overlapY = -overlapY2
	}

	return &Intersection{
		OverlapX: overlapX,
		OverlapY: overlapY,
	}, true
}
