package engine

import (
	"context"
	"errors"
	"os/signal"
	"syscall"
	"time"

	"github.com/adm87/flick/pkg/logger"
	"github.com/hajimehoshi/ebiten/v2"
)

type RunOptions struct {
	Update   GameUpdate
	Draw     GameDraw
	Layout   GameLayout
	Shutdown GameShutdown

	Logger logger.Logger

	GameOptions *ebiten.RunGameOptions

	WindowTitle               string
	WindowWidth, WindowHeight int
	FPS                       int
	CursorMode                int
	Fullscreen                bool
}

type RunOption func(*RunOptions)

func Run(opts ...RunOption) error {
	options := &RunOptions{
		Update:       &noopUpdate{},
		Draw:         &noopDraw{},
		Layout:       &noopLayout{},
		Shutdown:     &noopShutdown{},
		Logger:       logger.N,
		WindowTitle:  "Untitled Game",
		WindowWidth:  800,
		WindowHeight: 600,
		FPS:          60,
		CursorMode:   0,
		Fullscreen:   false,
	}

	for _, opt := range opts {
		opt(options)
	}

	ebiten.SetWindowTitle(options.WindowTitle)
	ebiten.SetWindowSize(options.WindowWidth, options.WindowHeight)
	ebiten.SetFullscreen(options.Fullscreen)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetCursorMode(ebiten.CursorModeType(options.CursorMode))

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	var errs []error

	game := &shell{
		ctx:      ctx,
		update:   options.Update,
		draw:     options.Draw,
		layout:   options.Layout,
		shutdown: options.Shutdown,
		logger:   options.Logger,
		time:     newGameTime(options.FPS),
	}

	if err := ebiten.RunGameWithOptions(game, options.GameOptions); err != nil {
		errs = append(errs, err)
	}

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	if err := game.Shutdown(shutdownCtx); err != nil {
		errs = append(errs, err)
	}

	return errors.Join(errs...)
}

func WithWindowTitle(title string) RunOption {
	return func(opts *RunOptions) {
		opts.WindowTitle = title
	}
}

func WithWindowSize(width, height int) RunOption {
	return func(opts *RunOptions) {
		opts.WindowWidth = width
		opts.WindowHeight = height
	}
}

func WithFPS(fps int) RunOption {
	return func(opts *RunOptions) {
		opts.FPS = fps
	}
}

func WithCursorMode(cursorMode int) RunOption {
	return func(opts *RunOptions) {
		opts.CursorMode = cursorMode
	}
}

func WithFullscreen(fullscreen bool) RunOption {
	return func(opts *RunOptions) {
		opts.Fullscreen = fullscreen
	}
}

func WithGameOptions(gameOptions *ebiten.RunGameOptions) RunOption {
	return func(opts *RunOptions) {
		opts.GameOptions = gameOptions
	}
}

func WithShutdown(shutdown GameShutdown) RunOption {
	return func(opts *RunOptions) {
		opts.Shutdown = shutdown
	}
}

func WithUpdate(update GameUpdate) RunOption {
	return func(opts *RunOptions) {
		opts.Update = update
	}
}

func WithDraw(draw GameDraw) RunOption {
	return func(opts *RunOptions) {
		opts.Draw = draw
	}
}

func WithLayout(layout GameLayout) RunOption {
	return func(opts *RunOptions) {
		opts.Layout = layout
	}
}

func WithLogger(logger logger.Logger) RunOption {
	return func(opts *RunOptions) {
		opts.Logger = logger
	}
}
