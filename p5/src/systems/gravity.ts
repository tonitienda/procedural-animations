/*


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

Translate to typescript
*/

import { World } from "../world";

export const newGravitySystem = (world: World) => {
  return {
    update: () => {
      for (const entity in Object.keys(world.gravitationalPulls)) {
        const gravitationalPull = world.gravitationalPulls[entity];

        if (world.velocities[entity]) {
          world.velocities[entity].Y += gravitationalPull.Acceleration;
        }
      }
    },
  };
};
