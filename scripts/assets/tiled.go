package assets

import (
	"encoding/xml"
	"log/slog"
	"path"

	"github.com/adm87/tiled"
)

func resolveSourcePath(basePath, source string) string {
	resolvedPath := path.Join(path.Dir(basePath), source)
	resolvedPath = path.Clean(resolvedPath)
	return resolvedPath
}

// ========== TMX Importer ==========

type tmxImporter struct {
	log *slog.Logger
}

func (ti *tmxImporter) AssetTypes() []string {
	return []string{"tmx"}
}

func (ti *tmxImporter) Import(handle AssetHandle, data []byte) (any, error) {
	var tmx *tiled.Tmx

	if err := xml.Unmarshal(data, &tmx); err != nil {
		ti.log.Error("Failed to unmarshal TMX", slog.String("error", err.Error()))
		return nil, err
	}

	for i := range tmx.Tilesets {
		tmx.Tilesets[i].Source = resolveSourcePath(string(handle), tmx.Tilesets[i].Source)
	}

	return tmx, nil
}

func TmxImporter(log *slog.Logger) AssetImporter {
	return &tmxImporter{log: log}
}

// ========== TSX Importer ==========

type tsxImporter struct {
	log *slog.Logger
}

func (tsi *tsxImporter) AssetTypes() []string {
	return []string{"tsx"}
}

func (tsi *tsxImporter) Import(handle AssetHandle, data []byte) (any, error) {
	var tsx *tiled.Tsx

	if err := xml.Unmarshal(data, &tsx); err != nil {
		tsi.log.Error("Failed to unmarshal TSX", slog.String("error", err.Error()))
		return nil, err
	}

	tsx.Image.Source = resolveSourcePath(string(handle), tsx.Image.Source)

	return tsx, nil
}

func TsxImporter(log *slog.Logger) AssetImporter {
	return &tsxImporter{log: log}
}

// ========== TX Importer ==========

type txImporter struct {
	log *slog.Logger
}

func (txi *txImporter) AssetTypes() []string {
	return []string{"tx"}
}

func (txi *txImporter) Import(handle AssetHandle, data []byte) (any, error) {
	var tx *tiled.Tx

	if err := xml.Unmarshal(data, &tx); err != nil {
		txi.log.Error("Failed to unmarshal TX", slog.String("error", err.Error()))
		return nil, err
	}

	tx.Tileset.Source = resolveSourcePath(string(handle), tx.Tileset.Source)

	return tx, nil
}

func TxImporter(log *slog.Logger) AssetImporter {
	return &txImporter{log: log}
}
