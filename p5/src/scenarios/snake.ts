/*

package scenarios

import (
	"image/color"
	"log"

	"github.com/tonitienda/procedural-animations-go/src/pkg/components"
	"github.com/tonitienda/procedural-animations-go/src/pkg/entities"
	"github.com/tonitienda/procedural-animations-go/src/pkg/systems"
	"github.com/tonitienda/procedural-animations-go/src/pkg/world"
)

func AddSnake(headx, heady int, world *world.World) {

	bodyPartsRadius := []int{38, 40}

	for i := 31; i > 1; i-- {
		bodyPartsRadius = append(bodyPartsRadius, i)
	}

	log.Println("Adding Snake Head")

	snakeParts := make([]entities.Entity, 0)

	snakeHead := world.AddEntity()
	snakeParts = append(snakeParts, snakeHead)

	world.AddComponents(snakeHead,
		&components.LeadMovement{MaxSpeed: 3},
		&components.Orientation{Radians: 0},
		&components.Position{X: float64(headx), Y: float64(heady)},
		&components.InitialPosition{X: float64(headx), Y: float64(heady)},
		&components.Velocity{X: 1, Y: 1},
		&components.Circle{Radius: float64(bodyPartsRadius[0]), StrokeColor: color.RGBA{255, 255, 255, 64}, ShowCenter: true},
	)

	prev := snakeHead
	yposition := heady
	for idx, radius := range bodyPartsRadius[1:] {
		yposition += int(radius)
		part := world.AddEntity()
		world.AddComponents(part,
			&components.DistanceConstraint{Prev: prev, Distance: 32},
			&components.Orientation{Radians: 0},
			&components.Position{X: float64(headx), Y: float64(yposition)},
			&components.InitialPosition{X: float64(headx), Y: float64(yposition)},
			&components.Velocity{X: 0, Y: 0},
			&components.Circle{Radius: float64(bodyPartsRadius[idx+1]), StrokeColor: color.RGBA{64, 128, 64, 64}, ShowCenter: true},
		)

		snakeParts = append(snakeParts, part)
		prev = part

	}

	snake := world.AddEntity()
	world.AddComponents(snake, &components.Snake{Segments: snakeParts})

	log.Println("Snake added")

}

func StartSnakeScenario(world *world.World) {
	log.Println("Starting snake scenario")

	world.Reset()

	world.AddSystem(systems.NewFollowMouseSystem(world))
	world.AddSystem(systems.NewDistanceConstraintSystem(world))
	world.AddSystem(systems.NewPositionSystem(world))
	world.AddRenderSystem(systems.NewSnakeRenderSystem(world))
	//world.AddRenderSystem(systems.NewCircleRenderSystem(world))

	AddSnake(400, 300, world)

}

convert to ts
*/

import { newDistanceConstraintSystem } from "../systems/distance-constraint";
import { newFollowMouseClickSystem } from "../systems/follow-mouse-click";
import { newPositionSystem } from "../systems/position";
import { newSnakeRenderSystem } from "../systems/snake-render";
import { World } from "../world";

type Snake = {
  Segments: number[];
};

function addSnake(headx: number, heady: number, world: World) {
  const bodyPartsRadius = [38, 40];

  for (let i = 31; i > 1; i--) {
    bodyPartsRadius.push(i);
  }

  console.log("Adding Snake Head");

  const snakeParts: number[] = [];

  const snakeHead = world.addEntity();
  snakeParts.push(snakeHead);

  world.addComponents(snakeHead, {
    leadMovement: { MaxSpeed: 3 },
    orientation: { Radians: 0 },
    position: { X: headx, Y: heady },
    initialPosition: { X: headx, Y: heady },
    velocity: { X: 1, Y: 1 },
    circle: {
      Radius: bodyPartsRadius[0],
      StrokeColor: [255, 255, 255, 64],
      ShowCenter: true,
    },
  });

  let prev = snakeHead;
  let yposition = heady;
  for (let idx = 0; idx < bodyPartsRadius.length - 1; idx++) {
    yposition += bodyPartsRadius[idx];
    const part = world.addEntity();
    world.addComponents(part, {
      distanceConstraint: { Prev: prev, Distance: 32 },
      orientation: { Radians: 0 },
      position: { X: headx, Y: yposition },
      initialPosition: { X: headx, Y: yposition },
      velocity: { X: 0, Y: 0 },
      circle: {
        Radius: bodyPartsRadius[idx + 1],
        StrokeColor: [64, 128, 64, 64],
        ShowCenter: true,
      },
    });

    snakeParts.push(part);
    prev = part;
  }

  const snake = world.addEntity();
  world.addComponents(snake, { snake: { Segments: snakeParts } });

  console.log("Snake added");
}

export function startSnakeScenario(world: World) {
  console.log("Starting snake scenario");

  //world.reset();

  world.addSystem(newFollowMouseClickSystem(world));
  world.addSystem(newDistanceConstraintSystem(world));
  world.addSystem(newPositionSystem(world));
  world.addRenderSystem(newSnakeRenderSystem(world));
  //world.addRenderSystem(newCircleRenderSystem(world));

  addSnake(400, 300, world);
}
