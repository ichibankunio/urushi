package urushi


type Tweener struct {
	Repeat int
	Target interface{}
	Delay int
	// Easing EasingMode
	start float64
	delta float64
	sigma float64
	dst float64
	duration float64

	ccRatio float64
}

func NewTweener(start float64, dst float64, duration float64) *Tweener {
	return &Tweener{
		Repeat: 0,
		Target: nil,
		Delay:  0,
		start: start,
		delta: 0,
		sigma: 0,
		dst: dst,
		duration: duration,
		ccRatio: (dst-start) / duration,
	}
}

func (t *Tweener) Update(cc int) (float64, bool) {
	t.sigma += 1/ 60
	// t.sigma += (t.dst-t.start) / t.duration
	value := easeIn(float64(cc) * t.ccRatio, t.start, t.dst - t.start, t.sigma)
	// t.delta = value - t.start

	return value, value == t.dst
}

func easeIn(t, b, c, d float64) float64 {
	t /= d
	return c*t*t + b
}


// func (t Tweener) To(tag string) *Tweener {

// }
