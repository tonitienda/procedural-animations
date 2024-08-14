import { World } from "../world";
import p5 from "p5";

export const newCircleRenderSystem = (world: World) => {
  return {
    draw: (p: p5, op: any) => {
      for (const entity in Object.keys(world.circles)) {
        const position = world.positions[entity];
        const circle = world.circles[entity];

        if (circle) {
          if (circle.FillColor) {
            p.fill(
              circle.FillColor[0],
              circle.FillColor[1],
              circle.FillColor[2]
            );
            p.ellipse(
              position.X,
              position.Y,
              circle.Radius * 2,
              circle.Radius * 2
            );
          }

          if (circle.StrokeColor) {
            p.stroke(
              circle.StrokeColor[0],
              circle.StrokeColor[1],
              circle.StrokeColor[2]
            );
            p.strokeWeight(1);
            p.noFill();
            p.ellipse(
              position.X,
              position.Y,
              circle.Radius * 2,
              circle.Radius * 2
            );
          }

          if (circle.ShowCenter) {
            p.fill(255, 0, 0);
            p.ellipse(position.X, position.Y, 2, 2);
          }
        }
      }
    },
  };
};
