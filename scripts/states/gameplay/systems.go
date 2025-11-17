package gameplay

import (
	"image/color"

	"github.com/adm87/flick/scripts/actors"
	"github.com/adm87/flick/scripts/components"
	"github.com/adm87/flick/scripts/game"
	"github.com/adm87/flick/scripts/systems/camera"
	"github.com/adm87/flick/scripts/systems/debug"
	"github.com/hajimehoshi/ebiten/v2"
)

func (s *state) registerSystems(g game.Game) {

	// ============ Early Update Systems ============

	g.AddUpdateSystems(game.EarlyUpdatePhase,
		debug.PollDebugInput,
	)

	// ============ Late Update Systems ============

	g.AddUpdateSystems(game.LateUpdatePhase,

		// Camera Follow Player
		func(ctx game.Context) error {
			return camera.FollowTarget(ctx, actors.Player.MustFirst(ctx.ECS()))
		},

		// Camera Bounds Check
		func(ctx game.Context) error {
			worldBoundsEntry := actors.WorldBounds.MustFirst(ctx.ECS())
			bounds := components.Rectangle.Get(worldBoundsEntry).Bounds(0, 0)
			return camera.BoundsCheck(ctx, bounds)
		},
	)

	// ============ Draw Systems ============

	g.AddDrawSystems(

		// Debug Colliders
		func(ctx game.Context, screen *ebiten.Image) error {
			debugEntry := actors.Debug.MustFirst(ctx.ECS())
			if components.Debug.Get(debugEntry).ShowColliders() {
				view := components.Transform.Get(actors.Camera.MustFirst(ctx.ECS())).InvMatrix()
				if err := debug.DrawEntityColliders(ctx, screen, view, actors.Solid.Iter(ctx.ECS()), color.RGBA{B: 255, A: 255}); err != nil {
					return err
				}
			}
			return nil
		},

		// Debug Collision Grid
		func(ctx game.Context, screen *ebiten.Image) error {
			debugEntry := actors.Debug.MustFirst(ctx.ECS())
			if components.Debug.Get(debugEntry).ShowStaticGrid() {
				view := components.Transform.Get(actors.Camera.MustFirst(ctx.ECS())).InvMatrix()
				cells := s.world.QueryCells(camera.Viewport(ctx))
				if err := debug.DrawCollisionGrid(ctx, screen, view, cells, GridCellSize, color.RGBA{R: 255, A: 255}); err != nil {
					return err
				}
			}
			return nil
		},

		// Debug Player Info
		func(ctx game.Context, screen *ebiten.Image) error {
			debugEntry := actors.Debug.MustFirst(ctx.ECS())
			if components.Debug.Get(debugEntry).ShowPlayer() {
				view := components.Transform.Get(actors.Camera.MustFirst(ctx.ECS())).InvMatrix()
				if err := debug.DrawEntityColliders(ctx, screen, view, actors.Player.Iter(ctx.ECS()), color.RGBA{G: 255, A: 255}); err != nil {
					return err
				}
			}
			return nil
		},

		// Debug FPS
		func(ctx game.Context, screen *ebiten.Image) error {
			debugEntry := actors.Debug.MustFirst(ctx.ECS())
			if components.Debug.Get(debugEntry).ShowFPS() {
				return debug.DrawFPS(ctx, screen)
			}
			return nil
		},
	)
}
