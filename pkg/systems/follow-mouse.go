package systems

import (
	"fmt"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tonitienda/procedural-animations/pkg/components"
	"github.com/tonitienda/procedural-animations/pkg/entities"
)

type FollowMouse struct {
	Positions     map[entities.Entity]*components.Position
	LeadMovements map[entities.Entity]*components.LeadMovement
}

func getNewPosition(mx, my float32, maxSpeed float64, position *components.Position) *components.Position {
	// Move the head
	// FIXME - Current functionality calculates movement separately from x and y axis
	// Improve it to calculate the movement vector and normalize it
	fmt.Println("Mouse: ", mx, my)
	fmt.Println("Position: ", position.X, position.Y)

	fmt.Println("Max Speed: ", maxSpeed)
	diffX, diffY := float32(mx)-position.X, float32(my)-position.Y

	fmt.Println("Difference: ", diffX, diffY)

	stepX := float32(0.0)
	stepY := float32(0.0)

	if diffX > 0 {
		stepX = float32(math.Min(maxSpeed, float64(diffX)))
	} else {
		stepX = float32(math.Max(-maxSpeed, float64(diffX)))
	}

	if diffY > 0 {
		stepY = float32(math.Min(maxSpeed, float64(diffY)))
	} else {
		stepY = float32(math.Max(-maxSpeed, float64(diffY)))
	}

	fmt.Println("Step: ", stepY, stepX)
	fmt.Println("-------------------")

	return &components.Position{
		X: position.X + stepX,
		Y: position.Y + stepY,
	}

}

func (f *FollowMouse) Update() {

	mx, my := ebiten.CursorPosition()

	for entity, leadMovement := range f.LeadMovements {
		position, ok := f.Positions[entity]
		if !ok {
			continue
		}

		f.Positions[entity] = getNewPosition(float32(mx), float32(my), float64(leadMovement.MaxSpeed), position)
	}
}
