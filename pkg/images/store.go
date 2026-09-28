package images

import (
	"fmt"
	"image"

	"github.com/adm87/flick/pkg/resources"
	"github.com/adm87/flick/pkg/types/structures/slotmap"
	"github.com/hajimehoshi/ebiten/v2"
)

var ErrFrameOutOfBounds = fmt.Errorf("frame out of bounds")

type ImageStore struct {
	images *slotmap.SlotMap[*ebiten.Image]
	frames map[resources.ResourceHandle][]*ebiten.Image
}

func NewImageStore() *ImageStore {
	return &ImageStore{
		images: slotmap.New[*ebiten.Image](256),
		frames: make(map[resources.ResourceHandle][]*ebiten.Image),
	}
}

func (is *ImageStore) Get(handle resources.ResourceHandle) (*ebiten.Image, error) {
	k := slotmap.Unpack(handle)

	img, err := is.images.Get(k)
	if err != nil {
		return nil, resources.GetError(err)
	}

	return img, nil
}

func (is *ImageStore) Set(handle resources.ResourceHandle, img *ebiten.Image) error {
	k := slotmap.Unpack(handle)

	if err := is.images.Set(k, img); err != nil {
		return resources.GetError(err)
	}

	return nil
}

func (is *ImageStore) Put(img *ebiten.Image) resources.ResourceHandle {
	k := is.images.Insert(img)
	return k.Pack()
}

func (is *ImageStore) Delete(handle resources.ResourceHandle) error {
	k := slotmap.Unpack(handle)

	img, err := is.images.Delete(k)
	if err != nil {
		return resources.GetError(err)
	}

	img.Deallocate()
	is.ClearFrames(handle)

	return nil
}

func (is *ImageStore) GetFrame(handle resources.ResourceHandle, frame int) (*ebiten.Image, error) {
	if frames, ok := is.frames[handle]; ok {
		if frame < 0 || frame >= len(frames) {
			return nil, ErrFrameOutOfBounds
		}
		return frames[frame], nil
	}
	// No frames found for this handle, fallback to the main image
	return is.Get(handle)
}

// SliceFramesUniform slices the main image into uniform frames based on the specified frame width and height.
func (is *ImageStore) SliceFramesUniform(handle resources.ResourceHandle, frameWidth, frameHeight int) {
	img, err := is.Get(handle)
	if err != nil {
		return
	}

	bounds := img.Bounds()
	var frames []image.Rectangle
	for y := bounds.Min.Y; y < bounds.Max.Y; y += frameHeight {
		for x := bounds.Min.X; x < bounds.Max.X; x += frameWidth {
			frames = append(frames, image.Rect(x, y, x+frameWidth, y+frameHeight))
		}
	}

	is.SliceFrames(handle, frames)
}

// SliceFrames slices the main image into frames based on the specified rectangles.
func (is *ImageStore) SliceFrames(handle resources.ResourceHandle, frames []image.Rectangle) {
	img, err := is.Get(handle)
	if err != nil {
		return
	}

	is.ClearFrames(handle)

	slicedFrames := make([]*ebiten.Image, len(frames))
	for i, rect := range frames {
		subImg := img.SubImage(rect).(*ebiten.Image)
		slicedFrames[i] = subImg
	}

	is.frames[handle] = slicedFrames
}

// ClearFrames removes all sliced frames associated with the given resource handle.
func (is *ImageStore) ClearFrames(handle resources.ResourceHandle) {
	delete(is.frames, handle)
}
