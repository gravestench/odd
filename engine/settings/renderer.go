package settings

const (
	defaultWidth, defaultHeight = 1024, 768
)

type Renderer struct {
	Resolution struct {
		Width, Height int
	}
}

func (r *Renderer) init() {
	r.Resolution.Width = defaultWidth
	r.Resolution.Height = defaultHeight
}
