package realm

import "github.com/gravestench/akara"

// TODO

type System struct {
	akara.BaseSystem
}

func (sys *System) Update() {
	// noop
}

func (sys *System) Init(_ *akara.World) {
	// noop
}

func (sys *System) IsInitialized() bool {
	return true
}

