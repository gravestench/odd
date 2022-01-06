package common

import (
	"github.com/gravestench/akara"

	"github.com/gravestench/director"

	"github.com/gravestench/odd/engine/components"
)

// SceneExtension acts as an extension of the generic director scene
type SceneExtension struct {
	director.Scene
	Components struct {
		*director.SceneComponents
		CommandRegistration components.CommandRegistrationFactory
	}
	Sys struct {
		*director.SceneSystems
	}
	Add struct {
		*director.SceneObjects
	}
}

func (s *SceneExtension) NewEntity() director.Entity {
	return s.Scene.Director.NewEntity()
}

func (s *SceneExtension) Init(_ *akara.World) {
	s.Components.SceneComponents = &s.Scene.Components
	s.Sys.SceneSystems = &s.Scene.Director.Sys
	s.Add.SceneObjects = &s.Scene.Add
	s.SetTickFrequency(60)
}

func (s *SceneExtension) IsInitialized() bool {
	panic("implement me")
}
