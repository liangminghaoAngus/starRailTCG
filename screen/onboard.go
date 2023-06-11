package screen

import (
	"fmt"
	"io/ioutil"
	"starRailTCG/widgets"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/furex/v2"
	"golang.org/x/image/font"

	_ "embed"
)

//go:embed html/onBoard.html
var onBoardHtml string

// func loadButtonImage() (*widget.ButtonImage, error) {

// 	idle := e_image.NewNineSliceColor(color.NRGBA{R: 170, G: 170, B: 180, A: 255})

// 	hover := e_image.NewNineSliceColor(color.NRGBA{R: 130, G: 130, B: 150, A: 255})

// 	pressed := e_image.NewNineSliceColor(color.NRGBA{R: 100, G: 100, B: 120, A: 255})

// 	return &widget.ButtonImage{
// 		Idle:    idle,
// 		Hover:   hover,
// 		Pressed: pressed,
// 	}, nil
// }

func loadFont(size float64) (font.Face, error) {
	fontFile, err := ioutil.ReadFile("static/AlibabaPuHuiTi-3-75-SemiBold.ttf")
	if err != nil {
		panic("font file err")
	}
	ttfFont, err := truetype.Parse(fontFile)
	if err != nil {
		return nil, err
	}

	return truetype.NewFace(ttfFont, &truetype.Options{
		Size:    size,
		DPI:     72,
		Hinting: font.HintingFull,
	}), nil
}

func NewOnBoardScreen() *furex.View {

	furex.Debug = true

	// buttonImage, _ := loadButtonImage()

	scale := ebiten.DeviceScaleFactor()
	face, _ := loadFont(24 * scale)
	titleFace, _ := loadFont(50 * scale)
	view := furex.Parse(onBoardHtml, &furex.ParseOptions{
		Width:  1280 * int(scale),
		Height: 720 * int(scale),
		Components: map[string]furex.Component{
			"button": func() *furex.View {
				return &furex.View{
					Height: 50 * int(scale),
					Width:  400 * int(scale),
					Handler: &widgets.Button{
						FontFace: face,
						OnClick: func(attrs map[string]string) {
							fmt.Printf("%+v", attrs)
							println("button click")
						},
					},
				}
			},
			"onboard-title": func() *furex.View {
				return &furex.View{
					Height:     50 * int(scale),
					Direction:  furex.Row,
					AlignItems: furex.AlignItemCenter,
					Justify:    furex.JustifyStart,
					Handler:    &widgets.Text{FontFace: titleFace},
				}
			},
		},
	})

	return view
}
