package splash

import (
	"image/color"
	"time"

	"github.com/gravestench/akara"
	"github.com/gravestench/odd/engine/scenes/common"
)

// TODO

type Scene struct {
	common.SceneExtension
}

func (scene *Scene) Key() string {
	return "splash"
}

func (scene *Scene) Update() {
 // noop
}

func (scene *Scene) Init(w *akara.World) {
	scene.SceneExtension.Init(w)

	scene.makeShapes()
	scene.fadeInOut()
}

func (scene *Scene) makeShapes() {
	ww := scene.Sys.Renderer.Window.Width
	wh := scene.Sys.Renderer.Window.Height
	cw := wh / 8
	pad := cw * 11 / 10
	cellWidth := cw + pad

	const (
		numShapes = 3
	)

	white := color.RGBA{R: 255, G: 255, B: 255, A: 255}
	//black := color.RGBA{A: 255}

	x, y := (ww/2)-cellWidth*2, wh/2
	layer := numShapes * 2

	for idx := 0; idx < numShapes; idx++ {
		x += cellWidth

		c := scene.Add.Circle(x, y, cw, white, nil)
		circleRenderOrder := scene.Components.RenderOrder.Add(c)

		if idx == 0 {
			circleRenderOrder.Value = layer
			layer--
			continue
		}


		circleRenderOrder.Value = layer - 2

		r := scene.Add.Rectangle(x, y, cw, cw*2, white, nil)
		rectRenderOrder := scene.Components.RenderOrder.Add(r)

		rectOrigin, _ := scene.Components.Origin.Get(r)
		rectOrigin.X = 1

		rectRenderOrder.Value = layer - 1

		layer -= 2
	}
}

func (scene *Scene) fadeInOut() {
	t1 := scene.Sys.Tweens.New()
	t2 := scene.Sys.Tweens.New()

	t1.Ease("Sine")
	t2.Ease("Sine")

	delay := time.Second
	duration := time.Second

	t1.
		Delay(delay).
		Time(duration).
		OnStart(func(){
			if scene.Viewports == nil {
				return
			}

			opacity, _ := scene.Components.Opacity.Get(scene.Viewports[0])
			opacity.Value = 0
		}).
		OnUpdate(func(progress float64){
			if scene.Viewports == nil {
				return
			}

			opacity, _ := scene.Components.Opacity.Get(scene.Viewports[0])
			opacity.Value = progress
		}).
		OnComplete(func(){
			if scene.Viewports == nil {
				return
			}

			opacity, _ := scene.Components.Opacity.Get(scene.Viewports[0])
			opacity.Value = 1
			scene.Sys.Tweens.Remove(t1)
			scene.Sys.Tweens.Add(t2)
		})

	t2.
		Time(duration).
		OnStart(func(){
			opacity, _ := scene.Components.Opacity.Get(scene.Viewports[0])
			opacity.Value = 1
		}).
		OnUpdate(func(progress float64){
			opacity, _ := scene.Components.Opacity.Get(scene.Viewports[0])
			opacity.Value = 1 - progress
		}).
		OnComplete(func(){
			opacity, _ := scene.Components.Opacity.Get(scene.Viewports[0])
			opacity.Value = 0
			scene.Sys.Tweens.Remove(t2)
		})

	scene.Sys.Tweens.Add(t1)
}

func (scene *Scene) IsInitialized() bool {
	return true
}
