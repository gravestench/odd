package engine

import (
	"encoding/json"
	"fmt"
	"github.com/gravestench/odd/engine/systems/network/client"
	"github.com/gravestench/odd/engine/systems/network/server"
	"github.com/gravestench/odd/engine/systems/shell"
	"os"
	"path"

	"github.com/gravestench/director"
	"github.com/kirsle/configdir"

	"github.com/gravestench/odd/engine/mpq_loader"
	"github.com/gravestench/odd/engine/settings"

	"github.com/gravestench/odd/engine/scenes/odd_error"
	"github.com/gravestench/odd/engine/scenes/splash"
	"github.com/gravestench/odd/engine/scenes/terminal"
)

const (
	oddAppName     = "odd"
	oddAppSettings = "settings.json"
	oddWindowTitle = "ODD Engine"
)

// New creates a new engine instance
func New() *Engine {
	e := &Engine{
		Director: director.New(),
	}

	e.Sys.Renderer.Window.Title = oddWindowTitle

	err := e.initSettings()
	if err != nil {
		errMsg := fmt.Sprintf("could not initialize configuration: %w", err)
		e.showErrorScene(errMsg)

		return e
	}

	err = e.initLoaders()
	if err != nil {
		errMsg := fmt.Sprintf("could not initialize loaders: %w", err)
		e.showErrorScene(errMsg)

		return e
	}

	e.initEngineSystems() // we always init all systems

	switch e.Mode {
	case ModeClient, ModeHost:
		// client and host are graphical, so need the graphical scenes
		e.initEngineScenes()
	case ModeServer:
		// the dedicated server only needs the non-graphical systems
		// nothing else to do here
	}

	return e
}

// Engine represents the core ODD engine
type Engine struct {
	*director.Director
	settings.Settings
	Mode
}

func (e *Engine) initSettings() error {
	err := configdir.MakePath(configdir.LocalConfig(oddAppName))
	if err != nil {
		return fmt.Errorf("could not create local config directory: %w", err)
	}

	settingsPath := path.Join(configdir.LocalConfig(oddAppName), oddAppSettings)

	if _, err := os.Stat(settingsPath); err != nil {
		err := e.initSettingsFile(&e.Settings, settingsPath)
		if err != nil {
			return err
		}
	} else {
		fh, err := os.Open(settingsPath)
		if err != nil {
			panic(err)
		}

		defer fh.Close()

		decoder := json.NewDecoder(fh)

		if err := decoder.Decode(&e.Settings); err != nil {
			_ = e.initSettingsFile(&e.Settings, settingsPath)
		}
	}

	return nil
}

func (e *Engine) initLoaders() error {
	for _, mpqName := range e.Settings.Archives.LoadOrder {
		mpqPath := path.Join(e.Settings.Archives.Directory, mpqName)

		p, err := mpq_loader.New(mpqPath)
		if err != nil {
			return fmt.Errorf("could not find MPQ: %w", err)
		}

		e.Sys.Load.AddProvider(p)
	}

	return nil
}

func (e *Engine) initEngineSystems() {
	e.AddSystem(&shell.System{}, true)
	e.AddSystem(&client.System{}, true)
	e.AddSystem(&server.System{}, true)
}

func (e *Engine) initEngineScenes() {
	e.AddScene(&splash.Scene{})
	e.AddScene(&terminal.Scene{})
}

func (e *Engine) initSettingsFile(s *settings.Settings, settingsPath string) error {
	s.Init()

	fileHandle, err := os.Create(settingsPath)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(fileHandle)
	encoder.SetIndent("", "\t")

	err = encoder.Encode(&s)
	if err != nil {
		return err
	}

	return nil
}

func (e *Engine) showErrorScene(errMsg string) {
	e.Sys.Renderer.Window.Width, e.Sys.Renderer.Window.Height = 600, 300
	e.AddScene(odd_error.New(errMsg))
}