package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tonitienda/procedural-animations-go/src/pkg/scenarios"
	"github.com/tonitienda/procedural-animations-go/src/pkg/systems"
	"github.com/tonitienda/procedural-animations-go/src/pkg/world"
)

type Game struct {
	world              *world.World
	circleRenderSystem *systems.CircleRenderSystem
}

func NewGame() *Game {
	w := world.NewWorld()

	rs := systems.NewCircleRenderSystem(w)

	return &Game{
		world:              w,
		circleRenderSystem: rs,
	}
}

func (g *Game) Update() error {
	g.world.Update() // Update the game world and systems
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 0, 255})
	g.circleRenderSystem.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 800, 600
}

func main() {
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("v0.0.1-alpha")
	game := NewGame()

	scenarios.StartScenario(game.world)

	game.world.AddSystem(systems.NewGravitySystem(game.world))

	game.world.AddSystem(systems.NewBoundaryBouncingSystem(game.world, 800, 600))
	game.world.AddSystem(systems.NewPositionSystem(game.world))

	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
