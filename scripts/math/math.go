package math

func Abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

func Sign(x float32) float32 {
	if x > 0 {
		return 1
	} else if x < 0 {
		return -1
	}
	return 0
}
