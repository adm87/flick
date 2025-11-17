package main

import (
	"context"
	"errors"
	"os"

	"github.com/adm87/flick/cmd"
	"github.com/adm87/flick/scripts"
	"github.com/hajimehoshi/ebiten/v2"
)

var version = "0.0.0-unreleased"

func main() {
	g := scripts.NewGame(context.Background(), version)
	if err := cmd.Boot(g).ExecuteContext(g.Context()); err != nil && !errors.Is(err, ebiten.Termination) {
		os.Exit(1)
	}
}
