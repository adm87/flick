package main

import (
	"bytes"
	"fmt"
	"go/format"
	"os"
	"path/filepath"
	"text/template"

	_ "embed"

	"github.com/adm87/flick/content"
	"github.com/adm87/flick/pkg/aseprite"
	"github.com/adm87/flick/pkg/assert"
	"github.com/adm87/flick/pkg/templating"
	"github.com/alexflint/go-arg"
)

//go:embed template.tmpl
var templateFile []byte

type cmdArgs struct {
	WorkingDir string `arg:"--working-dir" help:"working directory"`
	Output     string `arg:"--output" help:"output directory for generated keys"`
}

type libaryModel struct {
	Model    *aseprite.Model
	Name     string
	ImgPath  string
	JsonPath string
}

func main() {
	cfg, err := aseprite.LoadConfig(content.AsepriteConfig)
	assert.NoError(err)

	var v cmdArgs
	arg.MustParse(&v)

	output, err := filepath.Abs(v.Output)
	assert.NoError(err)

	assert.NoError(ensureDest(output))

	var models []*libaryModel

	root := filepath.Join(v.WorkingDir, cfg.ResourceDir)
	for i := range cfg.Libraries {
		lib := &cfg.Libraries[i]
		jsonPath := filepath.Join(root, lib.Name, lib.Content.Json)

		path, err := filepath.Abs(jsonPath)
		assert.NoError(err)

		model, err := getLibraryModel(path)
		if err != nil {
			assert.NoError(fmt.Errorf("failed to load model for %s: %w", path, err))
		}

		models = append(models, &libaryModel{
			Name:     lib.Name,
			Model:    model,
			ImgPath:  filepath.Join(lib.Name, lib.Content.Image),
			JsonPath: filepath.Join(lib.Name, lib.Content.Json),
		})
	}

	input := struct {
		Libraries []*libaryModel
	}{
		Libraries: models,
	}

	tmpl, err := template.
		New("generated").
		Funcs(templating.Functions()).
		Parse(string(templateFile))
	assert.NoError(err)

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, input)
	assert.NoError(err)

	outputFile := filepath.Join(output, "generated.go")

	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		_ = os.WriteFile(outputFile, buf.Bytes(), 0644)
		assert.NoError(fmt.Errorf("generated source is invalid: %w (raw output written to %s)", err, outputFile))
		return
	}

	err = os.WriteFile(outputFile, formatted, 0644)
	assert.NoError(err)
}

func ensureDest(dest string) error {
	stat, err := os.Stat(dest)
	if err == nil {
		if !stat.IsDir() {
			return fmt.Errorf("output path %q exists and is not a directory", dest)
		}
		return nil
	}
	if !os.IsNotExist(err) {
		return fmt.Errorf("failed to stat output path %q: %w", dest, err)
	}
	return os.MkdirAll(dest, 0755)
}

func getLibraryModel(path string) (*aseprite.Model, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return aseprite.LoadModel(data)
}
