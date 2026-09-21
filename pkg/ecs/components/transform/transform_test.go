package transform_test

import (
	"strconv"
	"testing"

	"github.com/adm87/flick/pkg/ecs/components/transform"
	"github.com/go-openapi/testify/v2/require"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
)

var benchSizes = []int{1, 10, 100, 1000, 10000}

func TestTransform(t *testing.T) {
	var testCases = []struct {
		name string
		Run  func(t *testing.T, world donburi.World)
	}{
		{
			name: "Get Transform",
			Run: func(t *testing.T, world donburi.World) {
				entry := world.Entry(world.Create(
					transform.TransformComponent,
				))
				_, ok := transform.GetTransform(entry)
				require.True(t, ok)
			},
		},
		{
			name: "Get Transform Fails Without Component",
			Run: func(t *testing.T, world donburi.World) {
				entry := world.Entry(world.Create(donburi.NewComponentType[struct{}]()))
				_, ok := transform.GetTransform(entry)
				require.False(t, ok)
			},
		},
		{
			name: "Fresh Transform Is Clean With Defaults",
			Run: func(t *testing.T, world donburi.World) {
				entry := world.Entry(world.Create(transform.TransformComponent))
				tr, _ := transform.GetTransform(entry)
				require.True(t, tr.IsDirty())

				sx, sy := tr.Scale()
				require.Equal(t, 1.0, sx)
				require.Equal(t, 1.0, sy)
			},
		},
		{
			name: "Position Accessors",
			Run: func(t *testing.T, world donburi.World) {
				entry := world.Entry(world.Create(
					transform.TransformComponent,
				))

				tr, ok := transform.GetTransform(entry)
				require.True(t, ok)

				tr.SetPosition(1, 2)
				require.True(t, tr.IsDirty())

				// Re-fetch the transform to simulate a new access after modification
				tr, ok = transform.GetTransform(entry)
				require.True(t, ok)

				x, y := tr.Position()
				require.Equal(t, 1.0, x)
				require.Equal(t, 2.0, y)
			},
		},
		{
			name: "Rotation Accessors",
			Run: func(t *testing.T, world donburi.World) {
				entry := world.Entry(world.Create(
					transform.TransformComponent,
				))

				tr, ok := transform.GetTransform(entry)
				require.True(t, ok)

				tr.SetRotation(45)

				// Re-fetch the transform to simulate a new access after modification
				tr, ok = transform.GetTransform(entry)
				require.True(t, ok)

				require.True(t, tr.IsDirty())

				rotation := tr.Rotation()
				require.Equal(t, 45.0, rotation)
			},
		},
		{
			name: "Scale Accessors",
			Run: func(t *testing.T, world donburi.World) {
				entry := world.Entry(world.Create(
					transform.TransformComponent,
				))

				tr, ok := transform.GetTransform(entry)
				require.True(t, ok)

				tr.SetScale(2, 3)
				require.True(t, tr.IsDirty())

				// Re-fetch the transform to simulate a new access after modification
				tr, ok = transform.GetTransform(entry)
				require.True(t, ok)

				x, y := tr.Scale()
				require.Equal(t, 2.0, x)
				require.Equal(t, 3.0, y)
			},
		},
		{
			name: "Origin Accessors",
			Run: func(t *testing.T, world donburi.World) {
				entry := world.Entry(world.Create(
					transform.TransformComponent,
				))

				tr, ok := transform.GetTransform(entry)
				require.True(t, ok)

				tr.SetOrigin(4, 5)
				require.True(t, tr.IsDirty())

				// Re-fetch the transform to simulate a new access after modification
				tr, ok = transform.GetTransform(entry)
				require.True(t, ok)

				x, y := tr.Origin()
				require.Equal(t, 4.0, x)
				require.Equal(t, 5.0, y)
			},
		},
		{
			name: "Matrix Accessors",
			Run: func(t *testing.T, world donburi.World) {
				entry := world.Entry(world.Create(
					transform.TransformComponent,
				))

				tr, ok := transform.GetTransform(entry)
				require.True(t, ok)

				matrix := transform.GetTransformMatrix(entry, tr)
				require.Equal(t, matrix, ebiten.GeoM{})

				tr.SetPosition(1, 2)

				matrix = transform.GetTransformMatrix(entry, tr)
				require.NotEqual(t, matrix, ebiten.GeoM{})
			},
		},
		{
			name: "Matrix Applies Origin, Scale, Then Position",
			Run: func(t *testing.T, world donburi.World) {
				entry := world.Entry(world.Create(transform.TransformComponent))
				tr, _ := transform.GetTransform(entry)

				tr.SetOrigin(1, 1)
				tr.SetScale(2, 2)
				tr.SetPosition(10, 10)

				m := transform.GetTransformMatrix(entry, tr)
				x, y := m.Apply(1, 1)
				require.Equal(t, 10.0, x)
				require.Equal(t, 10.0, y)

				x, y = m.Apply(0, 0)
				require.Equal(t, 8.0, x)
				require.Equal(t, 8.0, y)
			},
		},
		{
			name: "Cached Matrix Updates And Clears Dirty",
			Run: func(t *testing.T, world donburi.World) {
				entry := world.Entry(world.Create(
					transform.TransformComponent,
					transform.MatrixComponent,
				))
				tr, _ := transform.GetTransform(entry)

				tr.SetPosition(1, 2)
				m := transform.GetTransformMatrix(entry, tr)
				require.False(t, tr.IsDirty())

				x, y := m.Apply(0, 0)
				require.Equal(t, 1.0, x)
				require.Equal(t, 2.0, y)

				cached := transform.MatrixComponent.Get(entry).GeoM
				require.Equal(t, cached, m)
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			world := donburi.NewWorld()
			tc.Run(t, world)
		})
	}
}

