package game

import (
	"context"
	"log/slog"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
)

const (
	TargetWidth  = 1280
	TargetHeight = 720

	RenderScale = 0.15
)

// Screen represents the game's screen dimensions.
type Screen struct {
	Width  float32 // TargetWidth with applied RenderScale
	Height float32 // TargetHeight with applied RenderScale
}

// Context represents the game context, providing access to the ECS world, logger, and screen.
type Context interface {
	Context() context.Context
	ECS() donburi.World
	Log() *slog.Logger
	Screen() Screen
}

// Game represents the main game interface, extending Context with game-specific methods.
type Game interface {
	Context

	// OnStart callbacks are called right before the game loop starts.
	OnStart(func(g Game) error)

	// SetState changes the current game state.
	SetState(state State) error

	// AddUpdateSystems registers update systems for a specific phase.
	// The order systems are added determines the order they are executed within that phase.
	AddUpdateSystems(phase Phase, systems ...UpdateSystem)

	// AddDrawSystems registers draw systems.
	// The order systems are added determines the order they are executed.
	AddDrawSystems(systems ...DrawSystem)

	// ClearSystems removes all registered update and draw systems.
	ClearSystems()

	// Run starts the game loop.
	Run() error
}

type gameContext struct {
	ctx   context.Context
	world donburi.World

	screen Screen

	updateSystems map[Phase][]UpdateSystem
	drawSystems   []DrawSystem

	onStartCallbacks []func(g Game) error

	logger *slog.Logger
	sm     *statemachine
}

func NewGame(ctx context.Context) Game {
	l := slog.Default().With(slog.String("version", "0"))
	w := donburi.NewWorld()
	s := Screen{float32(TargetWidth) * RenderScale, float32(TargetHeight) * RenderScale}
	return &gameContext{
		ctx:           ctx,
		logger:        l,
		world:         w,
		screen:        s,
		updateSystems: make(map[Phase][]UpdateSystem),
		drawSystems:   make([]DrawSystem, 0),
		sm:            NewStateMachine(),
	}
}

func (g *gameContext) setupWindow() {
	ebiten.SetWindowTitle("Game")
	ebiten.SetWindowSize(TargetWidth, TargetHeight)
}

func (g *gameContext) onStart() error {
	for _, callback := range g.onStartCallbacks {
		if err := callback(g); err != nil {
			return err
		}
	}
	return nil
}

// ========== Context interface implementation ==========

func (g *gameContext) Context() context.Context {
	return g.ctx
}

func (g *gameContext) ECS() donburi.World {
	return g.world
}

func (g *gameContext) Log() *slog.Logger {
	return g.logger
}

func (g *gameContext) Screen() Screen {
	return g.screen
}

func (g *gameContext) OnStart(f func(g Game) error) {
	g.onStartCallbacks = append(g.onStartCallbacks, f)
}

func (g *gameContext) SetState(state State) error {
	return g.sm.ChangeState(g, state)
}

func (g *gameContext) Run() error {
	g.setupWindow()

	if err := g.onStart(); err != nil {
		return err
	}

	return ebiten.RunGame(g)
}

// ========== Ebiten game interface implementation ==========

func (g *gameContext) Update() error {
	if err := g.callUpdatePhase(EarlyUpdatePhase, 1); err != nil {
		return err
	}

	// TASK: Calculate number of fixed update steps based on elapsed time
	if err := g.callUpdatePhase(FixedUpdatePhase, 1); err != nil {
		return err
	}

	if err := g.callUpdatePhase(LateUpdatePhase, 1); err != nil {
		return err
	}

	return nil
}

func (g *gameContext) Draw(screen *ebiten.Image) {
	g.callDrawPhase(screen)
}

func (g *gameContext) Layout(outsideWidth, outsideHeight int) (int, int) {
	return int(g.screen.Width), int(g.screen.Height)
}
