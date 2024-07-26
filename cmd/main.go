package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tonitienda/procedural-animations/pkg/animals"
	"github.com/tonitienda/procedural-animations/pkg/systems"
	"github.com/tonitienda/procedural-animations/pkg/world"
)

type Game struct {
	World        *world.World
	renderSystem *systems.RenderSystem
}

func (g *Game) Update() error {
	g.World.Update()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.renderSystem.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 800, 600
}

func main() {
	w := world.NewWorld()
	animals.AddSnake(400, 300, w)

	w.AddSystem(&systems.FollowMouse{
		Positions:     w.Positions,
		LeadMovements: w.LeadMovements,
	})

	rs := &systems.RenderSystem{
		Positions: w.Positions,
		Sizes:     w.Sizes,
	}

	game := &Game{
		World:        w,
		renderSystem: rs,
	}
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("Procedural animations")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}

}
