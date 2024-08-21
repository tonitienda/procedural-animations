import p5 from "p5";
import { newBackgroundRenderSystem } from "../systems/background-render-system";
import { World } from "../world";

export function startCatScenario(world: World, p: p5) {
  world.addRenderSystem(newBackgroundRenderSystem(world, p));
  world.addRenderSystem(world);
}
