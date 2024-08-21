import p5 from "p5";
import { newWorld } from "./world";
import { startBouncingBallsScenario } from "./scenarios/bouncing-circles";
import { startSnakeScenario } from "./scenarios/snake";
import { startCatScenario } from "./scenarios/cat";

//export function startSketch(scenario: string) {
//  alert("startSketch");
const BOUNCONG_BALLS = "bouncing-balls";
const SNAKE = "snake";
const CAT = "cat";

const scenario = CAT;
const sketch = (p: p5) => {
  const world = newWorld(p);

  world.setup();

  switch (scenario) {
    case BOUNCONG_BALLS:
      startBouncingBallsScenario(world);
      break;
    case SNAKE:
      startSnakeScenario(world);
      break;

    case CAT:
      startCatScenario(world, p);
      break;
    default:
      startBouncingBallsScenario(world);
  }

  world.draw();
};

new p5(sketch);
//}
