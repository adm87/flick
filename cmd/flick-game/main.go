package main

import (
	"log"

	_ "github.com/adm87/flick/cmd/flick-game/internal/diagnostics"

	"github.com/adm87/flick/cmd/flick-game/internal/game"
)

func main() {
	if err := game.Run(); err != nil {
		log.Fatal(err)
	}
}
