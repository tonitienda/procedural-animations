package systems

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/tonitienda/procedural-animations/pkg/components"
	"github.com/tonitienda/procedural-animations/pkg/entities"
)

type RenderSystem struct {
	Positions map[entities.Entity]*components.Position
	Sizes     map[entities.Entity]*components.Size
}

func (rs *RenderSystem) Draw(screen *ebiten.Image) {
	circleColor := color.RGBA{255, 255, 255, 255} // Red color with full opacity

	for entity, position := range rs.Positions {
		if size, ok := rs.Sizes[entity]; ok {
			// Render the entity
			// This example assumes you're drawing circles for entities

			vector.StrokeCircle(screen, float32(position.X), float32(position.Y), float32(size.Radius), float32(1), circleColor, true)
			vector.DrawFilledCircle(screen, float32(position.X), float32(position.Y), float32(1), circleColor, true)
		}
	}
}
