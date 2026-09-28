package types_test

import (
	"encoding/json"
	"testing"

	"github.com/adm87/flick/pkg/types"
	"gopkg.in/yaml.v3"
)

const (
	hexJson string = `{"color":"#ffffffff"}`
	hexYaml string = "color: \"#ffffffff\""
)

type testModel struct {
	Color types.HexColor `json:"color" yaml:"color"`
}

func TestHexColorUnmarshal(t *testing.T) {
	t.Run("JSON Unmarshal", func(t *testing.T) {
		var h testModel
		err := json.Unmarshal([]byte(hexJson), &h)
		if err != nil {
			t.Fatalf("Failed to unmarshal JSON: %v", err)
		}
		if h.Color != (types.HexColor{R: 0xff, G: 0xff, B: 0xff, A: 0xff}) {
			t.Errorf("Expected color %#v, got %#v", types.HexColor{R: 0xff, G: 0xff, B: 0xff, A: 0xff}, h.Color)
		}
	})
	t.Run("YAML Unmarshal", func(t *testing.T) {
		var h testModel
		err := yaml.Unmarshal([]byte(hexYaml), &h)
		if err != nil {
			t.Fatalf("Failed to unmarshal YAML: %v", err)
		}
		if h.Color != (types.HexColor{R: 0xff, G: 0xff, B: 0xff, A: 0xff}) {
			t.Errorf("Expected color %#v, got %#v", types.HexColor{R: 0xff, G: 0xff, B: 0xff, A: 0xff}, h.Color)
		}
	})
}
