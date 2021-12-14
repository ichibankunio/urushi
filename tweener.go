package urushi

import "github.com/hajimehoshi/ebiten/v2"


type Tweener struct {
	Repeat int
	Target interface{}
	Delay int
	// Easing EasingMode
	begin float64
	end float64
	duration float64
	easing Easing
}

type Easing int

const (
	EaseIn Easing = iota
)

func NewTweener(begin float64, end float64, duration float64, easing Easing) *Tweener {
	return &Tweener{
		Repeat: 0,
		Target: nil,
		Delay:  0,
		begin: begin,
		end: end,
		duration: duration,
		easing: easing,
	}
}

func SecToFrame(sec float64) float64 {
	return sec * float64(ebiten.MaxTPS())
}

func (t *Tweener) Update(cc int) (float64, bool) {
	// t.sigma ++
	// t.sigma += (t.dst-t.start) / t.duration
	var value float64
	switch t.easing {
	case EaseIn:
		value = easeIn(float64(cc), t.begin, t.end - t.begin, t.duration)
	}
	// t.delta = value - t.start

	return value, value == t.end
}

func easeIn(t, b, c, d float64) float64 {
	t /= d
	return c*t*t + b
}


// func (t Tweener) To(tag string) *Tweener {

// }
