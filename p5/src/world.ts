import p5 from "p5";
import {
  Position,
  Velocity,
  Circle,
  GravitationalPull,
  BounceBoundaries,
  LeadMovement,
  DistanceConstraint,
  ChainLink,
  Snake,
  Orientation,
} from "./components/index";
import { i, s } from "../node_modules/vite/dist/node/types.d-aGj9QkWt";

interface System {
  update: () => void;
}

interface RenderSystem {
  draw: (p: p5, op: any) => void;
}

// Move to specific types
type ComponentsList = {
  initialPosition?: Position;
  position?: Position;
  velocity?: Velocity;
  circle?: Circle;
  gravitationalPull?: GravitationalPull;
  bounceBoundaries?: BounceBoundaries;
  leadMovement?: LeadMovement;
  distanceConstraint?: DistanceConstraint;
  chainLink?: ChainLink;
  snake?: Snake;
  orientation?: Orientation;
};

export interface World {
  p: p5;
  setup: () => void;
  draw: () => void;
  update: () => void;
  addEntity: () => number;
  addComponents: (entity: number, components: ComponentsList) => void;
  reset: () => void;
  addSystem: (system: System) => void;
  addRenderSystem: (system: RenderSystem) => void;

  // Components
  positions: { [key: number]: Position };
  initialPositions: { [key: number]: Position };
  velocities: { [key: number]: Velocity };
  circles: { [key: number]: Circle };
  gravitationalPulls: { [key: number]: GravitationalPull };
  bounceBoundaries: { [key: number]: BounceBoundaries };
  leadMovements: { [key: number]: LeadMovement };
  distanceConstraints: { [key: number]: DistanceConstraint };
  chainLinks: { [key: number]: ChainLink };
  snakes: { [key: number]: Snake };
  orientations: { [key: number]: Orientation };
}

export const newWorld = (p: p5): World => {
  const systems: System[] = [];
  const renderSystems: RenderSystem[] = [];
  let entityId = 0;
  let positions: { [key: number]: Position } = {};
  let initialPositions: { [key: number]: Position } = {};
  let velocities: { [key: number]: Velocity } = {};
  let circles: { [key: number]: Circle } = {};
  let gravitationalPulls: { [key: number]: GravitationalPull } = {};
  let bounceBoundaries: { [key: number]: BounceBoundaries } = {};
  let leadMovements: { [key: number]: LeadMovement } = {};
  let distanceConstraints: { [key: number]: DistanceConstraint } = {};
  let chainLinks: { [key: number]: ChainLink } = {};
  let snakes: { [key: number]: Snake } = {};
  let orientations: { [key: number]: Orientation } = {};

  const update = () => {
    for (const system of systems) {
      system.update();
    }
  };

  return {
    p,
    positions,
    initialPositions,
    velocities,
    circles,
    gravitationalPulls,
    bounceBoundaries,
    leadMovements,
    distanceConstraints,
    chainLinks,
    snakes,
    orientations,
    setup: () => {
      p.setup = () => {
        p.createCanvas(1600, 1200);
      };
    },
    reset: () => {
      entityId = 0;
      positions = {};
      initialPositions = {};
      velocities = {};
      circles = {};
      gravitationalPulls = {};
      bounceBoundaries = {};
      leadMovements = {};
      distanceConstraints = {};
      chainLinks = {};
      snakes = {};
      orientations = {};
    },
    draw: () => {
      p.draw = () => {
        update();
        p.background(0);
        for (const system of renderSystems) {
          system.draw(p, {});
        }
      };
    },
    update,
    addEntity: () => {
      return entityId++;
    },

    addSystem: (system: System) => {
      systems.push(system);
    },
    addRenderSystem: (system: RenderSystem) => {
      renderSystems.push(system);
    },
    addComponents: (entity: number, components: ComponentsList) => {
      if (components.position) {
        positions[entity] = components.position;
      }
      if (components.velocity) {
        velocities[entity] = components.velocity;
      }
      if (components.circle) {
        circles[entity] = components.circle;
      }
      if (components.gravitationalPull) {
        gravitationalPulls[entity] = components.gravitationalPull;
      }
      if (components.bounceBoundaries) {
        bounceBoundaries[entity] = components.bounceBoundaries;
      }
      if (components.leadMovement) {
        leadMovements[entity] = components.leadMovement;
      }
      if (components.distanceConstraint) {
        distanceConstraints[entity] = components.distanceConstraint;
      }
      if (components.chainLink) {
        chainLinks[entity] = components.chainLink;
      }
      if (components.snake) {
        snakes[entity] = components.snake;
      }
      if (components.orientation) {
        orientations[entity] = components.orientation;
      }
      if (components.initialPosition) {
        initialPositions[entity] = components.initialPosition;
      }
    },
  };
};
