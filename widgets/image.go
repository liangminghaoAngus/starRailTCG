package widgets

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/furex/v2"
)

type Image struct {
	Img     *ebiten.Image
	Item    image.Rectangle
	OnClick func(attr map[string]string)

	Attrs   map[string]string
	pressed bool
	// attrs map[string]string
}

func (i *Image) HandlePress(x, y int, t ebiten.TouchID) {
	i.pressed = true
}

func (i *Image) HandleRelease(x, y int, isCancel bool) {
	i.pressed = false
	if !isCancel {
		if i.OnClick != nil {
			i.OnClick(i.Attrs)
		}
	}
}

func (i *Image) Draw(screen *ebiten.Image, frame image.Rectangle, view *furex.View) {
	if i.Img == nil {
		return
	}

	// sw, sh := frame.Bounds().Dx(), frame.Bounds().Dy()
	sw, sh := frame.Min.X+frame.Dx()/2, frame.Min.Y+frame.Dy()/2
	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Scale(0.6, 0.6)
	scale := ebiten.DeviceScaleFactor()
	opt.GeoM.Scale(scale, scale)
	opt.GeoM.Translate(float64(sw-(i.Item.Dx()/2)), float64(sh-(i.Item.Bounds().Dy()/2)-10))
	opt.Filter = ebiten.FilterLinear

	screen.DrawImage(i.Img, opt)
}
