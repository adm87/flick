package aseprite

import (
	"gopkg.in/yaml.v3"
)

type Config struct {
	ResourceDir string          `yaml:"resource_directory"`
	Libraries   []LibraryConfig `yaml:"libraries"`
}

type LibraryConfig struct {
	Name    string               `yaml:"name"`
	Content LibraryContentConfig `yaml:"content"`
}

type LibraryContentConfig struct {
	Image string `yaml:"image"`
	Json  string `yaml:"json"`
}

func LoadConfig(raw []byte) (*Config, error) {
	var cfg Config
	err := yaml.Unmarshal(raw, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
