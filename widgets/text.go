package widgets

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/tinne26/etxt"
	"github.com/yohamta/furex/v2"
	"golang.org/x/image/font"
)

var (
	R    *etxt.Renderer
	Font = "x14y20pxScoreDozer"
)

type Text struct {
	Color     color.Color
	Shadow    bool
	HorzAlign etxt.HorzAlign
	VertAlign etxt.VertAlign
	Text      string

	FontFace font.Face
}

var (
	_ furex.Drawer = (*Text)(nil)
)

func (t *Text) Draw(screen *ebiten.Image, frame image.Rectangle, view *furex.View) {
	if t.Shadow {
		vector.DrawFilledRect(
			screen, float32(frame.Min.X), float32(frame.Min.Y), float32(len(view.Text)*6+4), float32(frame.Dy()), color.RGBA{0, 0, 0, 50}, false)
	}
	x, y := frame.Min.X+frame.Dx()/2, frame.Min.Y+frame.Dy()/2
	if t.HorzAlign == etxt.Left {
		x = frame.Min.X
	}
	if t.VertAlign == etxt.Top {
		y = frame.Min.Y
	}
	if t.Color != nil {
		// text.R.SetColor(t.Color)
	} else {
		t.Color = color.White
	}

	textStr := ""
	if t.Text == "" {
		textStr = view.Text
	} else {
		textStr = t.Text
	}
	textBox := text.BoundString(t.FontFace, textStr)
	minW := textBox.Dx() / 2
	minH := textBox.Dy() / 2
	text.Draw(screen, textStr, t.FontFace, x-minW, y+minH, t.Color)
}
