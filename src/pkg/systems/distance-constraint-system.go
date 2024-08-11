package systems

import (
	"math"

	"github.com/tonitienda/procedural-animations-go/src/pkg/components"
	"github.com/tonitienda/procedural-animations-go/src/pkg/entities"
	"github.com/tonitienda/procedural-animations-go/src/pkg/world"
)

func NewDistanceConstraintSystem(world *world.World) *DistanceConstraingSystem {
	return &DistanceConstraingSystem{
		distanceConstraints: world.DistanceConstraints,
		positions:           world.Positions,
		velocities:          world.Velocities,
		orientations:        world.Orientations,
	}
}

type DistanceConstraingSystem struct {
	distanceConstraints map[entities.Entity]*components.DistanceConstraint
	positions           map[entities.Entity]*components.Position
	velocities          map[entities.Entity]*components.Velocity
	orientations        map[entities.Entity]*components.Orientation
}

func (rs *DistanceConstraingSystem) Update() {
	for entity, distanceConstraint := range rs.distanceConstraints {
		if anchor, ok := rs.positions[distanceConstraint.Prev]; ok {
			if position, ok := rs.positions[entity]; ok {
				// Normalize the vector that connect point with anchor

				// Calculate the vector that connects the anchor with the point
				diifx, diffy := float64(position.X-anchor.X), float64(position.Y-anchor.Y)

				// Normalize the vector by dividing it by its length
				dotProduct := float64(diifx*diifx + diffy*diffy)

				normalizedX, normalizedY := float64(0.0), float64(0.0)

				if dotProduct != 0 {
					normalizedX = diifx / float64(math.Sqrt(float64(dotProduct)))
					normalizedY = diffy / float64(math.Sqrt(float64(dotProduct)))
				}

				if velocity, ok := rs.velocities[entity]; ok {
					velocity.X = normalizedX*distanceConstraint.Distance + anchor.X - position.X
					velocity.Y = normalizedY*distanceConstraint.Distance + anchor.Y - position.Y
				}
				if orientation, ok := rs.orientations[entity]; ok {
					orientation.Radians = math.Atan2(normalizedY, normalizedX)
				}
			}
		}
	}
}
