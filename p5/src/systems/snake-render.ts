type XY = [number, number];

import p5 from "p5";
import { World } from "../world";

function calculatePerpendicularPoints(
  pos: { X: number; Y: number },
  orientation: { Radians: number },
  radius: number
): [XY, XY] {
  const normVX = Math.cos(orientation.Radians);
  const normVY = Math.sin(orientation.Radians);

  const perpX = -normVY;
  const perpY = normVX;

  const point1: XY = [pos.X + perpX * radius, pos.Y + perpY * radius];

  const point2: XY = [pos.X - perpX * radius, pos.Y - perpY * radius];

  return [point1, point2];
}

export const newSnakeRenderSystem = (world: World) => {
  return {
    draw: (p: p5, op: any) => {
      // Calculate the points for the skin
      for (const entity of Object.keys(world.snakes) as any as number[]) {
        const snake = world.snakes[entity];

        const head = snake.Segments[0];
        const headOrientation = world.orientations[head].Radians;
        const second = snake.Segments[1];
        const secondOrientation = world.orientations[second].Radians;

        const rightEye: XY = [
          world.positions[second].X +
            Math.cos(secondOrientation + Math.PI / 3) *
              world.circles[second].Radius,
          world.positions[second].Y +
            Math.sin(secondOrientation + Math.PI / 3) *
              world.circles[second].Radius,
        ];

        const leftEye: XY = [
          world.positions[second].X +
            Math.cos(secondOrientation - Math.PI / 3) *
              world.circles[second].Radius,
          world.positions[second].Y +
            Math.sin(secondOrientation - Math.PI / 3) *
              world.circles[second].Radius,
        ];

        const rightHeadPoint = [
          world.positions[head].X +
            Math.cos(headOrientation + Math.PI / 8) *
              world.circles[head].Radius,
          world.positions[head].Y +
            Math.sin(headOrientation + Math.PI / 8) *
              world.circles[head].Radius,
        ];

        // Start with a point in the head that is almost parallel to the orientation
        // but a little bit to the left side
        const leftHeadPoint = [
          world.positions[head].X +
            Math.cos(headOrientation - Math.PI / 8) *
              world.circles[head].Radius,
          world.positions[head].Y +
            Math.sin(headOrientation - Math.PI / 8) *
              world.circles[head].Radius,
        ];

        const leftSide: XY[] = [];
        const rightSide: XY[] = [];
        for (let segment of snake.Segments) {
          const pos = world.positions[segment];
          const orientation = world.orientations[segment];
          const c = world.circles[segment];
          const [rightPoint, leftPoint] = calculatePerpendicularPoints(
            pos,
            orientation,
            c.Radius
          );

          leftSide.push(leftPoint);
          rightSide.push(rightPoint);
        }

        const snakeContour = [
          leftHeadPoint,
          ...leftSide,
          ...rightSide.reverse(),
          rightHeadPoint,
          leftHeadPoint,
        ];

        // Fill the snake with a greenish color
        p.fill(100, 200, 100);
        p.beginShape();

        for (let i = 0; i < snakeContour.length; i++) {
          const point = snakeContour[i];
          p.curveVertex(point[0], point[1]);
        }

        p.endShape(p.CLOSE);

        // Draw the eyes
        p.fill(100, 100, 220);
        p.circle(rightEye[0], rightEye[1], 10);
        p.circle(leftEye[0], leftEye[1], 10);
      }
    },
  };
};
