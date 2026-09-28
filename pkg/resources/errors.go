package resources

import (
	"errors"

	"github.com/adm87/flick/pkg/types/structures/slotmap"
)

var (
	ErrInvalidHandle = errors.New("invalid handle")
	ErrNotFound      = errors.New("not found")
	ErrImportFailed  = errors.New("import failed")
	ErrDeleteFailed  = errors.New("delete failed")
	ErrUnknownType   = errors.New("unknown type")
	ErrDuplicate     = errors.New("duplicate")
)

// GetError translates slotmap errors into resource package errors.
func GetError(err error) error {
	switch {
	case errors.Is(err, slotmap.ErrInvalidKey):
		return ErrInvalidHandle
	case errors.Is(err, slotmap.ErrNotFound):
		return ErrNotFound
	default:
		return err
	}
}
