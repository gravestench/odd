package menu_credits

import (
	"time"

	"github.com/gravestench/akara"
	"github.com/gravestench/director/pkg/systems/scene"
)

// TODO

type Scene struct {
	scene.Scene
}

func (scene *Scene) Update(duration time.Duration) {
	// noop
}

func (scene *Scene) Init(_ *akara.World) {
	// noop
}

func (scene *Scene) IsInitialized() bool {
	return true
}
