// import { World } from "../world";

// export const newPositionConstraintSystem = (world: World) => {
//   return {
//     update: () => {
//       for (const entity of Object.keys(
//         world.positionConstraints
//       ) as any as number[]) {
//         const positionConstraint = world.positionConstraints[entity];

//         const anchor = world.positions[positionConstraint.Prev];
//         const anchorOrientation = world.orientations[positionConstraint.Prev];

//         const position = world.positions[entity];
//         const orientation = world.orientations[entity];

//         console.log("positionConstraint", positionConstraint);

//         // Position the entity at the distance and angle from the anchor
//         if (anchor && position && orientation && anchorOrientation) {
//           orientation.Radians =
//             anchorOrientation.Radians + positionConstraint.Radians;

//           // Place the entity at the distance from the anchor
//           position.X =
//             anchor.X +
//             positionConstraint.Distance * Math.cos(orientation.Radians);

//           position.Y =
//             anchor.Y +
//             positionConstraint.Distance * Math.sin(orientation.Radians);
//         }
//       }
//     },
//   };
// };
