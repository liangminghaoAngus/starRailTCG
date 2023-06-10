package screen

import (
	"image/color"
	"starRailTCG/widgets"

	e_image "github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/golang/freetype/truetype"
	"github.com/yohamta/furex/v2"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"

	_ "embed"
)

//go:embed html/onBoard.html
var onBoardHtml string

func loadButtonImage() (*widget.ButtonImage, error) {

	idle := e_image.NewNineSliceColor(color.NRGBA{R: 170, G: 170, B: 180, A: 255})

	hover := e_image.NewNineSliceColor(color.NRGBA{R: 130, G: 130, B: 150, A: 255})

	pressed := e_image.NewNineSliceColor(color.NRGBA{R: 100, G: 100, B: 120, A: 255})

	return &widget.ButtonImage{
		Idle:    idle,
		Hover:   hover,
		Pressed: pressed,
	}, nil
}

func loadFont(size float64) (font.Face, error) {
	ttfFont, err := truetype.Parse(goregular.TTF)
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

	// face, _ := loadFont(24)
	view := furex.Parse(onBoardHtml, &furex.ParseOptions{
		Width:  1280,
		Height: 720,
		Components: map[string]furex.Component{
			"button": &widgets.Button{
				OnClick: func() {
					println("button click")
				},
			},
		},
	})

	return view
}
