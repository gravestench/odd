package terminal

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/gravestench/akara"
	"github.com/gravestench/director/pkg/easing"
	"github.com/gravestench/director/pkg/systems/tween"
	"image/color"
	"strings"
	"time"

	"github.com/gravestench/director/pkg/systems/input"
	"github.com/gravestench/odd/engine/scenes/common"
)

const (
	defaultHistory = 100 // commands
)

// Scene is a graphical frontend for a shell
type Scene struct {
	common.SceneExtension

	window akara.EID
	text akara.EID

	input struct {
		index int
		buffer []string
	}

	output struct {
		index int
		buffer []string
	}

	isVisible bool

	tween struct {
		show *tween.Tween
		hide *tween.Tween
	}
}

func (scene *Scene) Key() string {
	return "terminal"
}

func (scene *Scene) Update() {
	txt, found := scene.Components.Text.Get(scene.text)
	if found {
		txt.String = strings.Join(scene.output.buffer, "\r\n")
	}
}

func (scene *Scene) Init(w *akara.World) {
	scene.SceneExtension.Init(w)
	scene.bindInput()
	scene.makeWindow()
	scene.setupTweens()
}

func (scene *Scene) IsInitialized() bool {
	return scene.Director != nil
}

func (scene *Scene) bindInput() {
	scene.bindOpenClose()
	scene.bindPrintableKeys()
}

func (scene *Scene) makeWindow() {
	ww, wh := scene.Sys.Renderer.Window.Width, scene.Sys.Renderer.Window.Height

	gray := color.RGBA{
		R: 196,
		G: 196,
		B: 196,
		A: 230,
	}

	scene.window = scene.Add.Rectangle(0, wh, ww, wh, gray, nil)
	scene.text = scene.Add.Label("", 200, 200, wh/20, "", color.RGBA{R:222, G:222, B:222, A:255})

	wNode, wFound := scene.Components.SceneGraphNode.Get(scene.window)
	tNode, tFound := scene.Components.SceneGraphNode.Get(scene.text)
	if wFound && tFound {
		tNode.SetParent(wNode.Node)
	}

	origin, found := scene.Components.Origin.Get(scene.window)
	if found {
		origin.X = 0
		origin.Y = 0
	}
}

func (scene *Scene) setupTweens() {
	scene.tween.show = scene.Sys.Tweens.New()
	scene.tween.hide = scene.Sys.Tweens.New()

	const (
		fullyVisible = 0.65
	)

	showFrom, hideFrom := 0., 1.

	duration := time.Second / 2

	origin, found := scene.Components.Origin.Get(scene.window)
	if !found {
		origin = scene.Components.Origin.Add(scene.window)
	}

	alpha, found := scene.Components.Opacity.Get(scene.window)
	if !found {
		alpha = scene.Components.Opacity.Add(scene.window)
	}

	scene.tween.show.
		Time(duration).
		Ease(easing.Sine).
		OnStart(func() {
			scene.Sys.Tweens.Remove(scene.tween.hide)
		}).
		OnUpdate(func(progress float64) {
			alpha.Value = progress
			origin.Y = (fullyVisible * progress) - showFrom
		}).
		OnComplete(func() {
			hideFrom = origin.Y
			scene.Sys.Tweens.Remove(scene.tween.show)
		})

	scene.tween.hide.
		Time(duration).
		Ease(easing.Sine).
		OnStart(func() {
			scene.Sys.Tweens.Remove(scene.tween.show)
		}).
		OnUpdate(func(progress float64) {
			alpha.Value = 1 - progress
			origin.Y = fullyVisible - (progress * hideFrom)
		}).
		OnComplete(func() {
			showFrom = origin.Y
			scene.Sys.Tweens.Remove(scene.tween.hide)
		})
}

func (scene *Scene) bindOpenClose() {
	e := scene.NewEntity()
	v := scene.Components.Interactive.Add(e)
	v.KeyVector.Set(int(input.KeyGraveAccent), true)

	v.Callback = func() (preventPropagation bool) {
		scene.isVisible = !scene.isVisible

		if scene.isVisible {
			scene.Sys.Tweens.Add(scene.tween.show.Start())
			scene.Sys.Tweens.Add(scene.tween.hide.Stop())
		} else {
			scene.Sys.Tweens.Add(scene.tween.hide.Start())
			scene.Sys.Tweens.Add(scene.tween.show.Stop())
		}

		return !scene.isVisible
	}
}

func (scene *Scene) bindPrintableKeys() {
	regularKeys := []input.Key{
		input.Key0, input.Key1, input.Key2, input.Key3, input.Key4, input.Key5, input.Key6,
		input.Key7, input.Key8, input.Key9, input.KeyA, input.KeyB, input.KeyC, input.KeyD,
		input.KeyE, input.KeyF, input.KeyG, input.KeyH, input.KeyI, input.KeyJ, input.KeyK,
		input.KeyL, input.KeyM, input.KeyN, input.KeyO, input.KeyP, input.KeyQ, input.KeyR,
		input.KeyS, input.KeyT, input.KeyU, input.KeyV, input.KeyW, input.KeyX, input.KeyY,
		input.KeyZ, input.KeyApostrophe, input.KeyBackslash, input.KeyComma, input.KeyEnter, input.KeyEqual,
		input.KeyKP0, input.KeyKP1, input.KeyKP2, input.KeyKP3, input.KeyKP4, input.KeyKP5,
		input.KeyKP6, input.KeyKP7, input.KeyKP8, input.KeyKP9, input.KeyKPAdd,
		input.KeyKPDecimal, input.KeyKPDivide, input.KeyKPEnter,
		input.KeyKPEqual, input.KeyKPMultiply, input.KeyKPSubtract, input.KeyLeft,
		input.KeyLeftBracket, input.KeyMenu, input.KeyMinus, input.KeyPeriod, input.KeyRightBracket,
		input.KeySemicolon, input.KeySlash,
		input.KeySpace,
	}

	for _, k := range regularKeys {
		e := scene.NewEntity()
		v := scene.Components.Interactive.Add(e)
		v.KeyVector.Set(int(k), true)
		v.Callback = func() (preventPropagation bool) {
			if !scene.isVisible {
				return false
			}

			if len(scene.output.buffer) < 1 {
				scene.output.buffer = append(scene.output.buffer, "")
			}

			ch := string(rl.GetKeyPressed())
			scene.output.buffer[len(scene.output.buffer)-1] += ch

			return scene.isVisible
		}
	}
}
