package world

import (
	"github.com/tonitienda/procedural-animations/pkg/components"
	"github.com/tonitienda/procedural-animations/pkg/entities"
)

type World struct {
	entities      []entities.Entity
	Positions     map[entities.Entity]*components.Position
	Sizes         map[entities.Entity]*components.Size
	LeadMovements map[entities.Entity]*components.LeadMovement
	Systems       []System
	nextEntityID  entities.Entity
}

type System interface {
	Update()
}

func NewWorld() *World {
	return &World{
		Positions:     make(map[entities.Entity]*components.Position),
		Sizes:         make(map[entities.Entity]*components.Size),
		LeadMovements: make(map[entities.Entity]*components.LeadMovement),
		Systems:       []System{},
		nextEntityID:  0,
	}
}

func (w *World) AddEntity() entities.Entity {
	id := w.nextEntityID
	w.nextEntityID++
	w.entities = append(w.entities, id)
	return id
}

func (w *World) AddComponents(entity entities.Entity, args ...interface{}) {
	for _, component := range args {
		switch c := component.(type) {
		case *components.Position:
			w.Positions[entity] = c
		case *components.Size:
			w.Sizes[entity] = c
		case *components.LeadMovement:
			w.LeadMovements[entity] = c
		}
	}
}

func (w *World) AddSystem(system System) {
	w.Systems = append(w.Systems, system)
}

func (w *World) Update() {
	for _, system := range w.Systems {
		system.Update()
	}
}
