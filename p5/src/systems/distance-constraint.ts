import { World } from "../world";

export const newDistanceConstraintSystem = (world: World) => {
  return {
    update: () => {
      console.log("world.distanceConstraints", world.distanceConstraints);
      for (const entity of Object.keys(
        world.distanceConstraints
      ) as any as number[]) {
        console.log("entity", entity);
        const distanceConstraint = world.distanceConstraints[entity];
        console.log("distanceConstraint", distanceConstraint);

        const anchor = world.positions[distanceConstraint.Prev];
        const position = world.positions[entity];

        if (anchor && position) {
          // Normalize the vector that connect point with anchor

          // Calculate the vector that connects the anchor with the point
          const diifx = position.X - anchor.X;
          const diffy = position.Y - anchor.Y;

          // Normalize the vector by dividing it by its length
          const dotProduct = diifx * diifx + diffy * diffy;

          let normalizedX = 0.0;
          let normalizedY = 0.0;

          if (dotProduct !== 0) {
            normalizedX = diifx / Math.sqrt(dotProduct);
            normalizedY = diffy / Math.sqrt(dotProduct);
          }

          if (world.velocities[entity]) {
            world.velocities[entity].X =
              normalizedX * distanceConstraint.Distance + anchor.X - position.X;
            world.velocities[entity].Y =
              normalizedY * distanceConstraint.Distance + anchor.Y - position.Y;
          }
          if (world.orientations[entity]) {
            world.orientations[entity].Radians = Math.atan2(
              anchor.Y - position.Y,
              anchor.X - position.X
            );
          }
        }
      }
    },
  };
};
