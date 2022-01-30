package ui

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/ichibankunio/urushi"
	"golang.org/x/image/font"
)

type Button struct {
	spr *urushi.Sprite
	txt *urushi.TxtSpr
}

func NewButton(txt string, centerX int, centerY int, width int, height int, fontface font.Face) *Button {
	return &Button{
		spr: urushi.NewSprite(newButtonImg(width, height), float64(centerX - width / 2), float64(centerY - height / 2)),
		txt: urushi.NewTxtSpr(txt, float64(centerX - width / 2), float64(centerY - height / 2), color.Black, fontface, 0, 0, false),
	}
}

func (b *Button) Draw(screen *ebiten.Image) {
	b.spr.Draw(screen)
	b.txt.Draw(screen)
}

func newButtonImg(width, height int) *ebiten.Image {
	bg := ebiten.NewImage(width, height)
	bg.Fill(color.White)
	src := ebiten.NewImage(1, 1)
	src.Fill(color.Black)
	line := width / 8
	var path vector.Path

	path.MoveTo(0, 0)
	path.LineTo(float32(width), 0)
	path.LineTo(float32(width), float32(height))
	path.LineTo(0, float32(height))
	path.LineTo(0, 0)
	path.MoveTo(float32(line), float32(line))
	path.LineTo(float32(width - line), float32(line))
	path.LineTo(float32(width - line), float32(height - line))
	path.LineTo(float32(line), float32(height - line))
	path.LineTo(float32(line), float32(line))
	
	
	op := &ebiten.DrawTrianglesOptions{
		FillRule: ebiten.EvenOdd,
	}

	vs, is := path.AppendVerticesAndIndicesForFilling(nil, nil)
	bg.DrawTriangles(vs, is, src, op)
	
	return bg
}