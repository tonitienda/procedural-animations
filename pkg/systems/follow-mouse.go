package systems

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tonitienda/procedural-animations/pkg/components"
	"github.com/tonitienda/procedural-animations/pkg/world"
)

type FollowMouse struct {
	positions     map[world.Entity]*components.Position
	leadMovements map[world.Entity]*components.LeadMovement
}

func (f *FollowMouse) Update() {

	mx, my := ebiten.CursorPosition()

	for entity, _ := range f.leadMovements {
		position, ok := f.positions[entity]
		if !ok {
			continue
		}
	}

	return nil
}
