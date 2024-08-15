/*

package systems

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tonitienda/procedural-animations-go/src/pkg/components"
	"github.com/tonitienda/procedural-animations-go/src/pkg/entities"
	"github.com/tonitienda/procedural-animations-go/src/pkg/world"
)

func NewFollowMouseSystem(world *world.World) *FollowMouse {
	return &FollowMouse{
		Positions:     world.Positions,
		Velocities:    world.Velocities,
		LeadMovements: world.LeadMovements,
		Orientations:  world.Orientations,
	}
}

type FollowMouse struct {
	Positions     map[entities.Entity]*components.Position
	LeadMovements map[entities.Entity]*components.LeadMovement
	Velocities    map[entities.Entity]*components.Velocity
	Orientations  map[entities.Entity]*components.Orientation
}

func getNewVelocity(mx, my float64, maxSpeed float64, position *components.Position) *components.Velocity {
	// Move the head towards the mouse setting its velocity
	// The head should move towards the mouse at a speed of 3 pixels per frame

	// Calculate the difference between the mouse and the head
	diffX, diffY := mx-position.X, my-position.Y

	// Calculate the step that the head should take
	stepX := float64(0.0)
	stepY := float64(0.0)

	// The total speed should be taken into account as the whole vector.
	// X and Y components should be normalized and multiplied by the max speed
	// to get the step that the head should take
	if diffX > 0 {
		stepX = math.Min(maxSpeed, float64(diffX))
	} else {
		stepX = math.Max(-maxSpeed, float64(diffX))
	}

	if diffY > 0 {
		stepY = math.Min(maxSpeed, float64(diffY))
	} else {
		stepY = math.Max(-maxSpeed, float64(diffY))
	}

	// Normalize the vector
	dotProduct := float64(stepX*stepX + stepY*stepY)

	normalizedX, normalizedY := float64(0.0), float64(0.0)

	if dotProduct != 0 {
		normalizedX = stepX / float64(math.Sqrt(float64(dotProduct)))
		normalizedY = stepY / float64(math.Sqrt(float64(dotProduct)))
	}

	return &components.Velocity{
		X: normalizedX * maxSpeed,
		Y: normalizedY * maxSpeed,
	}

}

func (f *FollowMouse) Update() {

	mx, my := ebiten.CursorPosition()

	for entity, leadMovement := range f.LeadMovements {

		position, ok := f.Positions[entity]
		if !ok {
			continue
		}

		_, ok = f.Velocities[entity]
		if ok {
			f.Velocities[entity] = getNewVelocity(float64(mx), float64(my), float64(leadMovement.MaxSpeed), position)

		}

		_, ok = f.Orientations[entity]

		if ok && (position.X != float64(mx) || position.Y != float64(my)) {
			f.Orientations[entity] = &components.Orientation{
				Radians: math.Atan2(float64(my)-position.Y, float64(mx)-position.X),
			}
		}

	}
}


convert to ts, p5 and change follow mouse position to click position

*/

import { World } from "../world";

function getNewVelocity(
  mx: number,
  my: number,
  maxSpeed: number,
  position: { X: number; Y: number }
) {
  const diffX = mx - position.X;
  const diffY = my - position.Y;

  let stepX = 0.0;
  let stepY = 0.0;

  if (diffX > 0) {
    stepX = Math.min(maxSpeed, diffX);
  } else {
    stepX = Math.max(-maxSpeed, diffX);
  }

  if (diffY > 0) {
    stepY = Math.min(maxSpeed, diffY);
  } else {
    stepY = Math.max(-maxSpeed, diffY);
  }

  const dotProduct = stepX * stepX + stepY * stepY;

  let normalizedX = 0.0;
  let normalizedY = 0.0;

  if (dotProduct !== 0) {
    normalizedX = stepX / Math.sqrt(dotProduct);
    normalizedY = stepY / Math.sqrt(dotProduct);
  }

  return {
    X: normalizedX * maxSpeed,
    Y: normalizedY * maxSpeed,
  };
}

export const newFollowMouseClickSystem = (world: World) => {
  return {
    update: () => {
      const mx = world.p.mouseX;
      const my = world.p.mouseY;

      for (const entity of Object.keys(
        world.leadMovements
      ) as any as number[]) {
        const lead = world.leadMovements[entity];
        const position = world.positions[entity];

        if (position) {
          if (world.velocities[entity]) {
            world.velocities[entity] = getNewVelocity(
              mx,
              my,
              lead.MaxSpeed,
              position
            );
          }

          if (world.orientations[entity]) {
            world.orientations[entity].Radians = Math.atan2(
              my - position.Y,
              mx - position.X
            );
          }
        }
      }
    },
  };
};
