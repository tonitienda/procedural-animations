import p5 from "p5";
import { newWorld } from "./world";
import { startBouncingBallsScenario } from "./scenarios/bouncing-circles";

const sketch = (p: p5) => {
  const world = newWorld(p);

  startBouncingBallsScenario(world);

  world.setup();
  world.draw();
};

new p5(sketch);
