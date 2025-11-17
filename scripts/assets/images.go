package assets

import (
	"bytes"
	"log/slog"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type imageImporter struct {
	log *slog.Logger
}

func (ii *imageImporter) AssetTypes() []string {
	return []string{"png", "jpg", "jpeg"}
}

func (ii *imageImporter) Import(handle AssetHandle, data []byte) (any, error) {
	img, _, err := ebitenutil.NewImageFromReader(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	return img, nil
}

func ImageImporter(log *slog.Logger) AssetImporter {
	return &imageImporter{log: log}
}
