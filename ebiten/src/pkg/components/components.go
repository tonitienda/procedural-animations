package components

import (
	"image/color"

	"github.com/tonitienda/procedural-animations-go/src/pkg/entities"
)

type Position struct {
	X, Y float64
}

type InitialPosition struct {
	X, Y float64
}

type Velocity struct {
	X, Y float64
}

type Circle struct {
	Radius      float64
	FillColor   color.Color
	StrokeColor color.Color
	ShowCenter  bool
}

type GravitationalPull struct {
	Acceleration float64
}

type BounceBoundaries struct {
	BounceFactor float64
}

type LeadMovement struct {
	MaxSpeed float32
}

type DistanceConstraint struct {
	Prev     entities.Entity
	Distance float64
}

type ChainLink struct {
	Prev entities.Entity
	Next entities.Entity
}

type Snake struct {
	Segments []entities.Entity
}

type Orientation struct {
	Radians float64
}
