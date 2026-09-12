package game

import "time"

type Time interface {
	DeltaTime() float64
	FixedDeltaTime() float64
	FixedSteps() int
}

type gtime struct {
	lastTick       time.Time
	deltaTime      time.Duration
	fixedDeltaTime time.Duration
	accumulator    time.Duration
	fixedSteps     int
}

const (
	AccumulatorMax = time.Second
	DeltaTimeMax   = time.Second
)

func newGameTime(fps int) *gtime {
	fps = max(fps, 1)
	return &gtime{
		fixedDeltaTime: time.Second / time.Duration(fps),
	}
}

func (t *gtime) DeltaTime() float64 {
	return t.deltaTime.Seconds()
}

func (t *gtime) FixedDeltaTime() float64 {
	return t.fixedDeltaTime.Seconds()
}

func (t *gtime) FixedSteps() int {
	return t.fixedSteps
}

func (t *gtime) tick() {
	now := time.Now()

	if t.lastTick.IsZero() {
		t.lastTick = now
		return
	}

	t.deltaTime = now.Sub(t.lastTick)
	t.lastTick = now

	if t.deltaTime < 0 {
		t.deltaTime = 0
	}

	if t.deltaTime > DeltaTimeMax {
		t.deltaTime = DeltaTimeMax
	}

	t.accumulator += t.deltaTime
	t.fixedSteps = 0

	if t.accumulator > AccumulatorMax {
		t.accumulator = AccumulatorMax
	}

	for t.accumulator >= t.fixedDeltaTime {
		t.accumulator -= t.fixedDeltaTime
		t.fixedSteps++
	}
}
