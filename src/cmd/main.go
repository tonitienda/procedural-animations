package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tonitienda/procedural-animations-go/src/pkg/components"
	"github.com/tonitienda/procedural-animations-go/src/pkg/systems"
	"github.com/tonitienda/procedural-animations-go/src/pkg/world"
)

type Game struct {
	world        *world.World
	renderSystem *systems.RenderSystem
}

func NewGame() *Game {
	w := world.NewWorld()

	rs := systems.NewRenderSystem(w)

	return &Game{
		world:        w,
		renderSystem: rs,
	}
}

func (g *Game) Update() error {
	g.world.Update() // Update the game world and systems
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 0, 255})
	g.renderSystem.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 800, 600
}

func main() {
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("v0.0.1-alpha")
	game := NewGame()

	ball := game.world.AddEntity()
	game.world.AddComponents(ball,
		&components.Position{X: 400, Y: 300},
		&components.Velocity{X: 0, Y: 0},
		&components.Renderable{Radius: 10, Color: color.RGBA{255, 0, 0, 255}},
		&components.GravitationalPull{Acceleration: 0.1})

	game.world.AddSystem(systems.NewGravitySystem(game.world))
	game.world.AddSystem(systems.NewPositionSystem(game.world))

	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
