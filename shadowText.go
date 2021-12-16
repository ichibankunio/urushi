package urushi

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

type TxtSprShadow struct {
	Spr *Sprite
	Txt string
	Clr color.Color
	ShadowClr color.Color
	Font font.Face
	PadUp int
	PadLeft int
	ShadowX int
	ShadowY int

	Hidden bool
}

func NewTxtSprShadow(txt string, x, y float64, clr color.Color, shadowClr color.Color, font font.Face, padUp, padLeft, shadowX, shadowY int, hidden bool) *TxtSprShadow {
	var bgImg *ebiten.Image
	width := text.BoundString(font, txt).Dx()
	bgImg = ebiten.NewImage(width+padLeft*2, font.Metrics().Height.Ceil()+padUp*2)
	
	// back := NewTxtSpr(txt, x + float64(shadowX), y + float64(shadowY), shadowClr, color.Transparent, font, padUp, padLeft)
	// front := NewTxtSpr(txt, x, y, clr, color.Transparent, font, padUp, padLeft)
	
	tss := &TxtSprShadow{
		Txt: txt,
		Clr: clr,
		ShadowClr: shadowClr,
		Spr: NewSprite(bgImg, x, y),
		Font: font,
		PadUp: padUp,
		PadLeft: padLeft,
		ShadowX: shadowX,
		ShadowY: shadowY,
		Hidden: false,
	}
	return tss
}

func (t *TxtSprShadow) SetCenter(center int) {
	width := text.BoundString(t.Font, t.Txt).Dx()
	t.Spr.X = float64(center - width/2 - t.PadLeft)
}

func (t *TxtSprShadow) SetText(txt string) {
	t.Txt = txt
	width := text.BoundString(t.Font, txt).Dx()
	if width+t.PadLeft*2 == 0 {
		width = 1
	}
	t.Spr.Img = ebiten.NewImage(width+t.PadLeft*2, t.Font.Metrics().Height.Ceil()+t.PadUp*2)
}

func (t *TxtSprShadow) Draw(screen *ebiten.Image) {
	if !t.Hidden {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(t.Spr.X + float64(t.PadLeft + t.ShadowX), t.Spr.Y-float64(t.Font.Metrics().Height.Ceil()/8+t.Font.Metrics().Height.Ceil()+t.PadUp + t.ShadowY))
		op.ColorM.Scale(colorToScale(t.ShadowClr))
		text.DrawWithOptions(screen, t.Txt, t.Font, op)

		op.GeoM.Translate(t.Spr.X + float64(t.PadLeft), t.Spr.Y-float64(t.Font.Metrics().Height.Ceil()/8+t.Font.Metrics().Height.Ceil()+t.PadUp))
		op.ColorM.Scale(colorToScale(t.Clr))
		text.DrawWithOptions(screen, t.Txt, t.Font, op)

		// t.Spr.Draw(screen)
		// text.Draw(screen, t.Txt, t.Font, int(t.Spr.X) + t.PadLeft + t.ShadowX, int(t.Spr.Y)-t.Font.Metrics().Height.Ceil()/8+t.Font.Metrics().Height.Ceil()+t.PadUp + t.ShadowY, t.ShadowClr)
		// text.Draw(screen, t.Txt, t.Font, int(t.Spr.X) + t.PadLeft, int(t.Spr.Y)-t.Font.Metrics().Height.Ceil()/8+t.Font.Metrics().Height.Ceil()+t.PadUp, t.Clr)
		// text.Draw(t.Spr.Img, t.Txt, t.Font, t.PadLeft + t.ShadowX, -t.Font.Metrics().Height.Ceil()/8+t.Font.Metrics().Height.Ceil()+t.PadUp + t.ShadowY, t.ShadowClr)
		// text.Draw(t.Spr.Img, t.Txt, t.Font, t.PadLeft, -t.Font.Metrics().Height.Ceil()/8+t.Font.Metrics().Height.Ceil()+t.PadUp, t.Clr)
		// t.Spr.Draw(screen)
	}
}