import p5 from "p5";
import { newWorld } from "./world";
import { startBouncingBallsScenario } from "./scenarios/bouncing-circles";
import { startSnakeScenario } from "./scenarios/snake";

const sketch = (p: p5) => {
  const world = newWorld(p);

  //startBouncingBallsScenario(world);
  startSnakeScenario(world);

  world.setup();
  world.draw();
};

new p5(sketch);
