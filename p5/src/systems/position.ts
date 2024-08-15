import { World } from "../world";

export const newPositionSystem = (world: World) => {
  return {
    update: () => {
      for (const entity of Object.keys(world.positions) as any as number[]) {
        const velocity = world.velocities[entity];

        if (world.positions[entity] && velocity) {
          world.positions[entity].X += velocity.X;
          world.positions[entity].Y += velocity.Y;
        }
      }
    },
  };
};
