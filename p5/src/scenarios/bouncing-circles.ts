import { newBoundaryBouncingSystem } from "../systems/boundary-bouncing";
import { newCircleRenderSystem } from "../systems/circle-render";
import { newGravitySystem } from "../systems/gravity";
import { newPositionSystem } from "../systems/position";
import { World } from "../world";
type BouncingBall = {
  PosX: number;
  PosY: number;
  VelocityX: number;
  VelocityY: number;
  Radius: number;
  Color: number[];
};

/*

func addBouncingBall(world *world.World, settings BouncingBall) {

	ball := world.AddEntity()
	world.AddComponents(ball,
		&components.Position{X: settings.PosX, Y: settings.PosY},
		&components.Velocity{X: settings.VelocityX, Y: settings.VelocityY},
		&components.Circle{Radius: settings.Radius, FillColor: settings.Color},
		&components.GravitationalPull{Acceleration: 0.1},
		&components.BounceBoundaries{BounceFactor: 1})

}
 */

function addBouncingBall(world: World, settings: BouncingBall) {
  const ball = world.addEntity();
  world.addComponents(ball, {
    position: { X: settings.PosX, Y: settings.PosY },
    velocity: { X: settings.VelocityX, Y: settings.VelocityY },
    circle: { Radius: settings.Radius, FillColor: settings.Color },
    gravitationalPull: { Acceleration: 0.1 },
    bounceBoundaries: { BounceFactor: 1 },
  });
}

export function startBouncingBallsScenario(world: World) {
  //world.reset();

  world.addSystem(newGravitySystem(world));
  world.addSystem(newBoundaryBouncingSystem(world, 800, 600));
  world.addSystem(newPositionSystem(world));

  world.addRenderSystem(newCircleRenderSystem(world));

  addBouncingBall(world, {
    PosX: 400,
    PosY: 100,
    VelocityX: 0,
    VelocityY: 0,
    Radius: 30,
    Color: [255, 0, 0, 255],
  });

  addBouncingBall(world, {
    PosX: 0,
    PosY: 0,
    VelocityX: 5,
    VelocityY: 0,
    Radius: 25,
    Color: [0, 255, 0, 255],
  });

  addBouncingBall(world, {
    PosX: 0,
    PosY: 0,
    VelocityX: 2,
    VelocityY: 20,
    Radius: 20,
    Color: [0, 0, 255, 255],
  });
}
