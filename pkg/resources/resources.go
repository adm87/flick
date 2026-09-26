package resources

import (
	"errors"
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"
	"sync"

	"github.com/adm87/flick/pkg/logger"
)

var (
	ErrDuplicateImporterRegister = errors.New("duplicate importer register")
)

type ResourceHandle = uint64

type ResourcePath string

func (rp ResourcePath) String() string {
	return string(rp)
}

func (rp ResourcePath) Ext() string {
	ext := strings.Trim(filepath.Ext(string(rp)), ".")
	return strings.ToLower(ext)
}

type ResourceType string

func (rt ResourceType) String() string {
	return string(rt)
}

func (rt ResourceType) IsEmpty() bool {
	return rt == ""
}

type Resources struct {
	logger    logger.Logger
	importers map[ResourceType]ResourceImporter
	handles   map[ResourcePath]ResourceHandle
	mu        sync.Mutex
}

func NewResources(log logger.Logger) *Resources {
	return &Resources{
		importers: make(map[ResourceType]ResourceImporter),
		handles:   make(map[ResourcePath]ResourceHandle),
		logger:    log,
	}
}

func (r *Resources) GetHandle(path ResourcePath) (ResourceHandle, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	handle, exists := r.handles[path]
	if !exists {
		return 0, fmt.Errorf("%w: %s", ErrNotFound, path.String())
	}
	return handle, nil
}

func (r *Resources) RegisterImporter(importer ResourceImporter, resourceTypes []ResourceType) error {
	if importer == nil {
		return nil
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	for _, resourceType := range resourceTypes {
		if resourceType.IsEmpty() {
			continue
		}

		str := strings.Trim(resourceType.String(), ".")
		str = strings.ToLower(str)

		rt := ResourceType(str)

		if _, exists := r.importers[rt]; exists {
			return fmt.Errorf("%w: %s", ErrDuplicateImporterRegister, str)
		}

		r.importers[rt] = importer
	}

	return nil
}

func (r *Resources) Load(filesystem fs.FS, paths ...ResourcePath) error {
	return r.runOperation(paths, func(path ResourcePath) error {
		return r.loadResource(filesystem, path)
	})
}

func (r *Resources) Unload(paths ...ResourcePath) error {
	return r.runOperation(paths, r.unloadResource)
}

func (r *Resources) runOperation(paths []ResourcePath, operation func(path ResourcePath) error) error {
	seen := make(map[ResourcePath]struct{})
	errCh := make(chan error, len(paths))
	var wg sync.WaitGroup

	for _, path := range paths {
		if _, exists := seen[path]; exists {
			continue
		}
		seen[path] = struct{}{}

		wg.Add(1)
		go func(p ResourcePath) {
			defer wg.Done()
			if err := operation(p); err != nil {
				errCh <- err
			}
		}(path)
	}

	wg.Wait()
	close(errCh)

	var errs []error
	for err := range errCh {
		errs = append(errs, err)
	}

	return errors.Join(errs...)
}

func (r *Resources) loadResource(filesystem fs.FS, path ResourcePath) error {
	ext := path.Ext()
	rt := ResourceType(ext)

	if r.checkExists(path) {
		return fmt.Errorf("%w: %s", ErrDuplicate, path.String())
	}

	r.mu.Lock()
	importer, exists := r.importers[rt]
	r.mu.Unlock()

	if !exists {
		return fmt.Errorf("%w: %s", ErrUnknownType, ext)
	}

	data, err := fs.ReadFile(filesystem, path.String())
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return fmt.Errorf("%w: %w", ErrNotFound, err)
		}
		return err
	}

	handle, err := importer.Import(data)
	if err != nil {
		return err
	}

	r.mu.Lock()
	r.handles[path] = handle
	r.mu.Unlock()

	return nil
}

func (r *Resources) unloadResource(path ResourcePath) error {
	handle, exists := r.getHandle(path)
	if !exists {
		return fmt.Errorf("%w: %s", ErrNotFound, path.String())
	}

	r.mu.Lock()
	delete(r.handles, path)
	r.mu.Unlock()

	ext := path.Ext()
	rt := ResourceType(ext)

	r.mu.Lock()
	importer, exists := r.importers[rt]
	r.mu.Unlock()

	if !exists {
		return fmt.Errorf("%w: %s", ErrUnknownType, path.String())
	}

	err := importer.Delete(handle)
	if err != nil {
		return fmt.Errorf("%w: %s", err, path.String())
	}

	return nil
}

func (r *Resources) checkExists(path ResourcePath) bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	_, exists := r.handles[path]
	return exists
}

func (r *Resources) getHandle(path ResourcePath) (ResourceHandle, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	handle, exists := r.handles[path]
	return handle, exists
}
