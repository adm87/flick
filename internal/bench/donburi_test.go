package bench_test

import (
	"testing"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

type system[T any] struct {
	Component *donburi.ComponentType[T]
	Filter    filter.LayoutFilter
}

func newSystem[T any](component *donburi.ComponentType[T]) *system[T] {
	return &system[T]{
		Component: component,
		Filter:    filter.Contains(component),
	}
}

func (s *system[T]) run(e *donburi.Entry) {
	_ = s.Component.Get(e)
}

var globalTestComponent *donburi.ComponentType[struct{}]

func staticRun(e *donburi.Entry) {
	_ = globalTestComponent.Get(e)
}

func BenchmarkDonburi_GetComponent(b *testing.B) {
	b.Run("Single Entity Single Component Lookup", func(b *testing.B) {
		b.ReportAllocs()

		var testComponent = donburi.NewComponentType[struct{}](struct{}{})
		var world = donburi.NewWorld()

		entry := world.Entry(world.Create(testComponent))
		for b.Loop() {
			_ = testComponent.Get(entry)
		}
	})

	b.Run("Single Entity Multiple Component Lookup", func(b *testing.B) {
		b.ReportAllocs()

		var testComponent1 = donburi.NewComponentType[struct{}](struct{}{})
		var testComponent2 = donburi.NewComponentType[struct{}](struct{}{})
		var world = donburi.NewWorld()

		entry := world.Entry(world.Create(testComponent1, testComponent2))
		for b.Loop() {
			_ = testComponent1.Get(entry)
			_ = testComponent2.Get(entry)
		}
	})

	b.Run("Multiple Entity Single Component Lookup (100)", func(b *testing.B) {
		b.ReportAllocs()

		var testComponent = donburi.NewComponentType[struct{}](struct{}{})
		var world = donburi.NewWorld()

		entries := make([]*donburi.Entry, 100)
		for i := range 100 {
			entries[i] = world.Entry(world.Create(testComponent))
		}

		for b.Loop() {
			for i := range 100 {
				_ = testComponent.Get(entries[i])
			}
		}
	})

	b.Run("Multiple Entity Multiple Component Lookup (100)", func(b *testing.B) {
		b.ReportAllocs()

		var testComponent1 = donburi.NewComponentType[struct{}](struct{}{})
		var testComponent2 = donburi.NewComponentType[struct{}](struct{}{})
		var world = donburi.NewWorld()

		entries := make([]*donburi.Entry, 100)
		for i := range 100 {
			entries[i] = world.Entry(world.Create(testComponent1, testComponent2))
		}

		for b.Loop() {
			for i := range 100 {
				_ = testComponent1.Get(entries[i])
				_ = testComponent2.Get(entries[i])
			}
		}
	})

	b.Run("Query Single Component (closure)", func(b *testing.B) {
		b.ReportAllocs()

		var testComponent = donburi.NewComponentType[struct{}](struct{}{})
		var world = donburi.NewWorld()

		for range 100 {
			world.Create(testComponent)
		}

		query := donburi.NewQuery(filter.Contains(testComponent))
		for b.Loop() {
			query.Each(world, func(e *donburi.Entry) {
				_ = testComponent.Get(e)
			})
		}
	})

	b.Run("Query Single Component (system)", func(b *testing.B) {
		b.ReportAllocs()

		var testComponent = donburi.NewComponentType[struct{}](struct{}{})
		var world = donburi.NewWorld()

		for range 100 {
			world.Create(testComponent)
		}

		sys := newSystem(testComponent)
		quy := donburi.NewQuery(sys.Filter)

		for b.Loop() {
			quy.Each(world, sys.run)
		}
	})

	b.Run("Query Single Component (static func)", func(b *testing.B) {
		b.ReportAllocs()

		globalTestComponent = donburi.NewComponentType[struct{}](struct{}{})
		var world = donburi.NewWorld()

		for range 100 {
			world.Create(globalTestComponent)
		}

		query := donburi.NewQuery(filter.Contains(globalTestComponent))
		for b.Loop() {
			query.Each(world, staticRun)
		}
	})
}
