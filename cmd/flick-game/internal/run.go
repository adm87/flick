package internal

import (
	"os"

	"github.com/adm87/flick/pkg/engine"
	"github.com/adm87/flick/pkg/logger"
)

func Run() error {
	log := logger.NewLogger(os.Stdout)

	return engine.Run(
		engine.WithLogger(log),
	)
}
