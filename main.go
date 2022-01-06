package main

import (
	"github.com/gravestench/odd/engine"
)

func main() {
	e := engine.New()

	applyCommandLineFlags(e)

	if err := e.Run(); err != nil {
		panic(err)
	}
}

func applyCommandLineFlags(e *engine.Engine) *engine.Engine {
	return e
}
