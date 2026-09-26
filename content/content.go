package content

import (
	"embed"

	"github.com/adm87/flick/pkg/resources"
)

const ConfResourcePath resources.ResourcePath = "embedded/conf.yaml"

//go:embed embedded
var embeddedFS embed.FS

func EmbeddedFS() embed.FS {
	return embeddedFS
}
