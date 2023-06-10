package widgets

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/yohamta/furex/v2"
	"golang.org/x/image/font"
)

type Button struct {
	Color    color.Color
	attrs    map[string]string
	OnClick  func(attrs map[string]string)
	FontFace font.Face

	mouseover bool
	pressed   bool
}

func (b *Button) HandlePress(x, y int, t ebiten.TouchID) {
	b.pressed = true
}

func (b *Button) HandleRelease(x, y int, isCancel bool) {
	b.pressed = false
	if !isCancel {
		if b.OnClick != nil {
			b.OnClick(b.attrs)
		}
	}
}

func (b *Button) HandleMouseEnter(x, y int) bool {
	b.mouseover = true
	return true
}

func (b *Button) HandleMouseLeave() {
	b.mouseover = false
}

func (b *Button) Draw(screen *ebiten.Image, frame image.Rectangle, view *furex.View) {
	x, y := float64(frame.Min.X+frame.Dx()/2), float64(frame.Min.Y+frame.Dy()/2)
	b.attrs = view.Attrs
	// todo
	// sprite := view.Attrs["sprite"]
	// spritePressed := view.Attrs["sprite_pressed"]

	// opts := ganim8.DrawOpts(x, y, 0, 1, 1, .5, .5)
	// if b.mouseover {
	// 	opts.ColorM.Scale(1.1, 1.1, 1.1, 1)
	// }

	buttonColor := color.NRGBA{R: 170, G: 170, B: 180, A: 255}
	buttonHover := color.NRGBA{R: 130, G: 130, B: 150, A: 255}
	buttonPressed := color.NRGBA{R: 100, G: 100, B: 120, A: 255}

	// vector.DrawFilledRect(screen, float32(frame.Min.X), float32(frame.Min.Y), float32(view.Width), float32(view.Height), buttonColor, false)

	if b.pressed {
		vector.DrawFilledRect(screen, float32(frame.Min.X), float32(frame.Min.Y), float32(view.Width), float32(view.Height), buttonPressed, false)
	} else if b.mouseover {
		vector.DrawFilledRect(screen, float32(frame.Min.X), float32(frame.Min.Y), float32(view.Width), float32(view.Height), buttonHover, false)
	} else {
		vector.DrawFilledRect(screen, float32(frame.Min.X), float32(frame.Min.Y), float32(view.Width), float32(view.Height), buttonColor, false)
	}

	if b.Color == nil {
		b.Color = color.White
	}
	textBox := text.BoundString(b.FontFace, view.Text)
	minW := textBox.Dx() / 2
	minH := textBox.Dy() / 2
	text.Draw(screen, view.Text, b.FontFace, int(x)-minW, int(y)+minH, b.Color)
}
