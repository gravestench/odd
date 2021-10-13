package engine

import (
	"encoding/json"
	"fmt"
	"github.com/gravestench/odd/engine/scenes/terminal"
	"os"
	"path"

	"github.com/gravestench/odd/engine/scenes/odd_error"
	"github.com/gravestench/odd/engine/scenes/splash"

	"github.com/kirsle/configdir"

	director "github.com/gravestench/director/pkg"

	"github.com/gravestench/odd/engine/mpq_loader"
	"github.com/gravestench/odd/engine/settings"
)

const (
	oddAppName     = "odd"
	oddAppSettings = "settings.json"
	oddWindowTitle = "ODD Engine"
)

type Mode int

const (
	Client Mode = iota
	Host
	Server
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

	e.initEngineSystems()

	if e.Mode != Server {
		e.initEngineScenes()
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
	//e.AddSystem(&ai.System{}) // TODO
	//e.AddSystem(&ai.System{}) // TODO
	//e.AddSystem(&combat_evaluation.System{}) // TODO
	//e.AddSystem(&item_craft.System{}) // TODO
	//e.AddSystem(&damage_calculation.System{}) // TODO
	//e.AddSystem(&entity_control.System{}) // TODO
	//e.AddSystem(&merchant_items.System{}) // TODO
	//e.AddSystem(&health_modifier_queue.System{}) // TODO
	//e.AddSystem(&merchant_mercenaries.System{}) // TODO
	//e.AddSystem(&client.System{}) // TODO
	//e.AddSystem(&server.System{}) // TODO
	//e.AddSystem(&player_party.System{}) // TODO
	//e.AddSystem(&player_quest.System{}) // TODO
	//e.AddSystem(&realm.System{}) // TODO
	//e.AddSystem(&shell.System{}) // TODO
	//e.AddSystem(&player_skills.System{}) // TODO
	//e.AddSystem(&item_socket.System{}) // TODO
	//e.AddSystem(&item_treasure.System{}) // TODO
}

func (e *Engine) initEngineScenes() {
	e.AddScene(&splash.Scene{})
	e.AddScene(&terminal.Scene{})
	//e.AddScene(&game_ui_chat.Scene{}) // TODO
	//e.AddScene(&game_ui_cube.Scene{}) // TODO
	//e.AddScene(&game_ui_escape_menu.Scene{}) // TODO
	//e.AddScene(&game_ui_help.Scene{}) // TODO
	//e.AddScene(&game_ui_hud.Scene{}) // TODO
	//e.AddScene(&game_ui_inventory.Scene{}) // TODO
	//e.AddScene(&game_ui_map.Scene{}) // TODO
	//e.AddScene(&game_ui_mercenary.Scene{}) // TODO
	//e.AddScene(&game_ui_options_config.Scene{}) // TODO
	//e.AddScene(&game_ui_party.Scene{}) // TODO
	//e.AddScene(&game_ui_quest.Scene{}) // TODO
	//e.AddScene(&game_ui_skills.Scene{}) // TODO
	//e.AddScene(&game_ui_stash.Scene{}) // TODO
	//e.AddScene(&game_ui_stats.Scene{}) // TODO
	//e.AddScene(&game_ui_tooltip.Scene{}) // TODO
	//e.AddScene(&game_ui_trade_mercenary.Scene{}) // TODO
	//e.AddScene(&game_ui_trade_merchant.Scene{}) // TODO
	//e.AddScene(&game_ui_trade_player.Scene{}) // TODO
	//e.AddScene(&game_weather.Scene{}) // TODO
	//e.AddScene(&game_world.Scene{}) // TODO
	//e.AddScene(&loading.Scene{}) // TODO
	//e.AddScene(&menu_character_create.Scene{}) // TODO
	//e.AddScene(&menu_character_load.Scene{}) // TODO
	//e.AddScene(&menu_cinematics.Scene{}) // TODO
	//e.AddScene(&menu_credits.Scene{}) // TODO
	//e.AddScene(&menu_end_game.Scene{}) // TODO
	//e.AddScene(&menu_main.Scene{}) // TODO
	//e.AddScene(&menu_realm_browser.Scene{}) // TODO
	//e.AddScene(&menu_realm_host.Scene{}) // TODO
	//e.AddScene(&menu_realm_join.Scene{}) // TODO
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