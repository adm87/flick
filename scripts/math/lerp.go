package math

func Lerp(a, b, t float32) float32 {
	return a + t*(b-a)
}

func SmootherStep(a, b, t float32) float32 {
	if t < 0 {
		t = 0
	} else if t > 1 {
		t = 1
	}
	t = t * t * t * (t*(t*6-15) + 10)
	return a + t*(b-a)
}
