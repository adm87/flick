package slotmap_test

import (
	"math"
	"testing"

	"github.com/adm87/flick/pkg/structures/slotmap"
	"github.com/go-openapi/testify/v2/require"
)

func TestSlotMap(t *testing.T) {
	t.Run("Should initialize with correct length and capacity", func(t *testing.T) {
		slotMap := slotmap.New[int](10)
		require.NotNil(t, slotMap)
		require.Equal(t, 0, slotMap.Len())
		require.Equal(t, 10, slotMap.Cap())
	})

	t.Run("Should clamp negative seed to zero", func(t *testing.T) {
		slotMap := slotmap.New[int](-5)
		require.NotNil(t, slotMap)
		require.Equal(t, 0, slotMap.Len())
		require.Equal(t, 0, slotMap.Cap())

		k := slotMap.Insert(7)
		v, ok := slotMap.Get(k)
		require.True(t, ok)
		require.Equal(t, 7, v)
	})

	t.Run("Should panic when seed exceeds uint32 index capacity", func(t *testing.T) {
		require.Panics(t, func() {
			_ = slotmap.New[int](math.MaxUint32 + 1)
		})
	})

	t.Run("Should insert and retrieve a value correctly", func(t *testing.T) {
		slotMap := slotmap.New[int](2)

		k := slotMap.Insert(42)
		v, ok := slotMap.Get(k)

		require.True(t, ok)
		require.Equal(t, 42, v)
	})

	t.Run("Should keep free list valid after grow", func(t *testing.T) {
		slotMap := slotmap.New[int](1)

		k := slotMap.Insert(10)
		_, ok := slotMap.Delete(k)
		require.True(t, ok)

		slotMap.Grow(2)

		k2 := slotMap.Insert(20)
		k3 := slotMap.Insert(30)
		k4 := slotMap.Insert(40)

		v2, ok := slotMap.Get(k2)
		require.True(t, ok)
		require.Equal(t, 20, v2)

		v3, ok := slotMap.Get(k3)
		require.True(t, ok)
		require.Equal(t, 30, v3)

		v4, ok := slotMap.Get(k4)
		require.True(t, ok)
		require.Equal(t, 40, v4)

		require.Equal(t, 3, slotMap.Len())
	})

	t.Run("Should set value and return old value", func(t *testing.T) {
		slotMap := slotmap.New[string](1)

		k := slotMap.Insert("old")
		oldValue, ok := slotMap.Set(k, "new")
		require.True(t, ok)
		require.Equal(t, "old", oldValue)

		updated, ok := slotMap.Get(k)
		require.True(t, ok)
		require.Equal(t, "new", updated)
	})

	t.Run("Should delete value and invalidate stale key", func(t *testing.T) {
		slotMap := slotmap.New[int](2)

		k := slotMap.Insert(99)
		require.Equal(t, 1, slotMap.Len())

		deleted, ok := slotMap.Delete(k)
		require.True(t, ok)
		require.Equal(t, 99, deleted)
		require.Equal(t, 0, slotMap.Len())

		_, ok = slotMap.Get(k)
		require.False(t, ok)

		_, ok = slotMap.Set(k, 100)
		require.False(t, ok)

		_, ok = slotMap.Delete(k)
		require.False(t, ok)
	})

	t.Run("Should return false for invalid keys", func(t *testing.T) {
		slotMap := slotmap.New[int](1)

		_, ok := slotMap.Get(slotmap.K(0))
		require.False(t, ok)

		_, ok = slotMap.Set(slotmap.K(0), 1)
		require.False(t, ok)

		_, ok = slotMap.Delete(slotmap.K(0))
		require.False(t, ok)

		invalidOutOfRangeKey := slotmap.K(^uint64(0))
		_, ok = slotMap.Get(invalidOutOfRangeKey)
		require.False(t, ok)
	})

	t.Run("Should range over occupied entries and support early stop", func(t *testing.T) {
		slotMap := slotmap.New[int](4)

		k1 := slotMap.Insert(10)
		slotMap.Insert(20)
		k3 := slotMap.Insert(30)
		_, ok := slotMap.Delete(k1)
		require.True(t, ok)

		seen := map[int]int{}
		slotMap.Range(func(_ slotmap.K, value int) bool {
			seen[value]++
			return true
		})

		require.Equal(t, 0, seen[10])
		require.Equal(t, 1, seen[20])
		require.Equal(t, 1, seen[30])

		visited := 0
		slotMap.Range(func(k slotmap.K, _ int) bool {
			visited++
			return k != k3
		})
		require.True(t, visited >= 1)
		require.True(t, visited <= 2)
	})
}
