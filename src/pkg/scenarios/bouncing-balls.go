package scenarios

import (
	"image/color"

	"github.com/tonitienda/procedural-animations-go/src/pkg/components"
	"github.com/tonitienda/procedural-animations-go/src/pkg/systems"
	"github.com/tonitienda/procedural-animations-go/src/pkg/world"
)

type BouncingBall struct {
	PosX      float64
	PosY      float64
	VelocityX float64
	VelocityY float64
	Radius    float64
	Color     color.RGBA
}

func addBouncingBall(world *world.World, settings BouncingBall) {

	ball := world.AddEntity()
	world.AddComponents(ball,
		&components.Position{X: settings.PosX, Y: settings.PosY},
		&components.Velocity{X: settings.VelocityX, Y: settings.VelocityY},
		&components.Circle{Radius: settings.Radius, FillColor: settings.Color},
		&components.GravitationalPull{Acceleration: 0.1},
		&components.BounceBoundaries{BounceFactor: 1})

}

func StartBouncingBallsScenario(world *world.World) {

	world.Reset()

	world.AddSystem(systems.NewGravitySystem(world))
	world.AddSystem(systems.NewBoundaryBouncingSystem(world, 800, 600))
	world.AddSystem(systems.NewPositionSystem(world))

	addBouncingBall(world, BouncingBall{
		PosX:      400,
		PosY:      100,
		VelocityX: 0,
		VelocityY: 0,
		Radius:    30, Color: color.RGBA{255, 0, 0, 255},
	})

	addBouncingBall(world, BouncingBall{
		PosX:      0,
		PosY:      0,
		VelocityX: 5,
		VelocityY: 0,
		Radius:    25, Color: color.RGBA{0, 255, 0, 255},
	})

	addBouncingBall(world, BouncingBall{
		PosX:      0,
		PosY:      0,
		VelocityX: 2,
		VelocityY: 20,
		Radius:    20, Color: color.RGBA{0, 0, 255, 255},
	})
}
