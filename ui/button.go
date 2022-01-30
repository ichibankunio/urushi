package ui

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/ichibankunio/urushi"
	"golang.org/x/image/font"
)

type UITheme int

const (
	ThemeRect UITheme = iota
	ThemeRound
)

type Button struct {
	spr *urushi.Sprite
	txt *urushi.TxtSpr
}

func NewButton(txt string, centerX int, centerY int, width int, height int, fontface font.Face, theme UITheme) *Button {
	return &Button{
		spr: urushi.NewSprite(newButtonImg(width, height, theme), float64(centerX - width / 2), float64(centerY - height / 2)),
		txt: urushi.NewTxtSpr(txt, float64(centerX - text.BoundString(fontface, txt).Dx()/2), float64(centerY - text.BoundString(fontface, txt).Dy()/2), color.Black, fontface, 0, 0, false),
	}
}

func (b *Button) Draw(screen *ebiten.Image) {
	b.spr.Draw(screen)
	b.txt.Draw(screen)
}

func newButtonImg(width, height int, theme UITheme) *ebiten.Image {
	bg := ebiten.NewImage(width, height)
	bg.Fill(color.White)
	src := ebiten.NewImage(1, 1)
	src.Fill(color.Black)

	var path vector.Path

	w := float32(width)
	h := float32(height)
	l := float32(math.Min(float64(width), float64(height)) / 16)

	switch theme {
	case ThemeRect:
		path.MoveTo(0, 0)
		path.LineTo(float32(width), 0)
		path.LineTo(float32(width), float32(height))
		path.LineTo(0, float32(height))
		path.LineTo(0, 0)
		path.MoveTo(l, l)
		path.LineTo(w - l, l)
		path.LineTo(w - l, h - l)
		path.LineTo(l, h - l)
		path.LineTo(l, l)

	case ThemeRound:
		path.MoveTo(l, 0)
		path.LineTo(w - l, 0)
		path.ArcTo(w-l, 0, w, l, l)
		path.LineTo(w, h-l)
		path.ArcTo(w, h-l, w-l, h, l)
		path.LineTo(l, h)
		path.ArcTo(l, h, 0, h-l, l)
		path.LineTo(0, l)
		path.ArcTo(0, l, l, 0, l)
		
	}

	
	op := &ebiten.DrawTrianglesOptions{
		FillRule: ebiten.EvenOdd,
	}

	vs, is := path.AppendVerticesAndIndicesForFilling(nil, nil)
	bg.DrawTriangles(vs, is, src, op)
	
	return bg
}