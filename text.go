package urushi

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
)

type TxtSpr struct {
	txt     string
	spr *Sprite
	clr     color.Color
	padUp   int
	padLeft int
	font    font.Face
	isVert  bool
	alpha float64
	hidden bool
}

func NewTxtSpr(txt string, x, y float64, clr color.Color, bgClr color.Color, font font.Face, padUp, padLeft int) *TextSprite {
	var bgImg *ebiten.Image

	width := text.BoundString(font, txt).Dx()
	if width+padLeft*2 == 0 {
		width = 1
	}
	bgImg = ebiten.NewImage(width+padLeft*2, font.Metrics().Height.Ceil()+padUp*2)

	bgImg.Fill(bgClr)	

	t := &TxtSpr{
		txt: txt,
		spr: NewSprite(bgImg, x, y),
		clr: clr,
		font: font,
		padUp: 0,
		padLeft: 0,
		alpha: 1,
	}
	return t
}

func (t *TxtSpr) SetCenter(center int) {
	if t.isVert {
		t.spr.X = float64(center - t.font.Metrics().Height.Ceil()/2 - t.padLeft)
	} else {
		width := text.BoundString(t.font, t.txt).Dx()
		t.spr.X = float64(center - width/2 - t.padLeft)
	}
}

func (t *TxtSpr) Draw(screen *ebiten.Image) {
	if !t.hidden {
		if t.isVert {
			t.spr.Draw(screen)
		
			for i, v := range []rune(t.txt) {
				if string(v) == "、" || string(v) == "。" {
					// Perhaps int(t.spr.x) is not good?
					text.Draw(screen, string(v), t.font, int(t.spr.X)+t.padLeft+t.font.Metrics().Height.Ceil()-text.BoundString(t.font, string(v)).Dx(), int(t.spr.Y)-t.font.Metrics().Height.Ceil()/8+t.padUp+i*t.font.Metrics().Height.Ceil()+t.font.Metrics().Height.Ceil()-text.BoundString(t.font, string(v)).Dy(), t.clr)
				} else {
					text.Draw(screen, string(v), t.font, int(t.spr.X)+t.padLeft, int(t.spr.Y)-t.font.Metrics().Height.Ceil()/8+t.padUp+t.font.Metrics().Height.Ceil()+i*t.font.Metrics().Height.Ceil(), t.clr)
	
				}
			}
		} else {
			t.spr.Draw(screen)
			text.Draw(screen, t.txt, t.font, int(t.spr.X)+t.padLeft, int(t.spr.Y)-t.font.Metrics().Height.Ceil()/8+t.font.Metrics().Height.Ceil()+t.padUp, t.clr)
		}
	}

}