package screen

import (
	"image/color"
	"starRailTCG/common"
	"starRailTCG/enums"

	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
)

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

func NewOnBoardScreen() *ebitenui.UI {

	buttonImage, _ := loadButtonImage()

	face, _ := loadFont(24)

	// construct a new container that serves as the root of the UI hierarchy
	rootContainer := widget.NewContainer(
		// the container will use a plain color as its background
		// widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(color.NRGBA{0x13, 0x1a, 0x22, 0xff})),

		// the container will use an anchor layout to layout its single child widget
		widget.ContainerOpts.Layout(widget.NewAnchorLayout(
			widget.AnchorLayoutOpts.Padding(widget.NewInsetsSimple(20)),
		)),
	)

	buttonContainer := widget.NewContainer(
		widget.ContainerOpts.Layout(
			widget.NewGridLayout(
				widget.GridLayoutOpts.Columns(1),
				//Define how much padding to inset the child content
				widget.GridLayoutOpts.Padding(widget.NewInsetsSimple(30)),
				//Define how far apart the rows and columns should be
				widget.GridLayoutOpts.Spacing(20, 10),
				//Define how to stretch the rows and columns. Note it is required to
				//specify the Stretch for each row and column.
				// widget.GridLayoutOpts.Stretch([]bool{true, false}, []bool{false, true}),
			)),
	)

	// 开始游戏按钮 start game button
	startGameBtn := widget.NewButton(
		widget.ButtonOpts.WidgetOpts(
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
			common.ChangeScreen(enums.ScreenGameMode)
		}),
	)
	// fmt.Println(startGameBtn)
	// add the button as a child of the container
	buttonContainer.AddChild(startGameBtn)
	// rootContainer.AddChild(startGameBtn)

	// 设置按钮
	settingBtn := widget.NewButton(
		widget.ButtonOpts.WidgetOpts(
			widget.WidgetOpts.MinSize(200, 80),
		),
		// specify the images to use
		widget.ButtonOpts.Image(buttonImage),
		widget.ButtonOpts.Text("SETTING", face, &widget.ButtonTextColor{
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
			// TODO
			println("setting btn clicked")
			// ChangeScreen(ScreenGameMode)
		}),
	)
	buttonContainer.AddChild(settingBtn)
	rootContainer.AddChild(buttonContainer)

	// construct the UI
	ui := &ebitenui.UI{
		Container: rootContainer,
	}

	return ui
}
