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
	// Move the head
	// FIXME - Current functionality calculates movement separately from x and y axis
	// Improve it to calculate the movement vector and normalize it
	// fmt.Println("Mouse: ", mx, my)
	// fmt.Println("Position: ", position.X, position.Y)

	// fmt.Println("Max Speed: ", maxSpeed)
	diffX, diffY := mx-position.X, my-position.Y

	// fmt.Println("Difference: ", diffX, diffY)

	stepX := float64(0.0)
	stepY := float64(0.0)

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

	// fmt.Println("Step: ", stepY, stepX)
	// fmt.Println("-------------------")

	return &components.Velocity{
		X: stepX,
		Y: stepY,
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