func BenchmarkTransform(b *testing.B) {
	var benchmarks = []struct {
		name string
		Run  func(b *testing.B, entries []*donburi.Entry, trs []*transform.TransformModel)
	}{
		{
			name: "Get Dirty Matrix With Transform",
			Run: func(b *testing.B, entries []*donburi.Entry, trs []*transform.TransformModel) {
				for i := range trs {
					trs[i].SetDirty(true)
					_ = transform.GetTransformMatrix(entries[i], trs[i])
				}
			},
		},
		{
			name: "Get Dirty Matrix Without Transform",
			Run: func(b *testing.B, entries []*donburi.Entry, trs []*transform.TransformModel) {
				for i := range trs {
					trs[i].SetDirty(true)
					_ = transform.GetMatrix(entries[i])
				}
			},
		},
		{
			name: "Get Cache Matrix With Transform",
			Run: func(b *testing.B, entries []*donburi.Entry, trs []*transform.TransformModel) {
				for i := range trs {
					_ = transform.GetTransformMatrix(entries[i], trs[i])
				}
			},
		},
		{
			name: "Get Cache Matrix Without Transform",
			Run: func(b *testing.B, entries []*donburi.Entry, trs []*transform.TransformModel) {
				for i := range trs {
					_ = transform.GetMatrix(entries[i])
				}
			},
		},
	}
	for _, bm := range benchmarks {
		for _, size := range benchSizes {
			b.Run(bm.name+"/size="+strconv.Itoa(size), func(b *testing.B) {
				entries, trs := makeBenchEntities(donburi.NewWorld(), size)
				b.ResetTimer()
				b.ReportAllocs()
				for b.Loop() {
					bm.Run(b, entries, trs)
				}
			})
		}
	}
}

func makeBenchEntities(world donburi.World, count int) ([]*donburi.Entry, []*transform.TransformModel) {
	entries := make([]*donburi.Entry, count)
	entities := world.CreateMany(
		count,
		transform.TransformComponent,
		transform.MatrixComponent,
	)
	for i, e := range entities {
		entries[i] = world.Entry(e)
	}
	trs := make([]*transform.TransformModel, 0, count)
	transform.TransformComponent.Each(world, func(e *donburi.Entry) {
		tr, _ := transform.GetTransform(e)
		tr.SetOrigin(8, 8)
		tr.SetScale(2, 2)
		tr.SetRotation(0.7)
		trs = append(trs, tr)
	})
	return entries, trs
}

// BenchmarkTransform/Get_Dirty_Matrix_With_Transform/size=1-18            99426847                12.08 ns/op            0 B/op          0 allocs/op
// BenchmarkTransform/Get_Dirty_Matrix_With_Transform/size=10-18           10127931               116.3 ns/op             0 B/op          0 allocs/op
// BenchmarkTransform/Get_Dirty_Matrix_With_Transform/size=100-18           1000000              1130 ns/op               0 B/op          0 allocs/op
// BenchmarkTransform/Get_Dirty_Matrix_With_Transform/size=1000-18            97090             12474 ns/op               0 B/op          0 allocs/op
// BenchmarkTransform/Get_Dirty_Matrix_With_Transform/size=10000-18            9868            119217 ns/op               0 B/op          0 allocs/op
// BenchmarkTransform/Get_Dirty_Matrix_Without_Transform/size=1-18         61739510                19.35 ns/op            0 B/op          0 allocs/op
// BenchmarkTransform/Get_Dirty_Matrix_Without_Transform/size=10-18         6534241               185.3 ns/op             0 B/op          0 allocs/op
// BenchmarkTransform/Get_Dirty_Matrix_Without_Transform/size=100-18         634248              1835 ns/op               0 B/op          0 allocs/op
// BenchmarkTransform/Get_Dirty_Matrix_Without_Transform/size=1000-18         62962             19132 ns/op               0 B/op          0 allocs/op
// BenchmarkTransform/Get_Dirty_Matrix_Without_Transform/size=10000-18         6160            188667 ns/op               0 B/op          0 allocs/op
// BenchmarkTransform/Get_Cache_Matrix_With_Transform/size=1-18            200012402                6.064 ns/op           0 B/op          0 allocs/op
// BenchmarkTransform/Get_Cache_Matrix_With_Transform/size=10-18           21825886                55.46 ns/op            0 B/op          0 allocs/op
// BenchmarkTransform/Get_Cache_Matrix_With_Transform/size=100-18           2338387               513.8 ns/op             0 B/op          0 allocs/op
// BenchmarkTransform/Get_Cache_Matrix_With_Transform/size=1000-18           215726              5533 ns/op               0 B/op          0 allocs/op
// BenchmarkTransform/Get_Cache_Matrix_With_Transform/size=10000-18           21288             56478 ns/op               0 B/op          0 allocs/op
// BenchmarkTransform/Get_Cache_Matrix_Without_Transform/size=1-18         100000000               11.55 ns/op            0 B/op          0 allocs/op
// BenchmarkTransform/Get_Cache_Matrix_Without_Transform/size=10-18        11296965               107.4 ns/op             0 B/op          0 allocs/op
// BenchmarkTransform/Get_Cache_Matrix_Without_Transform/size=100-18        1000000              1052 ns/op               0 B/op          0 allocs/op
// BenchmarkTransform/Get_Cache_Matrix_Without_Transform/size=1000-18        109059             10927 ns/op               0 B/op          0 allocs/op
// BenchmarkTransform/Get_Cache_Matrix_Without_Transform/size=10000-18        10000            107091 ns/op               0 B/op          0 allocs/op
