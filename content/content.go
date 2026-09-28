package content

import _ "embed"

//go:embed embedded/game.yaml
var GameConfig []byte

//go:embed embedded/aseprite.yaml
var AsepriteConfig []byte
