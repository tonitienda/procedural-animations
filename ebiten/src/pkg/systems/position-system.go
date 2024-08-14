package systems

import (
	"github.com/tonitienda/procedural-animations-go/src/pkg/components"
	"github.com/tonitienda/procedural-animations-go/src/pkg/entities"
	"github.com/tonitienda/procedural-animations-go/src/pkg/world"
)

func NewPositionSystem(world *world.World) *PositionSystem {
	return &PositionSystem{
		velocities: world.Velocities,
		positions:  world.Positions,
	}
}

type PositionSystem struct {
	velocities map[entities.Entity]*components.Velocity
	positions  map[entities.Entity]*components.Position
}

func (rs *PositionSystem) Update() {
	for entity, position := range rs.positions {
		if velocity, ok := rs.velocities[entity]; ok {

			// Update position
			position.X += velocity.X
			position.Y += velocity.Y
		}
	}
}
