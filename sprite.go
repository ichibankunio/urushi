package urushi

import (
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2"
)

type Sprite struct {
	Img *ebiten.Image
	X float64
	Y float64
	Vx float64
	Vy float64
	Hidden bool
	Alpha float64
	DrawOption func(*ebiten.DrawImageOptions)
	
	TouchID ebiten.TouchID
}

func NewSprite(img *ebiten.Image, x, y float64) *Sprite {
	spr := &Sprite{
		Img:   img,
		X:     x,
		Y:     y,
		Alpha: 1,
		DrawOption: func(*ebiten.DrawImageOptions){},
	}

	return spr
}

func (s *Sprite) SetCenter(center int) *Sprite {
	s.X = float64(center - s.Img.Bounds().Dx() / 2)
	return s
}

func (s *Sprite) Draw(screen *ebiten.Image) {
	if !s.Hidden {
		op := &ebiten.DrawImageOptions{}
		s.DrawOption(op)
	
		op.GeoM.Translate(s.X, s.Y)
		op.ColorM.Scale(1, 1, 1, s.Alpha)
		screen.DrawImage(s.Img, op)
	}
}


func (s *Sprite) IsJustTouched() bool {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) || inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		x, y := ebiten.CursorPosition()

		if x >= int(s.X) && x <= int(s.X)+s.Img.Bounds().Dx() && y >= int(s.Y) && y <= int(s.Y)+s.Img.Bounds().Dy() {
			return true
		}
	}

	// touch := inpututil.JustPressedTouchIDs()
	touch := inpututil.AppendJustPressedTouchIDs(nil)
	if len(touch) > 0 {
		for _, t := range touch {
			x, y := ebiten.TouchPosition(t)
			if x >= int(s.X) && x <= int(s.X)+s.Img.Bounds().Dx() && y >= int(s.Y) && y <= int(s.Y)+s.Img.Bounds().Dy() {
				s.touchID = t
				return true
			}
		}
		
	}

	return false
}

func (s *Sprite) IsTouchJustReleased() bool {
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()

		if x >= int(s.X) && x <= int(s.X)+s.Img.Bounds().Dx() && y >= int(s.Y) && y <= int(s.Y)+s.Img.Bounds().Dy() {
			return true
		}
	}

	if inpututil.IsTouchJustReleased(s.touchID) {
		return true
	}

	return false
}

func (s *Sprite) IsTouched() bool {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()

		if x >= int(s.X) && x <= int(s.X)+s.Img.Bounds().Dx() && y >= int(s.Y) && y <= int(s.Y)+s.Img.Bounds().Dy() {
			return true
		}
	}

	t := ebiten.TouchIDs()
	if len(t) > 0 {
		x, y := ebiten.TouchPosition(t[0])
		if x >= int(s.X) && x <= int(s.X)+s.Img.Bounds().Dx() && y >= int(s.Y) && y <= int(s.Y)+s.Img.Bounds().Dy() {
			return true
		}
	}

	return false
}



