package urushi


type Tweener struct {
	Repeat int
	Target interface{}
	Delay int
	// Easing EasingMode
	start float64
	delta float64
	sigma float64
}

func NewTweener(start float64) *Tweener {
	return &Tweener{
		Repeat: 0,
		Target: nil,
		Delay:  0,
		start: start,
		delta: 0,
		sigma: 0,
		// Easing: 0,
	}
}

func (t *Tweener) Update(cc int, destination float64) (float64, bool) {
	value := easeIn(float64(cc), t.start, destination - t.start, t.sigma)
	// t.delta = value - t.start
	t.sigma ++

	return value, false
}

func easeIn(t, b, c, d float64) float64 {
	t /= d
	return c*t*t*t*t + b
}


// func (t Tweener) To(tag string) *Tweener {

// }
