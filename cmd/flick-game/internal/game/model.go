package game

type Model struct {
	Config    *Config
	Renderers *Renderers
}

type Renderers struct {
	ImageRendererID uint64
}
