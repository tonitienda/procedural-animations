package scenarios

import (
	"image/color"
	"log"

	"github.com/tonitienda/procedural-animations-go/src/pkg/components"
	"github.com/tonitienda/procedural-animations-go/src/pkg/systems"
	"github.com/tonitienda/procedural-animations-go/src/pkg/world"
)

func AddSnake(headx, heady int, world *world.World) {
	bodyPartsRadius := []float64{20, 25, 30, 25, 20, 15, 15, 15, 15, 15, 15, 15, 15, 15, 10, 10, 10, 5, 5, 3}
	log.Println("Adding Snake Head")

	snakeHead := world.AddEntity()
	world.AddComponents(snakeHead,
		&components.LeadMovement{MaxSpeed: 3},
		&components.Position{X: float64(headx), Y: float64(heady)},
		&components.Circle{Radius: float64(bodyPartsRadius[0]), StrokeColor: color.White, ShowCenter: true})

	prev := snakeHead
	yposition := heady
	for idx, radius := range bodyPartsRadius[1:] {
		yposition += int(radius)
		part := world.AddEntity()
		world.AddComponents(part,
			&components.DistanceConstraint{Prev: prev, Distance: float64(radius)},
			&components.Position{X: float64(headx), Y: float64(yposition)},
			&components.Circle{Radius: float64(bodyPartsRadius[idx]), StrokeColor: color.White, ShowCenter: true})

		prev = part

	}

	log.Println("Snake added")

}

func StartSnakeScenario(world *world.World) {
	log.Println("Starting snake scenario")

	world.Reset()

	world.AddSystem(systems.NewFollowMouseSystem(world))
	world.AddSystem(systems.NewDistanceConstraintSystem(world))

	AddSnake(400, 300, world)

}
