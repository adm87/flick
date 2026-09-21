package rendering

import (
	"slices"

	"github.com/adm87/flick/pkg/ecs/components/renderable"
	"github.com/yohamta/donburi"
)

type sortkey uint64

const (
	idxBits = 24
	idxMask = 1<<idxBits - 1
)

// makeSortKey creates a sort key from the layer, z-index, and candidate index.
func makeSortKey(layer int8, z int32, idx int) sortkey {
	l := uint64(uint8(layer) ^ 0x80)
	zz := uint64(uint32(z) ^ 0x80000000)
	return sortkey(l<<56 | zz<<24 | uint64(idx))
}

func (k sortkey) index() int {
	return int(k & idxMask)
}

type RenderingCandidate struct {
	Renderable *renderable.RenderableModel
	Entry      *donburi.Entry
}

func (rc *RenderingCandidate) Reset() {
	rc.Renderable = nil
	rc.Entry = nil
}

type RenderingQueue struct {
	candidates []RenderingCandidate
	sortKeys   []sortkey
}

func NewRenderingQueue() *RenderingQueue {
	return &RenderingQueue{
		candidates: make([]RenderingCandidate, 0, 256),
		sortKeys:   make([]sortkey, 0, 256),
	}
}

func (rq *RenderingQueue) Reset() {
	rq.candidates = rq.candidates[:0]
	rq.sortKeys = rq.sortKeys[:0]
}

func (rq *RenderingQueue) Enqueue(entry *donburi.Entry, r *renderable.RenderableModel) {
	rq.candidates = append(rq.candidates, RenderingCandidate{Renderable: r, Entry: entry})
	rq.sortKeys = append(rq.sortKeys, makeSortKey(r.Layer(), r.ZIndex(), len(rq.candidates)-1))
}

func (rq *RenderingQueue) Candidates() []RenderingCandidate {
	return rq.candidates
}

func (rq *RenderingQueue) sortedKeys() []sortkey {
	slices.Sort(rq.sortKeys)
	return rq.sortKeys
}
