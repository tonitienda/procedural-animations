package systems

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/tonitienda/procedural-animations-go/src/pkg/components"
	"github.com/tonitienda/procedural-animations-go/src/pkg/entities"
	"github.com/tonitienda/procedural-animations-go/src/pkg/world"
)

func NewRenderSystem(world *world.World) *RenderSystem {
	return &RenderSystem{
		positions:   &world.Positions,
		renderables: &world.Renderables,
	}
}

type RenderSystem struct {
	positions   *map[entities.Entity]*components.Position
	renderables *map[entities.Entity]*components.Renderable
}

func (rs *RenderSystem) Draw(screen *ebiten.Image) {
	for entity, position := range *(rs.positions) {
		if renderable, ok := (*(rs.renderables))[entity]; ok {

			// Render the entity
			// This example assumes you're drawing circles for entities
			ebitenutil.DrawCircle(screen, position.X, position.Y, renderable.Radius, renderable.Color)
		}

	}
}
