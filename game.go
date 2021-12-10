package urushi

import (
	"math"
	"time"
)

type Game struct {
	lastFrameTime int
	thisFrameTime int
	deltaTime float64
	counter int
}

func (g *Game) Update() {
	g.thisFrameTime = time.Now().Nanosecond()
	
	g.deltaTime = float64(g.thisFrameTime-g.lastFrameTime) / math.Pow(10, float64(int(math.Log10(float64(g.thisFrameTime-g.lastFrameTime))+2)))
	if g.deltaTime < 0 || math.IsNaN(g.deltaTime) {
		g.deltaTime = 0
	}
	g.lastFrameTime = g.thisFrameTime

	g.counter++
}
