package urushi

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Scene struct {
	Update func(*Game) error
	Draw func(*ebiten.Image)
	ID SceneID
}

func ShiftSceneWithExpandingSprite(spr Sprite, clr color.Color, ID SceneID) *Scene {
	return &Scene{
		Update: func(g *Game) error {
			return nil
		},
		Draw: func(screen *ebiten.Image) {
			
		},
		ID: ID,
	}
}

func ShiftSceneWithShrinkingSprite(spr Sprite, clr color.Color, ID SceneID) *Scene {
	return &Scene{
		Update: func(g *Game) error {
			return nil
		},
		Draw: func(screen *ebiten.Image) {
			
		},
		ID: ID,
	}
}
