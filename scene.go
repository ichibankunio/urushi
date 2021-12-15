package urushi

import "github.com/hajimehoshi/ebiten/v2"

type Scene struct {
	Update func(*Game)
	Draw func(*ebiten.Image)
	ID SceneID
}

// func NewScene(update func(*Game), draw func(*ebiten.Image)) *Scene {
//     return &Scene{
// 		Update: 
// 	}
// }

// func (s *Scene) Update(g *Game) {

// }

// func (s *Scene) Draw(screen *ebiten.Image) {

// }

func ShiftSceneWithExpandingImage(next int) {

}