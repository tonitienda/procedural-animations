package systems

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/tonitienda/procedural-animations-go/src/pkg/components"
	"github.com/tonitienda/procedural-animations-go/src/pkg/entities"
	"github.com/tonitienda/procedural-animations-go/src/pkg/world"
)

func colorToRGBA(c color.Color) color.RGBA {
	r, g, b, a := c.RGBA()
	return color.RGBA{
		R: uint8(r >> 8),
		G: uint8(g >> 8),
		B: uint8(b >> 8),
		A: uint8(a >> 8),
	}
}

func NewCircleRenderSystem(world *world.World) *CircleRenderSystem {
	return &CircleRenderSystem{
		positions: &world.Positions,
		circles:   &world.Circles,
	}
}

type CircleRenderSystem struct {
	positions *map[entities.Entity]*components.Position
	circles   *map[entities.Entity]*components.Circle
}

func (rs *CircleRenderSystem) Draw(screen *ebiten.Image, op *world.DrawOptions) {
	for entity, position := range *(rs.positions) {
		if circle, ok := (*(rs.circles))[entity]; ok {
			if circle.FillColor != nil {
				vector.DrawFilledCircle(screen, float32(position.X), float32(position.Y), float32(circle.Radius), colorToRGBA(circle.FillColor), true)
			}

			if circle.StrokeColor != nil {
				vector.StrokeCircle(screen, float32(position.X), float32(position.Y), float32(circle.Radius), 1, colorToRGBA(circle.StrokeColor), true)
			}

			if circle.ShowCenter {
				vector.DrawFilledCircle(screen, float32(position.X), float32(position.Y), 2, circle.StrokeColor, true)
			}

		}

	}
}
