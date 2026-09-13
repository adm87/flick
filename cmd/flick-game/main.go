package main

import (
	"log"

	"github.com/adm87/flick/cmd/flick-game/internal"

	_ "github.com/adm87/flick/cmd/flick-game/internal/diagnostics"
)

func main() {
	if err := internal.Run(); err != nil {
		log.Fatalf("exited with error: %v", err)
	}
}
