package game

import (
	"github.com/adm87/flick/pkg/data"
	"github.com/adm87/flick/pkg/images"
	"github.com/adm87/flick/pkg/resources"
)

type Assets struct {
	Resources *resources.Resources
	Data      *data.DataStore
	Images    *images.ImageStore
}
