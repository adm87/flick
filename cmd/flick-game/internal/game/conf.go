package game

import "gopkg.in/yaml.v3"

type Config struct {
	Window WindowConfig `yaml:"window"`
}

type WindowConfig struct {
	Width      int  `yaml:"width"`
	Height     int  `yaml:"height"`
	Fullscreen bool `yaml:"fullscreen"`
}

func NewConfig(raw []byte) (*Config, error) {
	var cfg Config
	err := yaml.Unmarshal(raw, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
