package screen

import (
	"image"

	_ "embed"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/furex/v2"
)

//go:embed html/onGame.html
var GameHtml string

func NewGameScreen() *furex.View {
	furex.Debug = true

	scale := ebiten.DeviceScaleFactor()
	// face, _ := loadFont(24 * scale)

	view := furex.Parse(GameHtml, &furex.ParseOptions{
		Width:      1280 * int(scale),
		Height:     720 * int(scale),
		Components: map[string]furex.Component{},
		Handler: furex.NewHandler(furex.HandlerOpts{
			Draw: func(screen *ebiten.Image, frame image.Rectangle, v *furex.View) {
				println("game screen")
			},
		}),
	})

	return view
}
