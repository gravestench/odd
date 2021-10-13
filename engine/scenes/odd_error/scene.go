package odd_error

import (
	"github.com/gravestench/akara"
	"github.com/gravestench/director/pkg/systems/scene"
	"image/color"
)

// TODO

func New(message string) *Scene {
	s := &Scene{
		message: message,
	}

	return s
}

type Scene struct {
	scene.Scene
	message string
}

func (scene *Scene) Update() {

}

func (scene *Scene) Key() string {
	return "odd error"
}

func (scene *Scene) Init(_ *akara.World) {
	c := color.RGBA{R: 200, A: 255}
	ww, wh := scene.Sys.Renderer.Window.Width, scene.Sys.Renderer.Window.Height
	fontSize := wh / 10

	e := scene.Add.Label(scene.message, ww/2, wh/2, fontSize, "", c)

	trs, found := scene.Components.Transform.Get(e)
	if !found {
		return
	}

	trs.Translation.Set(float64(ww/2), float64(wh/2), trs.Translation.Z)
}

func (scene *Scene) IsInitialized() bool {
	return true
}

