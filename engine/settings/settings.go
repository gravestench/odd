package settings

type Settings struct {
	Archives Archives
	Renderer Renderer
	Keys     Keys
}

func (s *Settings) Init() {
	s.Archives.init()
	s.Renderer.init()
	s.Keys.init()
}
