package components

import (
	"github.com/gravestench/akara"
)

var _ akara.Component = &CommandRegistration{}

// CommandRegistration is a wrapper component for a color.CommandRegistration interface
type CommandRegistration struct {
	Enabled     bool
	Name        string
	Description string
	Args        []string
	Callback    func(args []string) error
}

// New creates a new color component
func (*CommandRegistration) New() akara.Component {
	return &CommandRegistration{}
}

// CommandRegistrationFactory is a wrapper for the generic component factory.
type CommandRegistrationFactory struct {
	*akara.ComponentFactory
}

// Add a CommandRegistration component to the given entity and return it
func (m *CommandRegistrationFactory) Add(id akara.EID) *CommandRegistration {
	return m.ComponentFactory.Add(id).(*CommandRegistration)
}

// Get returns the CommandRegistration component for the given entity, and a bool for whether or not it exists
func (m *CommandRegistrationFactory) Get(id akara.EID) (*CommandRegistration, bool) {
	component, found := m.ComponentFactory.Get(id)
	if !found {
		return nil, found
	}

	return component.(*CommandRegistration), found
}
