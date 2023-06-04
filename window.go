package main

import (
	"fmt"
	"image/color"
	_ "image/gif"
	"log"
	"starRailTCG/config"
	"strconv"

	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
)

type Game struct {
	i             uint8
	cfg           *config.Config
	OnboardScreen *ebitenui.UI
}

func NewGame(cfg *config.Config) *Game {

	ebiten.SetWindowSize(cfg.ScreenWidth, cfg.ScreenHeight) //窗口大小
	ebiten.SetWindowTitle(cfg.Title)                        //窗口标题

	onboard := NewOnBoardScreen()
	return &Game{
		cfg:           cfg,
		OnboardScreen: onboard,
	}
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// g.i++
	// if g.i < 255 {
	// 	screen.Fill(Hex2RGB("#0dceda", g.i))
	// } else {
	// 	g.i = 0
	// }

	screen.Fill(Hex2RGB("#0dceda", g.i))
	sw, sh := screen.Bounds().Dx(), screen.Bounds().Dy()
	// 加载背景图片
	bgGIF, bgGIFFile, err := ebitenutil.NewImageFromFile("./static/imgs/onboard.gif")
	if err != nil {
		log.Fatalf("ebitenutil.NewImageFromFile failed :%v", err)
	}
	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Scale(0.6, 0.6)
	// opt.GeoM.Translate(float64(g.cfg.ScreenWidth-bgGIFFile.Bounds().Dx())/2, float64(g.cfg.ScreenHeight-bgGIFFile.Bounds().Dy())/2)
	scale := ebiten.DeviceScaleFactor()
	// log.Printf("scale :%+v", scale)
	opt.GeoM.Scale(scale, scale)
	opt.GeoM.Translate(float64(sw-bgGIFFile.Bounds().Dx()-40)/2, float64(sh-bgGIFFile.Bounds().Dy()-40)/2)
	opt.Filter = ebiten.FilterLinear

	screen.DrawImage(bgGIF, opt)
	g.OnboardScreen.Draw(screen)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f\nFPS: %0.2f", ebiten.ActualTPS(), ebiten.ActualFPS()))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	s := ebiten.DeviceScaleFactor()
	return int(float64(outsideWidth) * s), int(float64(outsideHeight) * s)
}

func Hex2RGB(color16 string, alpha uint8) color.RGBA {
	r, _ := strconv.ParseInt(color16[:2], 16, 10)
	g, _ := strconv.ParseInt(color16[2:4], 16, 18)
	b, _ := strconv.ParseInt(color16[4:], 16, 10)
	return color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: alpha}
}

func NewOnBoardScreen() *ebitenui.UI {

	buttonImage, _ := loadButtonImage()

	face, _ := loadFont(24)

	// construct a new container that serves as the root of the UI hierarchy
	rootContainer := widget.NewContainer(
		// the container will use a plain color as its background
		// widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(color.NRGBA{0x13, 0x1a, 0x22, 0xff})),

		// the container will use an anchor layout to layout its single child widget
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
	)

	// 开始游戏按钮 start game button
	startGameBtn := widget.NewButton(
		widget.ButtonOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				HorizontalPosition: widget.AnchorLayoutPositionCenter,
				VerticalPosition:   widget.AnchorLayoutPositionEnd,
			}),
			widget.WidgetOpts.MinSize(200, 80),
		),
		// specify the images to use
		widget.ButtonOpts.Image(buttonImage),
		widget.ButtonOpts.Text("START", face, &widget.ButtonTextColor{
			Idle: color.RGBA{0xdf, 0xf4, 0xff, 0xff},
		}),
		// specify that the button's text needs some padding for correct display
		widget.ButtonOpts.TextPadding(widget.Insets{
			Left:   30,
			Right:  30,
			Top:    5,
			Bottom: 5,
		}),
		// add a handler that reacts to clicking the button
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			println("button clicked")
		}),
	)
	// add the button as a child of the container
	rootContainer.AddChild(startGameBtn)

	// construct the UI
	ui := &ebitenui.UI{
		Container: rootContainer,
	}

	return ui
}

func loadButtonImage() (*widget.ButtonImage, error) {
	idle := image.NewNineSliceColor(color.NRGBA{R: 170, G: 170, B: 180, A: 255})

	hover := image.NewNineSliceColor(color.NRGBA{R: 130, G: 130, B: 150, A: 255})

	pressed := image.NewNineSliceColor(color.NRGBA{R: 100, G: 100, B: 120, A: 255})

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
