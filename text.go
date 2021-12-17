package urushi

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
)

type TxtSpr struct {
	Txt     string
	Spr *Sprite
	Clr     color.Color
	PadUp   int
	PadLeft int
	Font    font.Face
	IsVert  bool
	Hidden bool
}

func NewTxtSpr(txt string, x, y float64, clr color.Color, font font.Face, padUp, padLeft int, hidden bool) *TxtSpr {
	var bgImg *ebiten.Image

	width := text.BoundString(font, txt).Dx()
	if width+padLeft*2 == 0 {
		width = 1
	}
	bgImg = ebiten.NewImage(width+padLeft*2, font.Metrics().Height.Ceil()+padUp*2)
	bgImg.Fill(color.White)
	// bgImg.Fill(bgClr)	

	t := &TxtSpr{
		Txt: txt,
		Spr: NewSprite(bgImg, x, y),
		Clr: clr,
		Font: font,
		PadUp: padUp,
		PadLeft: padLeft,
		Hidden: hidden,
	}
	return t
}

func (t *TxtSpr) SetCenter(center int) {
	if t.IsVert {
		t.Spr.X = float64(center - t.Font.Metrics().Height.Ceil()/2 - t.PadLeft)
	} else {
		width := text.BoundString(t.Font, t.Txt).Dx()
		t.Spr.X = float64(center - width/2 - t.PadLeft)
	}
}

func (t *TxtSpr) SetText(txt string) {
	t.Txt = txt
	width := text.BoundString(t.Font, txt).Dx()
	if width+t.PadLeft*2 == 0 {
		width = 1
	}
	t.Spr.Img = ebiten.NewImage(width+t.PadLeft*2, t.Font.Metrics().Height.Ceil()+t.PadUp*2)
}

func (t *TxtSpr) Draw(screen *ebiten.Image) {
	if !t.Hidden {
		if t.IsVert {
			t.Spr.Draw(screen)
		
			for i, v := range []rune(t.Txt) {
				if string(v) == "、" || string(v) == "。" {
					// Perhaps int(t.spr.x) is not good?
					text.Draw(screen, string(v), t.Font, int(t.Spr.X)+t.PadLeft+t.Font.Metrics().Height.Ceil()-text.BoundString(t.Font, string(v)).Dx(), int(t.Spr.Y)-t.Font.Metrics().Height.Ceil()/8+t.PadUp+i*t.Font.Metrics().Height.Ceil()+t.Font.Metrics().Height.Ceil()-text.BoundString(t.Font, string(v)).Dy(), t.Clr)
				} else {
					text.Draw(screen, string(v), t.Font, int(t.Spr.X)+t.PadLeft, int(t.Spr.Y)-t.Font.Metrics().Height.Ceil()/8+t.PadUp+t.Font.Metrics().Height.Ceil()+i*t.Font.Metrics().Height.Ceil(), t.Clr)
	
				}
			}
		} else {
			t.Spr.Draw(screen)

			op := &ebiten.DrawImageOptions{}
			// op.GeoM.Translate(t.Spr.X + float64(t.PadLeft), t.Spr.Y+float64(t.PadUp))
			op.GeoM.Translate(t.Spr.X + float64(t.PadLeft), t.Spr.Y+float64(t.Font.Metrics().Height.Ceil()*6/8 +1 + t.PadUp))

			// op.GeoM.Translate(t.Spr.X + float64(t.PadLeft), t.Spr.Y+float64(-t.Font.Metrics().Height.Ceil() + t.PadUp))
			op.ColorM.Scale(colorToScale(t.Clr))
			text.DrawWithOptions(screen, t.Txt, t.Font, op)

			// text.DrawWithOptions(screen, t.Txt, t.Font, int(t.Spr.X)+t.PadLeft, int(t.Spr.Y)-t.Font.Metrics().Height.Ceil()/8+t.Font.Metrics().Height.Ceil()+t.PadUp, op)
			// text.Draw(t.Spr.Img, t.Txt, t.Font, t.PadLeft, -t.Font.Metrics().Height.Ceil()/8+t.Font.Metrics().Height.Ceil()+t.PadUp, t.Clr)
			// t.Spr.Draw(screen)
		}
	}
}

func colorToScale(clr color.Color) (float64, float64, float64, float64) {
	r, g, b, a := clr.RGBA()
	return float64(uint8(r)) / 255, float64(uint8(g)) / 255, float64(uint8(b)) / 255, float64(uint8(a)) / 255
}