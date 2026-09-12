package slotmap_test

import (
	"testing"

	"github.com/adm87/flick/pkg/structures/slotmap"
)

func BenchmarkSlotMap(b *testing.B) {
	b.Run("Insert", func(b *testing.B) {
		b.ReportAllocs()
		sm := slotmap.New[int](1)
		for i := range b.N {
			_ = sm.Insert(i)
		}
	})

	b.Run("Get", func(b *testing.B) {
		b.ReportAllocs()
		sm := slotmap.New[int](1)
		keys := make([]slotmap.K, b.N)
		for i := 0; i < b.N; i++ {
			keys[i] = sm.Insert(i)
		}
		b.ResetTimer()
		for i := range b.N {
			_, _ = sm.Get(keys[i])
		}
	})

	b.Run("Delete", func(b *testing.B) {
		b.ReportAllocs()
		sm := slotmap.New[int](1)
		keys := make([]slotmap.K, b.N)
		for i := 0; i < b.N; i++ {
			keys[i] = sm.Insert(i)
		}
		b.ResetTimer()
		for i := range b.N {
			_, _ = sm.Delete(keys[i])
		}
	})

	b.Run("RangeFull", func(b *testing.B) {
		b.ReportAllocs()
		const size = 1 << 12
		sm, _ := benchmarkFilledMap(size)
		b.ResetTimer()

		for range b.N {
			count := 0
			sm.Range(func(_ slotmap.K, _ int) bool {
				count++
				return true
			})
			if count != size {
				b.Fatalf("expected %d items, got %d", size, count)
			}
		}
	})

	b.Run("RangeEarlyExit", func(b *testing.B) {
		b.ReportAllocs()
		const size = 1 << 12
		sm, _ := benchmarkFilledMap(size)
		b.ResetTimer()

		for range b.N {
			count := 0
			sm.Range(func(_ slotmap.K, _ int) bool {
				count++
				return count < 16
			})
			if count != 16 {
				b.Fatalf("expected early exit at 16, got %d", count)
			}
		}
	})
}

func benchmarkFilledMap(size int) (*slotmap.SlotMap[int], []slotmap.K) {
	sm := slotmap.New[int](1)
	keys := make([]slotmap.K, size)
	for i := range size {
		keys[i] = sm.Insert(i)
	}
	return sm, keys
}
