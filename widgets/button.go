package widgets

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/furex/v2"
)

type Button struct {
	Color   color.Color
	OnClick func()

	mouserover bool
	pressed    bool
}

func (b *Button) HandlePress(x, y int, t ebiten.TouchID) {
	b.pressed = true
}

func (b *Button) HandleRelease(x, y int, isCancel bool) {
	b.pressed = false
	if !isCancel {
		if b.OnClick != nil {
			b.OnClick()
		}
	}
}

func (b *Button) HandleMouseEnter(x, y int) bool {
	b.mouserover = true
	return true
}

func (b *Button) HandleMouseLeave() {
	b.mouserover = false
}

func (b *Button) Draw(screen *ebiten.Image, frame image.Rectangle, view *furex.View) {
	// x, y := float64(frame.Min.X+frame.Dx()/2), float64(frame.Min.Y+frame.Dy()/2)

	// todo

}
