package urushi

import (
	"bytes"
	"image/png"
	"log"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"golang.org/x/image/font"
)

func NewImageFromBytes(byteData []byte) *ebiten.Image {
	r := bytes.NewReader(byteData)
	p, _ := png.Decode(r)
	return ebiten.NewImageFromImage(p)
}

func NewFontFromBytes(byteData []byte, size int) font.Face {
	tt, err := truetype.Parse(byteData)
	if err != nil {
		log.Fatal(err)
	}

	return truetype.NewFace(tt, &truetype.Options{
		Size:    float64(size),
		DPI:     72,
		Hinting: font.HintingFull,
	})
}

func NewBGMFromBytes(b []byte, sampleRate int, context *audio.Context) *audio.Player {
	m, _ := mp3.DecodeWithSampleRate(sampleRate, bytes.NewReader(b))
	s := audio.NewInfiniteLoop(m, m.Length())
	p, _ := context.NewPlayer(s)
	// p, _ := audio.NewPlayer(audioContext, s)

	return p
}
