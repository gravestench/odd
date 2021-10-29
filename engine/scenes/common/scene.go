package common

import (
	"github.com/gravestench/akara"
	director "github.com/gravestench/director/pkg"
	"github.com/gravestench/director/pkg/common"
	"github.com/gravestench/director/pkg/systems/scene"

	"github.com/gravestench/odd/engine/components"
)

// SceneExtension acts as an extension of the generic director scene
type SceneExtension struct {
	scene.Scene
	Components struct {
		*common.BasicComponents
		CommandRegistration components.CommandRegistrationFactory
	}
	Sys struct {
		*director.DirectorSystems
	}
	Add struct {
		*scene.ObjectFactory
	}
}

func (s *SceneExtension) NewEntity() common.Entity {
	return s.Scene.Director.NewEntity()
}

func (s *SceneExtension) Init(_ *akara.World) {
	s.Components.BasicComponents = &s.Scene.Components
	s.Sys.DirectorSystems = &s.Scene.Director.Sys
	s.Add.ObjectFactory = &s.Scene.Add
	s.SetTickFrequency(60)
}

func (s *SceneExtension) IsInitialized() bool {
	panic("implement me")
}

