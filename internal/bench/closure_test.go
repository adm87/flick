package bench_test

import "testing"

const something = "something"

type mockstructclosure struct {
	doSomething func()
	value       string
}

func NewMockStructClosure() *mockstructclosure {
	m := &mockstructclosure{}
	m.doSomething = m.DoSomething
	return m
}

func (m *mockstructclosure) DoSomething() {
	m.value = something
}

func BenchmarkClosure(b *testing.B) {
	b.Run("New closure", func(b *testing.B) {
		m := NewMockStructClosure()
		b.ResetTimer()
		b.ReportAllocs()
		for b.Loop() {
			callSomething(func() {
				m.value = something
			})
		}
	})
	b.Run("Call method", func(b *testing.B) {
		m := NewMockStructClosure()
		b.ResetTimer()
		b.ReportAllocs()
		for b.Loop() {
			callSomething(m.DoSomething)
		}
	})
	b.Run("Call field", func(b *testing.B) {
		m := NewMockStructClosure()
		b.ResetTimer()
		b.ReportAllocs()
		for b.Loop() {
			callSomething(m.doSomething)
		}
	})
}

//go:noinline
func callSomething(something func()) {
	something()
}
