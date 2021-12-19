package urushi

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

type TxtSpr struct {
	Txt     string
	Spr *Sprite
	Clr     color.Color
	PadUp   int
	PadLeft int
	Font    font.Face
	isVert  bool
	Hidden bool
}

func NewTxtSpr(txt string, x, y float64, clr color.Color, font font.Face, padUp, padLeft int, hidden bool) *TxtSpr {
	var bgImg *ebiten.Image

	width := text.BoundString(font, txt).Dx()
	if width+padLeft*2 == 0 {
		width = 1
	}
	bgImg = ebiten.NewImage(width+padLeft*2, -text.BoundString(font, txt).Bounds().Min.Y + text.BoundString(font, txt).Bounds().Max.Y+padUp*2)
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
	if t.isVert {
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

func (t *TxtSpr) SetVertical(vertical bool) {
	height := 0
	height += t.Font.Metrics().Height.Ceil() * len([]rune(t.Txt))
	// for _, v := range []rune(t.Txt) {
	// 	switch string(v) {
	// 	case "ー":
	// 		v = rune("|")
	// 	case "　":
	// 		v = rune(" ")
	// 	}

	// 	height += t.Font.Metrics().Height.Ceil() * len([]rune(t.Txt))

	// }
	if height+t.PadUp*2 == 0 {
		height = 1
	}

	t.Spr.Img = ebiten.NewImage(t.Font.Metrics().Height.Ceil()+t.PadLeft*2, height+t.PadUp*2)

	t.isVert = true
}

func (t *TxtSpr) Draw(screen *ebiten.Image) {
	if !t.Hidden {
		if t.isVert {
			t.Spr.Draw(screen)
			yPos := 0.0
			for _, v := range []rune(t.Txt) {
				s := string(v)
				if s == "ー" {
					s = "|"
				}

				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(t.Spr.X - float64((t.Spr.Img.Bounds().Dx()-text.BoundString(t.Font, s).Bounds().Min.X) / 2 + t.PadLeft), t.Spr.Y+ yPos + float64(-text.BoundString(t.Font, s).Bounds().Min.Y + t.PadUp))
				op.ColorM.Scale(colorToScale(t.Clr))
				text.DrawWithOptions(screen, s, t.Font, op)

				// yPos += float64(-text.BoundString(t.Font, s).Bounds().Min.Y + text.BoundString(t.Font, s).Bounds().Max.Y)
				yPos += float64(t.Font.Metrics().Height.Ceil())
			}

		} else {
			t.Spr.Draw(screen)
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(t.Spr.X + float64(-text.BoundString(t.Font, t.Txt).Bounds().Min.X + t.PadLeft), t.Spr.Y+float64(-text.BoundString(t.Font, t.Txt).Bounds().Min.Y + t.PadUp))
			op.ColorM.Scale(colorToScale(t.Clr))
			text.DrawWithOptions(screen, t.Txt, t.Font, op)
		}
	}
}

func colorToScale(clr color.Color) (float64, float64, float64, float64) {
	r, g, b, a := clr.RGBA()
	return float64(uint8(r)) / 255, float64(uint8(g)) / 255, float64(uint8(b)) / 255, float64(uint8(a)) / 255
}