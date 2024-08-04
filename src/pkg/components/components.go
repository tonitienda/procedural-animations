package components

import "image/color"

type Position struct {
	X, Y float64
}

type Velocity struct {
	X, Y float64
}

type Circle struct {
	Radius float64
	Color  color.Color
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

type ChainMovement struct {
	Prev int
	Next int
}
