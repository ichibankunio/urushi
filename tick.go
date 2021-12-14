package urushi

import "github.com/hajimehoshi/ebiten/v2"

type TickF struct {
	span int
	content func(int, []interface{}) bool
	contentCounter int
	repeat int
	pause bool
	// content2 func(int, ...TickOptions) bool
}

// type TickOptions struct {
// 	target interface{}
// 	tweener Tweener
// }

func (t *TickF) rewind() {
	t.contentCounter = 0
	t.pause = true
}

func (t *TickF) Update(g *Game, i ...interface{}) {
	if !t.pause {
		if g.counter % t.span == 0 {
			if ebiten.MaxTPS() == 30 {
				t.content(t.contentCounter, i)
			}
			if t.content(t.contentCounter, i) || (t.contentCounter == t.repeat - 1 && t.repeat != -1) {
				t.rewind()
			}else {
				t.contentCounter++
			}
		}
	}
}

func NewTickF(span int, pause bool, repeat int, content func(int, []interface{}) bool) *TickF {
	return &TickF{
		span: span,
		content: content,
		pause: pause,
		repeat: repeat,
	}
}

// func NewTickF2(span int, content func (int, ...TickOptions) bool) *TickF {
// 	return &TickF{
// 		span: span,
// 		content2: content,
// 	}
// }

// func (t *TickF) Update2(g *Game, i ...interface{}) {
// 	if g.counter % t.span == 0 {
// 		if ebiten.MaxTPS() == 30 {
// 			t.content(t.contentCounter, i)
// 		}
// 		if t.content(t.contentCounter, i) {
// 			// t.rewind()
// 		}else {
// 			t.contentCounter++
// 		}
// 	}
// }

