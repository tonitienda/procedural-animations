export type Position = {
  X: number;
  Y: number;
};

export type InitialPosition = {
  X: number;
  Y: number;
};

export type Velocity = {
  X: number;
  Y: number;
};

export type Circle = {
  Radius: number;
  FillColor?: number[];
  StrokeColor?: number[];
  ShowCenter?: boolean;
};

export type GravitationalPull = {
  Acceleration: number;
};

export type BounceBoundaries = {
  BounceFactor: number;
};

export type LeadMovement = {
  MaxSpeed: number;
};

export type DistanceConstraint = {
  Prev: number;
  Distance: number;
};

export type PositionConstraint = {
  Prev: number;
  Distance: number;
  Radians: number;
};

export type ChainLink = {
  Prev: number;
  Next: number;
};

export type Snake = {
  Segments: number[];
};

export type Orientation = {
  Radians: number;
};

export type BodyPart = {
  Position: Position;
};

export type Cat = {
  Head: BodyPart;
};
