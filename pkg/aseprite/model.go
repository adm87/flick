package aseprite

import (
	"encoding/json"

	"github.com/adm87/flick/pkg/types"
)

type Model struct {
	Meta   *Meta    `json:"meta"`
	Frames []*Frame `json:"frames"`
}

type Meta struct {
	App       string     `json:"app"`
	Version   string     `json:"version"`
	Image     string     `json:"image"`
	Format    string     `json:"format"`
	Scale     string     `json:"scale"`
	FrameTags []FrameTag `json:"frameTags"`
	Size      FrameRect  `json:"size"`
}

type Frame struct {
	FileName         string    `json:"filename"`
	Rotated          bool      `json:"rotated"`
	Trimmed          bool      `json:"trimmed"`
	Duration         int       `json:"duration"`
	Frame            FrameRect `json:"frame"`
	SpriteSourceSize FrameRect `json:"spriteSourceSize"`
	SourceSize       FrameRect `json:"sourceSize"`
}

type FrameRect struct {
	X      int `json:"x"`
	Y      int `json:"y"`
	Width  int `json:"w"`
	Height int `json:"h"`
}

type FrameTag struct {
	Name      string         `json:"name"`
	Direction string         `json:"direction"`
	From      int            `json:"from"`
	To        int            `json:"to"`
	Color     types.HexColor `json:"color"`
}

func LoadModel(data []byte) (*Model, error) {
	var m Model
	err := json.Unmarshal(data, &m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}
