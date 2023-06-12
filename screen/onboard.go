package screen

import (
	"image"
	"image/color"
	"image/jpeg"
	"os"
	"starRailTCG/common"
	"starRailTCG/enums"
	"starRailTCG/widgets"
	"strconv"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
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
	fontFile, err := os.ReadFile("static/AlibabaPuHuiTi-3-75-SemiBold.ttf")
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
	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
	logo, logoFile, _ := ebitenutil.NewImageFromFile("./static/imgs/logo.png")
	bgImg, _, _ := ebitenutil.NewImageFromFile("./static/imgs/bgImg.jpeg")
	rootView := furex.NewHandler(furex.HandlerOpts{
		HandleRelease: func(x, y int, isCancel bool) {},
		HandlePress:   func(x, y int, t ebiten.TouchID) {},
		Update:        func(v *furex.View) {},
		Draw: func(screen *ebiten.Image, frame image.Rectangle, v *furex.View) {
			// screen.Fill(Hex2RGB("#0dceda", 0))
			// sw, sh := frame.Min.X+frame.Dx()/2, frame.Min.Y+frame.Dy()/2
			opt := &ebiten.DrawImageOptions{}
			opt.GeoM.Scale(0.75, 0.75)
			scale := ebiten.DeviceScaleFactor()
			opt.GeoM.Scale(scale, scale)
			// opt.GeoM.Translate(float64(sw-(bgImgFile.Bounds().Dx()/2)), float64(sh-(bgImgFile.Bounds().Dy()/2)-10))
			opt.Filter = ebiten.FilterLinear

			screen.DrawImage(bgImg, opt)
		},
	})

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
							if id, ok := attrs["id"]; ok {
								ButtonClickCall(id)
							}

							// fmt.Printf("%+v", attrs)
							// println("button click")
						},
					},
				}
			},
			"onboard-title": func() *furex.View {
				return &furex.View{
					Height:     65 * int(scale),
					Width:      400 * int(scale),
					Direction:  furex.Row,
					AlignItems: furex.AlignItemCenter,
					Justify:    furex.JustifyStart,
					Handler:    &widgets.Text{FontFace: titleFace},
				}
			},
			"title-panel-img": func() *furex.View {
				return &furex.View{
					Height:     100 * int(scale),
					Width:      400 * int(scale),
					Direction:  furex.Row,
					AlignItems: furex.AlignItemCenter,
					Justify:    furex.JustifyStart,
					Handler:    &widgets.Image{Img: logo, Item: logoFile.Bounds()},
				}
			},
		},
		Handler: rootView,
	})

	return view
}

func Hex2RGB(color16 string, alpha uint8) color.RGBA {
	r, _ := strconv.ParseInt(color16[:2], 16, 10)
	g, _ := strconv.ParseInt(color16[2:4], 16, 18)
	b, _ := strconv.ParseInt(color16[4:], 16, 10)
	return color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: alpha}
}

func ButtonClickCall(id string) {
	const (
		startGame = "startGame"
		settings  = "settings"
		exitGame  = "exitGame"
	)
	switch id {
	case startGame:
		common.ChangeScreen(enums.ScreenGameMode)
	case settings:

	case exitGame:
		os.Exit(1)
	}
}
