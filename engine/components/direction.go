package components

import (
	"github.com/gravestench/akara"
)

var _ akara.Component = &Direction{}

// Direction represents an entity's direction (that it is facing) in the isometric world.
// The 0-value of angle should be to the bottom-left of the screen.
type Direction struct {
	Angle float64
}

// New creates a new color component
func (*Direction) New() akara.Component {
	return &Direction{}
}

// DirectionFactory is a wrapper for the generic component factory.
type DirectionFactory struct {
	*akara.ComponentFactory
}

// Add a Direction component to the given entity and return it
func (m *DirectionFactory) Add(id akara.EID) *Direction {
	return m.ComponentFactory.Add(id).(*Direction)
}

// Get returns the Direction component for the given entity, and a bool for whether or not it exists
func (m *DirectionFactory) Get(id akara.EID) (*Direction, bool) {
	component, found := m.ComponentFactory.Get(id)
	if !found {
		return nil, found
	}

	return component.(*Direction), found
}
