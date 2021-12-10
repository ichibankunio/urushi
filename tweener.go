package urushi


type Tweener struct {
	Repeat int
	Target interface{}
	Delay int
	// Easing EasingMode
	
}

func NewTweener() *Tweener {
	return &Tweener{
		Repeat: 0,
		Target: nil,
		Delay:  0,
		// Easing: 0,
	}
}

// func (t Tweener) To(tag string) *Tweener {

// }
