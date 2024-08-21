type XY = [number, number];

import p5 from "p5";
import { World } from "../world";

export const newBackgroundRenderSystem = (world: World, p: p5) => {
  const treeTypes = ["summer", "winter", "autumn", "spring", "sakura"];
  // const forestBuffer = p.createGraphics(p.width * 2, p.height); // Buffer larger than canvas to allow scrolling
  const forestBuffer = p.createGraphics(2300, 600); // Buffer larger than canvas to allow scrolling

  let xPosition = 0;
  let aldreadyDrawn = false;
  return {
    draw: (p: p5, op: any) => {
      p.angleMode(p.DEGREES);

      // Fill the cat with a dark-grey color
      p.background(150, 220, 255);

      if (aldreadyDrawn) {
        p.image(forestBuffer, -xPosition, 0);

        xPosition += 5;
        return;
      }

      aldreadyDrawn = true;
      for (let x = xPosition; x < p.width * 2 + xPosition; x++) {
        // This is repeated
        const height = p.noise(x / 500) * 100 + 100;
        const screenX = x - xPosition;
        const surfaceY = p.height - height;

        const shouldDrawTree = p.noise(x + 100000) > 0.5 && x % 50 == 0;
        const treeTypeIdx = Math.round(p.noise(x + 150000) * treeTypes.length);

        if (shouldDrawTree) {
          //p.translate(x - xPosition, p.height - height - 25);

          //p.ellipse(screenX, surfaceY - 25, 50, 50);

          forestBuffer.push();
          forestBuffer.translate(screenX, surfaceY);
          branch(
            forestBuffer,
            50,
            0,
            x,
            screenX,
            surfaceY,
            treeTypes[treeTypeIdx]
          );
          forestBuffer.pop();
        }

        // p.line(x - xPosition, p.height, x - xPosition, p.height - height);
      }

      //drawTerrain(forestBuffer, xPosition);

      xPosition += 5;
    },
  };
};

function drawTerrain(p: p5, xPosition: number) {
  for (let x = 0; x < p.width; x++) {
    const height = p.noise((x + xPosition) / 500) * 100 + 100;

    p.stroke(100, 50, 0);
    p.line(x, p.height, x, p.height - height);
  }
}

function branch(
  p: p5,
  len: number,
  branchNumber: number,
  noisePosition: number,
  x: number,
  y: number,
  treeType: string
) {
  p.push();

  let trunkColor = [0, 0, 0];
  let leafColor = [0, 0, 0];

  if (treeType === "winter") {
    trunkColor = [35, 20, 0];
    leafColor = [
      255 - p.noise(noisePosition + 10000 * branchNumber) * 40,
      255 - p.noise(noisePosition + 10000 * branchNumber) * 10,
      255 - p.noise(noisePosition + 10000 * branchNumber) * 20,
      175,
    ];
  } else if (treeType === "autumn") {
    trunkColor = [70, 40, 0];
    leafColor = [
      80 + p.noise(noisePosition + 10000 * branchNumber) * 100,
      40 + p.noise(noisePosition + 10000 * branchNumber) * 100,
      0,
    ];
  } else if (treeType === "summer") {
    trunkColor = [70, 40, 0];
    leafColor = [
      200 + p.noise(noisePosition + 10000 * branchNumber) * 100,
      120 + p.noise(noisePosition + 10000 * branchNumber) * 100,
      40 + p.noise(noisePosition + 10000 * branchNumber) * 80,
    ];
  } else if (treeType === "spring") {
    trunkColor = [70, 40, 0];
    leafColor = [
      80 + p.noise(noisePosition + 10000 * branchNumber) * 100,
      120 + p.noise(noisePosition + 10000 * branchNumber) * 100,
      40 + p.noise(noisePosition + 10000 * branchNumber) * 100,
    ];
  }
  if (treeType === "sakura") {
    trunkColor = [70, 40, 0];
    leafColor = [
      80 + p.noise(noisePosition + 10000 * branchNumber),
      120 + p.noise(noisePosition + 12000 * branchNumber),
      40 + p.noise(noisePosition + 15000 * branchNumber),
    ];
  }

  // Draw a branch
  if (len > 10) {
    p.strokeWeight(p.map(len, 10, 100, 3, 15));
    p.stroke(trunkColor[0], trunkColor[1], trunkColor[2]);
    p.line(0, 0, 0, -len);
    p.translate(0, -len);

    const children =
      Math.floor(p.noise(noisePosition + 10000 * branchNumber) * 3) + 2;

    p.rotate(-30);

    for (let i = 0; i < children; i++) {
      branch(
        p,
        len *
          Math.min(0.9, p.noise(noisePosition + 10000 * branchNumber) + 0.3),
        branchNumber + 1 + 100 * i,
        noisePosition,
        x,
        y,
        treeType
      );
      p.rotate(60 / (children - 1));
    }
  } else {
    // Draw a leaf

    p.fill(leafColor[0], leafColor[1], leafColor[2], leafColor[3]);
    p.noStroke();
    p.beginShape();
    p.ellipse(0, 0, 25, 7);

    p.endShape(p.CLOSE);
  }
  p.pop();
}
