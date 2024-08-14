package systems

import (
	"github.com/tonitienda/procedural-animations-go/src/pkg/components"
	"github.com/tonitienda/procedural-animations-go/src/pkg/entities"
	"github.com/tonitienda/procedural-animations-go/src/pkg/world"
)

func NewGravitySystem(world *world.World) *GravitySystem {
	return &GravitySystem{
		gravitationalPulls: world.GravitationalPulls,
		velocities:         world.Velocities,
	}
}

type GravitySystem struct {
	gravitationalPulls map[entities.Entity]*components.GravitationalPull
	velocities         map[entities.Entity]*components.Velocity
}

func (rs *GravitySystem) Update() {
	for entity, gravitationalPull := range rs.gravitationalPulls {
		if Velocity, ok := rs.velocities[entity]; ok {

			// Apply gravity
			Velocity.Y += gravitationalPull.Acceleration // Consider using a delta time

		}
	}
}
