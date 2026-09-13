package engine

import "time"

type Time interface {
	DeltaTime() float64
	FixedDeltaTime() float64
	FixedSteps() int
}

type gametime struct {
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

func newGameTime(fps int) *gametime {
	fps = max(fps, 1)
	return &gametime{
		fixedDeltaTime: time.Second / time.Duration(fps),
	}
}

func (t *gametime) DeltaTime() float64 {
	return t.deltaTime.Seconds()
}

func (t *gametime) FixedDeltaTime() float64 {
	return t.fixedDeltaTime.Seconds()
}

func (t *gametime) FixedSteps() int {
	return t.fixedSteps
}

func (t *gametime) tick() {
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
