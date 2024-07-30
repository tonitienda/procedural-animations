package systems

import (
	"github.com/tonitienda/procedural-animations-go/src/pkg/components"
	"github.com/tonitienda/procedural-animations-go/src/pkg/entities"
	"github.com/tonitienda/procedural-animations-go/src/pkg/world"
)

func NewBoundaryBouncingSystem(world *world.World, width int, height int) *BoundaryBouncingSystem {
	return &BoundaryBouncingSystem{
		velocities:        world.Velocities,
		positions:         world.Positions,
		boundaryBouncings: world.BoundaryBouncings,
		width:             width,
		height:            height,
	}
}

type BoundaryBouncingSystem struct {
	velocities        map[entities.Entity]*components.Velocity
	positions         map[entities.Entity]*components.Position
	boundaryBouncings map[entities.Entity]*components.BounceBoundaries
	width             int
	height            int
}

func (rs *BoundaryBouncingSystem) Update() {
	for entity, bouncing := range rs.boundaryBouncings {
		if position, ok := rs.positions[entity]; ok {
			if velocity, ok := rs.velocities[entity]; ok {

				if position.X < 0 || position.X > float64(rs.width) {
					velocity.X = velocity.X * -bouncing.BounceFactor
				}

				if position.Y < 0 || position.Y > float64(rs.height) {
					velocity.Y = velocity.Y * -bouncing.BounceFactor
				}
			}
		}
	}
}
