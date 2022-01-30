package ui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/ichibankunio/urushi"
)

type config struct {
	texts [2]map[string]string
	soundBtn Button
	reviewButton Button
}



var configText [2]map[string]string

func init() {
	configText[0] = map[string]string{"sound": "サウンド", "review": "このアプリをレビューする", "back": "もどる", "bgm": "BGM音量", "se": "SE音量"}
	configText[1] = map[string]string{"sound": "サウンド", "review": "このアプリをレビューする", "back": "もどる", "bgm": "BGM音量", "se": "SE音量"}
}

func NewSettingScene(ID urushi.SceneID) *urushi.Scene {
	return &urushi.Scene{
		Update: func(g *urushi.Game) error {

			return nil
		},
		Draw: func(screen *ebiten.Image) {
			
		},
		ID: ID,
	}
}