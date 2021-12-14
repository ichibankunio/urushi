package urushi

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

type TxtSprShadow struct {
	spr *Sprite
	txt string
	clr color.Color
	shadowClr color.Color
	font font.Face
	padUp int
	padLeft int
	shadowX int
	shadowY int

	front *TxtSpr
	back *TxtSpr

	hidden bool
	alpha float64
}

func NewTxtSprShadow(txt string, x, y float64, clr color.Color, shadowClr color.Color, font font.Face, padUp, padLeft, shadowX, shadowY int, isVert bool) *TxtSprShadow {
	var bgImg *ebiten.Image
	if isVert {
		height := font.Metrics().Height.Ceil() * len([]rune(txt))
		bgImg = ebiten.NewImage(font.Metrics().Height.Ceil()+padLeft*2, height+padUp*2)
	} else {
		width := text.BoundString(font, txt).Dx()
		bgImg = ebiten.NewImage(width+padLeft*2, font.Metrics().Height.Ceil()+padUp*2)
	}
	back := NewTxtSpr(txt, x + float64(shadowX), y + float64(shadowY), shadowClr, color.Transparent, font, padUp, padLeft)
	front := NewTxtSpr(txt, x, y, clr, color.Transparent, font, padUp, padLeft)
	
	tss := &TxtSprShadow{
		txt: txt,
		clr: clr,
		shadowClr: shadowClr,
		spr: NewSprite(bgImg, x, y),
		font: font,
		padUp: padUp,
		padLeft: padLeft,
		shadowX: shadowX,
		shadowY: shadowY,
		back: back,
		front: front,
		hidden: false,
		alpha: 1,
	}
	return tss
}

func (tss *TxtSprShadow) SetCenter(center int) {
	width := text.BoundString(tss.font, tss.txt).Dx()
	tss.spr.X = float64(center - width/2 - tss.padLeft)
}

func (tss *TxtSprShadow) Draw(screen *ebiten.Image) {
	if !tss.hidden {
		tss.spr.Draw(screen)
		text.Draw(screen, tss.txt, tss.font, int(tss.spr.X)+tss.padLeft + tss.shadowX, int(tss.spr.Y)-tss.font.Metrics().Height.Ceil()/8+tss.font.Metrics().Height.Ceil()+tss.padUp + tss.shadowY, tss.shadowClr)
		text.Draw(screen, tss.txt, tss.font, int(tss.spr.X)+tss.padLeft, int(tss.spr.Y)-tss.font.Metrics().Height.Ceil()/8+tss.font.Metrics().Height.Ceil()+tss.padUp, tss.clr)
	}
}