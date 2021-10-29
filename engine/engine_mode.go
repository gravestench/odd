package engine

type Mode int

const (
	ModeClient Mode = iota
	ModeHost
	ModeServer
)