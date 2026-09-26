package images

import (
	"bytes"

	"github.com/adm87/flick/pkg/resources"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var imageTypes = []resources.ResourceType{"png", "jpg", "jpeg", "gif"}

func ImageTypes() []resources.ResourceType {
	return imageTypes
}

type ImageImporter struct {
	store *ImageStore
}

func NewImageImporter(store *ImageStore) *ImageImporter {
	return &ImageImporter{
		store: store,
	}
}

func (ii *ImageImporter) Import(raw []byte) (resources.ResourceHandle, error) {
	img, _, err := ebitenutil.NewImageFromReader(bytes.NewBuffer(raw))
	if err != nil {
		return 0, err
	}

	handle := ii.store.Put(img)
	return handle, nil
}

func (ii *ImageImporter) Delete(handle resources.ResourceHandle) error {
	if err := ii.store.Delete(handle); err != nil {
		return err
	}
	return nil
}
