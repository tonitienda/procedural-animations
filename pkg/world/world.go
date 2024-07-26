package world

import (
	"github.com/tonitienda/procedural-animations/pkg/components"
)

type Entity int

type World struct {
	entities     []Entity
	positions    map[Entity]*components.Position
	sizes        map[Entity]*components.Size
	movements    map[Entity]*components.Movement
	leads        map[Entity]*components.LeadMovement
	systems      []System
	nextEntityID Entity
}

type System interface {
	Update()
}

func NewWorld() *World {
	return &World{
		positions:    make(map[Entity]*components.Position),
		sizes:        make(map[Entity]*components.Size),
		movements:    make(map[Entity]*components.Movement),
		systems:      []System{},
		nextEntityID: 0,
	}
}

func (w *World) AddEntity() Entity {
	id := w.nextEntityID
	w.nextEntityID++
	w.entities = append(w.entities, id)
	return id
}

func (w *World) AddComponents(entity Entity, components ...interface{}) {
	for _, component := range components {
		switch c := component.(type) {
		case *components.Position:
			w.positions[entity] = c
		case *components.Size:
			w.sizes[entity] = c
		case *components.Movement:
			w.movements[entity] = c
		}
	}
}

func (w *World) AddSystem(system System) {
	w.systems = append(w.systems, system)
}

func (w *World) Update() {
	for _, system := range w.systems {
		system.Update()
	}
}
