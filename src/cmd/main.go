package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/tonitienda/procedural-animations-go/src/pkg/scenarios"
	"github.com/tonitienda/procedural-animations-go/src/pkg/systems"
	"github.com/tonitienda/procedural-animations-go/src/pkg/world"
)

type Scenario func(w *world.World)

var scenarioList = map[string]Scenario{
	"Bouncing cicles":  scenarios.StartBouncingBallsScenario,
	"Procedural Snake": scenarios.StartSnakeScenario,
}

var menuItems = []string{"Bouncing cicles", "Procedural Snake"}

type Game struct {
	world              *world.World
	circleRenderSystem *systems.CircleRenderSystem
	inMenu             bool
}

func NewGame() *Game {
	w := world.NewWorld()

	rs := systems.NewCircleRenderSystem(w)

	return &Game{
		world:              w,
		circleRenderSystem: rs,
		inMenu:             true,
	}
}

func (g *Game) Update() error {
	if g.inMenu {
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			x, y := ebiten.CursorPosition()
			for i, item := range menuItems {
				if x >= 100 && x <= 300 && y >= 50+i*30 && y <= 80+i*30 {
					if scenario, ok := scenarioList[item]; ok {
						scenario(g.world)
					}
					g.inMenu = false
				}
			}
		}
	} else {
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			x, y := ebiten.CursorPosition()

			if x >= 700 && x <= 800 && y >= 50 && y <= 80 {
				g.inMenu = true
			}
		}

		g.world.Update() // Update the game world and systems
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.inMenu {
		g.drawMenu(screen)
	} else {
		screen.Fill(color.RGBA{0, 0, 0, 255})

		ebitenutil.DebugPrintAt(screen, "Menu", 700, 50)
		g.world.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 800, 600
}

func (g *Game) drawMenu(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 0, 255})
	for i, item := range menuItems {
		ebitenutil.DebugPrintAt(screen, item, 100, 50+i*30)
	}
}

func main() {
	ebiten.SetWindowSize(1200, 800)
	ebiten.SetWindowTitle("v0.0.1-alpha")
	game := NewGame()

	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
