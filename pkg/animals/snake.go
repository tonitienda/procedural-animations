package animals

import (
	"github.com/tonitienda/procedural-animations/pkg/components"
	"github.com/tonitienda/procedural-animations/pkg/world"
)

func AddSnake(headx, heady int, world *world.World) {
	// 	bodySizes := []float32{20, 25, 30, 25, 20, 15, 15, 15, 15, 15, 15, 15, 15, 15, 10, 10, 10, 5, 5, 3}

	snakeHead := world.AddEntity()
	world.AddComponents(snakeHead, &components.LeadMovement{MaxSpeed: 3}, &components.Position{X: float32(headx), Y: float32(heady)}, &components.Size{Radius: 20})
}
