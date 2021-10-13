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

func (odd *SceneExtension) NewEntity() common.Entity {
	return odd.Scene.Director.NewEntity()
}

func (odd *SceneExtension) Init(_ *akara.World) {
	odd.Components.BasicComponents = &odd.Scene.Components
	odd.Sys.DirectorSystems = &odd.Scene.Director.Sys
	odd.Add.ObjectFactory = &odd.Scene.Add
}

func (odd *SceneExtension) IsInitialized() bool {
	panic("implement me")
}

