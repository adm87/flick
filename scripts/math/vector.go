package math

import stdmath "math"

func Distance(x1, y1, x2, y2 float32) float32 {
	dx := x2 - x1
	dy := y2 - y1
	return float32(stdmath.Sqrt(float64(dx*dx + dy*dy)))
}
