package settings

import "github.com/gravestench/director/pkg/systems/input"

type Keys struct {
	Terminal input.Key

	SkillTree, SpeedSkill,
	Skill1, Skill2, Skill3, Skill4,
	Skill5, Skill6, Skill7, Skill8,
	Skill9, Skill10, Skill11, Skill12,
	Skill13, Skill14, Skill15, Skill16,
	SkillNext, SkillPrev input.Key

	Belt1, Belt2, Belt3, Belt4, ShowBelt input.Key

	Inventory, Mercenary, Stats input.Key

	Party, Quests, Messages, Help input.Key

	ShowItems, SwapWeapons, Run, StandStill, ToggleRun input.Key

	MapToggle, MapCenter, MapFade,
	MapShowNames, MapShowEntities, MapMinimap input.Key

	SayHelp, SayFollow, SayForYou, SayThanks,
	SaySorry, SayBye, SayDie, SayRetreat input.Key
}

func (keys *Keys) init() {
	keys.Terminal = input.KeyGraveAccent

	keys.Skill1 = input.KeyF1
	keys.Skill2 = input.KeyF2
	keys.Skill3 = input.KeyF3
	keys.Skill4 = input.KeyF4
	keys.Skill5 = input.KeyF5
	keys.Skill6 = input.KeyF6
	keys.Skill7 = input.KeyF7
	keys.Skill8 = input.KeyF8

	keys.SayHelp = input.KeyKP1
	keys.SayFollow = input.KeyKP2
	keys.SayForYou = input.KeyKP3
	keys.SayThanks = input.KeyKP4
	keys.SaySorry = input.KeyKP5
	keys.SayBye = input.KeyKP6
	keys.SayDie = input.KeyKP7
	keys.SayRetreat = input.KeyKP8

	keys.SkillTree = input.KeyT

	keys.Belt1 = input.Key1
	keys.Belt2 = input.Key2
	keys.Belt3 = input.Key3
	keys.Belt4 = input.Key4

	keys.Inventory = input.KeyI
	keys.Stats = input.KeyC
	keys.Mercenary = input.KeyO
	keys.Party = input.KeyP
	keys.Quests = input.KeyQ
	keys.Messages = input.KeyM
	keys.Help = input.KeyH
	keys.MapToggle = input.KeyTab
	keys.ShowBelt = input.ModControlLeft
	keys.StandStill = input.ModShiftLeft
	keys.SwapWeapons = input.KeyW
	keys.ToggleRun = input.KeyR
}
