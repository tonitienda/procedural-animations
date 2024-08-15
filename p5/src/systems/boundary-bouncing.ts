import { World } from "../world";

export const newBoundaryBouncingSystem = (
  world: World,
  width: number,
  height: number
) => {
  return {
    update: () => {
      for (const entity of Object.keys(
        world.bounceBoundaries
      ) as any as number[]) {
        const bouncing = world.bounceBoundaries[entity];
        const position = world.positions[entity];
        const velocity = world.velocities[entity];

        if (position && velocity) {
          if (position.X < 0 || position.X > width) {
            velocity.X = velocity.X * -bouncing.BounceFactor;
          }

          if (position.Y < 0 || position.Y > height) {
            velocity.Y = velocity.Y * -bouncing.BounceFactor;
          }
        }
      }
    },
  };
};
