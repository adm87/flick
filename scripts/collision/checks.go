package collision

func AABBOverlap(boundsA, boundsB [4]float32) bool {
	return boundsA[0] < boundsB[2] && boundsA[2] > boundsB[0] && boundsA[1] < boundsB[3] && boundsA[3] > boundsB[1]
}
