package game

import (
	"errors"
	"log/slog"

	"github.com/hajimehoshi/ebiten/v2"
)

type Phase uint8

const (
	EarlyUpdatePhase Phase = iota
	FixedUpdatePhase
	LateUpdatePhase
)

func (p Phase) String() string {
	switch p {
	case EarlyUpdatePhase:
		return "EarlyUpdatePhase"
	case FixedUpdatePhase:
		return "FixedUpdatePhase"
	case LateUpdatePhase:
		return "LateUpdatePhase"
	default:
		return "UnknownPhase"
	}
}

type UpdateSystem func(ctx Context) error
type DrawSystem func(ctx Context, screen *ebiten.Image) error

func (g *gameContext) AddUpdateSystems(phase Phase, systems ...UpdateSystem) {
	if _, exists := g.updateSystems[phase]; !exists {
		g.updateSystems[phase] = []UpdateSystem{}
	}
	g.updateSystems[phase] = append(g.updateSystems[phase], systems...)
}

func (g *gameContext) AddDrawSystems(systems ...DrawSystem) {
	g.drawSystems = append(g.drawSystems, systems...)
}

func (g *gameContext) ClearSystems() {
	g.updateSystems = make(map[Phase][]UpdateSystem)
	g.drawSystems = make([]DrawSystem, 0)
}

func (g *gameContext) callUpdatePhase(phase Phase, steps int) error {
	systems, exists := g.updateSystems[phase]
	if !exists {
		return nil // No systems registered for this phase
	}
	for range max(0, steps) {
		for _, system := range systems {
			if err := system(g); err != nil {
				if errors.Is(err, ebiten.Termination) {
					g.logger.Info("termination requested")
				} else {
					g.logger.Error("update error", slog.String("phase", phase.String()), slog.Any("error", err))
				}
				return err
			}
		}
	}
	return nil
}

func (g *gameContext) callDrawPhase(screen *ebiten.Image) {
	for _, system := range g.drawSystems {
		if err := system(g, screen); err != nil {
			g.logger.Error("draw error", slog.Any("error", err))
			return
		}
	}
}
