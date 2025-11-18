package game

type Time struct {
	deltaTime      float64
	fixedDeltaTime float64
	fixedSteps     int
	maxSteps       int
	accumulator    float64
}

func NewTime(fixedDeltaTime float64, maxSteps int) *Time {
	return &Time{
		fixedDeltaTime: fixedDeltaTime,
		maxSteps:       maxSteps,
	}
}

func (t *Time) DeltaTime() float64      { return t.deltaTime }
func (t *Time) FixedDeltaTime() float64 { return t.fixedDeltaTime }
func (t *Time) FixedSteps() int         { return t.fixedSteps }
func (t *Time) Alpha() float64          { return t.accumulator / t.fixedDeltaTime }

func (t *Time) tick(elapsed float64) {
	t.deltaTime = elapsed
	t.accumulator += elapsed

	maxAccum := t.fixedDeltaTime * float64(t.maxSteps)
	if t.accumulator > maxAccum {
		t.accumulator = maxAccum
	}

	t.fixedSteps = 0
	for t.accumulator >= t.fixedDeltaTime {
		t.fixedSteps++
		t.accumulator -= t.fixedDeltaTime
	}
}
