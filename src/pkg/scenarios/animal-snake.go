package scenarios

import (
	"image/color"
	"log"

	"github.com/tonitienda/procedural-animations-go/src/pkg/components"
	"github.com/tonitienda/procedural-animations-go/src/pkg/systems"
	"github.com/tonitienda/procedural-animations-go/src/pkg/world"
)

func AddSnake(headx, heady int, world *world.World) {
	// 	bodySizes := []float32{20, 25, 30, 25, 20, 15, 15, 15, 15, 15, 15, 15, 15, 15, 10, 10, 10, 5, 5, 3}
	log.Println("Adding Snake Head")

	snakeHead := world.AddEntity()
	world.AddComponents(snakeHead, &components.LeadMovement{MaxSpeed: 3}, &components.Position{X: float64(headx), Y: float64(heady)}, &components.Circle{Radius: 20, Color: color.White})

	log.Println("Snake added")

}

func StartSnakeScenario(world *world.World) {
	log.Println("Starting snake scenario")

	world.Reset()

	world.AddSystem(systems.NewFollowMouseSystem(world))

	AddSnake(400, 300, world)

}
