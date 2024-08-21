type XY = [number, number];

import p5 from "p5";
import { World } from "../world";

export const newCatRenderSystem = (world: World) => {
  return {
    draw: (p: p5, op: any) => {
      // Fill the cat with a dark-grey color
      p.fill(50, 50, 75);
      p.stroke(255, 255, 255);
      p.strokeWeight(5);
      p.beginShape();
      p.curveTightness(0.1);
      p.curveVertex(840, 0);
      p.curveVertex(880, 35);
      p.curveVertex(835, 140);
      p.curveVertex(845, 175);
      p.endShape();
      p.beginShape();
      p.curveTightness(0.1);
      p.curveVertex(840, 0);
      p.curveVertex(880, 35);
      p.curveVertex(795, 140);
      p.curveVertex(875, 175);
      p.endShape();

      // Tail
      //   p.curveVertex(700, 150);
      //   p.curveVertex(400, 50);
      //   p.curveVertex(250, 50);
      //   p.curveVertex(150, 60);
      //   p.curveVertex(90, 130);
      //   p.curveVertex(100, 140);
      //   p.curveVertex(110, 130);
      //   p.curveVertex(150, 110);

      //   p.curveVertex(600, 150);

      //   p.curveVertex(500, 250);
      //   p.curveVertex(450, 280);
      //   p.curveVertex(400, 250);
      //   p.curveVertex(350, 250);
      //   p.curveVertex(250, 280);
      //   p.curveVertex(150, 310);
      //   p.curveVertex(145, 315);
      //   p.curveVertex(150, 320);
      //   p.curveVertex(250, 290);
      //   p.curveVertex(350, 265);
      //   p.curveVertex(400, 355);
      //   p.curveVertex(450, 390);
      //   p.curveVertex(430, 320);
      //   p.curveVertex(490, 330);
      //   p.curveVertex(550, 330);
      //   p.curveVertex(550, 390);
      //   p.curveVertex(570, 390);
      //   p.curveVertex(570, 330);
      //   p.curveVertex(600, 280);
      //   p.curveVertex(650, 260);
      //   p.curveVertex(660, 230);
      //   p.curveVertex(660, 200);
      //   p.curveVertex(600, 190);
    },
  };
};
