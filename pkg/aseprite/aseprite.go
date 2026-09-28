package aseprite

import (
	"errors"
	"fmt"
	"image"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/adm87/flick/pkg/data"
	"github.com/adm87/flick/pkg/images"
	"github.com/adm87/flick/pkg/resources"
)

var (
	ErrLibraryNotFound = errors.New("library not found")
	ErrLibraryError    = errors.New("library error")
)

type Aseprite struct {
	config *Config
	data   *data.DataStore
	images *images.ImageStore

	contentFS fs.FS
	libraries map[resources.ResourceHandle]*Model
}

func New(config *Config, dataStore *data.DataStore, imageStore *images.ImageStore, workingDir string) *Aseprite {
	contentFS := os.DirFS(filepath.Join(workingDir, config.ResourceDir))
	return &Aseprite{
		config:    config,
		data:      dataStore,
		images:    imageStore,
		contentFS: contentFS,
		libraries: make(map[resources.ResourceHandle]*Model),
	}
}

func (a *Aseprite) ContentFS() fs.FS {
	return a.contentFS
}

func (a *Aseprite) BuildLibrary(res *resources.Resources, libraryName string) error {
	lib, err := a.getLibraryConfig(libraryName)
	if err != nil {
		return err
	}

	imgPath := filepath.Join(libraryName, lib.Content.Image)
	imgHandle, err := res.GetHandle(resources.ResourcePath(imgPath))
	if err != nil {
		return fmt.Errorf("%w: %w", ErrLibraryError, err)
	}

	jsonPath := filepath.Join(libraryName, lib.Content.Json)
	jsonHandle, err := res.GetHandle(resources.ResourcePath(jsonPath))
	if err != nil {
		return fmt.Errorf("%w: %w", ErrLibraryError, err)
	}

	model, err := a.getModel(jsonHandle, a.data)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrLibraryError, err)
	}

	frames := a.getFrameRects(model.Frames)
	a.images.SliceFrames(imgHandle, frames)

	a.libraries[imgHandle] = model
	return nil
}

func (a *Aseprite) getLibraryConfig(libraryName string) (*LibraryConfig, error) {
	for i := range a.config.Libraries {
		lib := &a.config.Libraries[i]

		if lib.Name == libraryName {
			return lib, nil
		}
	}
	return nil, fmt.Errorf("%w: %s", ErrLibraryNotFound, libraryName)
}

func (a *Aseprite) getModel(handle resources.ResourceHandle, data *data.DataStore) (*Model, error) {
	raw, err := data.Get(handle)

	if err != nil {
		return nil, err
	}

	return LoadModel(raw)
}

func (a *Aseprite) getFrameRects(frames []*Frame) []image.Rectangle {
	rects := make([]image.Rectangle, len(frames))

	for i := range frames {
		frame := frames[i].Frame
		rects[i] = image.Rect(frame.X, frame.Y, frame.X+frame.Width, frame.Y+frame.Height)
	}

	return rects
}
