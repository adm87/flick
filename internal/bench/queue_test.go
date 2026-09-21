package bench_test

import "testing"

type item struct {
	Field1 int
	Field2 string
}

type queue struct {
	items []item
}

func BenchmarkSlice(b *testing.B) {
	b.Run("Benchmark queue - A", func(b *testing.B) {
		q := &queue{
			items: make([]item, 0),
		}
		b.ResetTimer()
		b.ReportAllocs()
		for b.Loop() {
			for range 100 {
				q.items = append(q.items, item{Field1: 1, Field2: "test"})
			}
			q.items = make([]item, 0)
		}
	})
	b.Run("Benchmark queue - B", func(b *testing.B) {
		q := &queue{
			items: make([]item, 0, 100),
		}
		b.ResetTimer()
		b.ReportAllocs()
		for b.Loop() {
			for range 100 {
				q.items = append(q.items, item{Field1: 1, Field2: "test"})
			}
			q.items = make([]item, 0, 100)
		}
	})
	b.Run("Benchmark queue - C", func(b *testing.B) {
		q := &queue{
			items: make([]item, 0, 100),
		}
		b.ResetTimer()
		b.ReportAllocs()
		for b.Loop() {
			for range 100 {
				q.items = append(q.items, item{Field1: 1, Field2: "test"})
			}
			q.items = q.items[:0]
		}
	})
	b.Run("Benchmark queue - D", func(b *testing.B) {
		q := &queue{
			items: make([]item, 50),
		}
		i := 0
		b.ResetTimer()
		b.ReportAllocs()
		for b.Loop() {
			for range 100 {
				if i >= len(q.items) {
					q.items = append(q.items, item{})
				}
				q.items[i] = item{Field1: 1, Field2: "test"}
				i++
			}
			i = 0
		}
	})
}
