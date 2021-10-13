package settings

import (
	"os"
)

type Archives struct {
	Directory string
	LoadOrder []string
}

func (a *Archives) init() {
	a.Directory, _ = os.Getwd()
	a.LoadOrder = []string{
		"patch_d2.mpq",
		"d2exp.mpq",
		"d2xmusic.mpq",
		"d2xtalk.mpq",
		"d2xvideo.mpq",
		"d2data.mpq",
		"d2char.mpq",
		"d2music.mpq",
		"d2sfx.mpq",
		"d2video.mpq",
		"d2speech.mpq",
	}
}
