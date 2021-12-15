package urushi

import (
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	lastFrameTime int
	thisFrameTime int
	deltaTime float64
	counter int

	scenes []*Scene
	State SceneID
}

type SceneID int

func (g *Game) Update() {
	g.thisFrameTime = time.Now().Nanosecond()
	
	g.deltaTime = float64(g.thisFrameTime-g.lastFrameTime) / math.Pow(10, float64(int(math.Log10(float64(g.thisFrameTime-g.lastFrameTime))+2)))
	if g.deltaTime < 0 || math.IsNaN(g.deltaTime) {
		g.deltaTime = 0
	}
	g.lastFrameTime = g.thisFrameTime

	g.counter++

	for _, scene := range g.scenes {
		if g.State == scene.ID {
			scene.Update(g)

			break
		}
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, scene := range g.scenes {
		if g.State == scene.ID {
			scene.Draw(screen)
			break
		}
	}
}

func (g *Game) AddScene(scene *Scene) {
	g.scenes = append(g.scenes, scene)
}
