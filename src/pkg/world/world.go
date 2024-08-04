package world

import (
	"github.com/tonitienda/procedural-animations-go/src/pkg/components"
	"github.com/tonitienda/procedural-animations-go/src/pkg/entities"
)

// FIXME - entities map should be private or readonly
type World struct {
	entities            []entities.Entity
	Positions           map[entities.Entity]*components.Position
	Velocities          map[entities.Entity]*components.Velocity
	GravitationalPulls  map[entities.Entity]*components.GravitationalPull
	Circles             map[entities.Entity]*components.Circle
	BoundaryBouncings   map[entities.Entity]*components.BounceBoundaries
	LeadMovements       map[entities.Entity]*components.LeadMovement
	DistanceConstraints map[entities.Entity]*components.DistanceConstraint
	systems             []System
	nextEntityID        entities.Entity
}

type System interface {
	Update()
}

func NewWorld() *World {
	return &World{
		Positions:           make(map[entities.Entity]*components.Position),
		Velocities:          make(map[entities.Entity]*components.Velocity),
		GravitationalPulls:  make(map[entities.Entity]*components.GravitationalPull),
		Circles:             make(map[entities.Entity]*components.Circle),
		BoundaryBouncings:   make(map[entities.Entity]*components.BounceBoundaries),
		LeadMovements:       make(map[entities.Entity]*components.LeadMovement),
		DistanceConstraints: make(map[entities.Entity]*components.DistanceConstraint),
		systems:             []System{},
		nextEntityID:        0,
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
		case *components.GravitationalPull:
			w.GravitationalPulls[entity] = c
		case *components.Circle:
			w.Circles[entity] = c
		case *components.Velocity:
			w.Velocities[entity] = c
		case *components.BounceBoundaries:
			w.BoundaryBouncings[entity] = c
		case *components.LeadMovement:
			w.LeadMovements[entity] = c
		case *components.DistanceConstraint:
			w.DistanceConstraints[entity] = c
		default:
			// Handle unknown component types if necessary
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

func (w *World) Reset() {
	w.systems = []System{}
	w.Positions = make(map[entities.Entity]*components.Position)
	w.Velocities = make(map[entities.Entity]*components.Velocity)
	w.GravitationalPulls = make(map[entities.Entity]*components.GravitationalPull)
	w.Circles = make(map[entities.Entity]*components.Circle)
	w.BoundaryBouncings = make(map[entities.Entity]*components.BounceBoundaries)
	w.LeadMovements = make(map[entities.Entity]*components.LeadMovement)
	w.DistanceConstraints = make(map[entities.Entity]*components.DistanceConstraint)
	w.nextEntityID = 0
}
