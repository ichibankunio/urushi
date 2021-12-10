package urushi

import "github.com/hajimehoshi/ebiten/v2"

type TickF struct {
	span int
	content func(int, ...interface{}) bool
	contentCounter int
	repeat int
}

func (t *TickF) update(g *Game, i ...interface{}) {
	if g.counter % t.span == 0 {
		if ebiten.MaxTPS() == 30 {
			t.content(t.contentCounter, i)
		}
		if t.content(t.contentCounter, i) {
			// t.rewind()
		}else {
			t.contentCounter++
		}
	}
	
}

func NewTickF(span int, content func(int, ...interface{}) bool) *TickF {
	return &TickF{
		span: span,
		content: content,
	}
}

