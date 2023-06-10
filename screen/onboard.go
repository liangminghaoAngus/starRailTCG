package screen

import (
	"fmt"
	"image"
	"image/color"
	"math"
	"starRailTCG/common"
	"starRailTCG/enums"

	"github.com/ebitenui/ebitenui"
	e_image "github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
)

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

func NewOnBoardScreen() *ebitenui.UI {

	ui := &ebitenui.UI{}

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
	// add the button as a child of the container
	buttonContainer.AddChild(startGameBtn)

	// 设置菜单弹窗
	// Create the titlebar for the window
	titleContainer := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(e_image.NewNineSliceColor(color.NRGBA{150, 150, 150, 255})),
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
	)
	titleContainer.AddChild(widget.NewText(
		widget.TextOpts.Text("Window Title", face, color.NRGBA{254, 255, 255, 255}),
		widget.TextOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
			HorizontalPosition: widget.AnchorLayoutPositionCenter,
			VerticalPosition:   widget.AnchorLayoutPositionCenter,
		})),
	))

	windowContainer := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(e_image.NewNineSliceColor(color.NRGBA{100, 100, 100, 255})),
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
	)
	windowContainer.AddChild(widget.NewText(
		widget.TextOpts.Text("Hello from window", face, color.NRGBA{254, 255, 255, 255}),
		widget.TextOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
			HorizontalPosition: widget.AnchorLayoutPositionCenter,
			VerticalPosition:   widget.AnchorLayoutPositionCenter,
		})),
	))
	window := widget.NewWindow(
		//Set the main contents of the window
		widget.WindowOpts.Contents(windowContainer),
		//Set the titlebar for the window (Optional)
		widget.WindowOpts.TitleBar(titleContainer, 25),
		//Set the window above everything else and block input elsewhere
		widget.WindowOpts.Modal(),
		//Set how to close the window. CLICK_OUT will close the window when clicking anywhere
		//that is not a part of the window object
		widget.WindowOpts.CloseMode(widget.CLICK_OUT),
		//Indicates that the window is draggable. It must have a TitleBar for this to work
		widget.WindowOpts.Draggable(),
		//Set the window resizeable
		widget.WindowOpts.Resizeable(),
		//Set the minimum size the window can be
		widget.WindowOpts.MinSize(200, 200),
		//Set the maximum size a window can be
		// widget.WindowOpts.MaxSize(300, 300),
		//Set the callback that triggers when a move is complete
		widget.WindowOpts.MoveHandler(func(args *widget.WindowChangedEventArgs) {
			fmt.Println("Window Moved")
		}),
		//Set the callback that triggers when a resize is complete
		widget.WindowOpts.ResizeHandler(func(args *widget.WindowChangedEventArgs) {
			fmt.Println("Window Resized")
		}),
	)

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
			println("setting btn clicked")
			// x, y := window.Contents.PreferredSize()
			s := ebiten.DeviceScaleFactor()
			sInt := int(math.Ceil(s - 0.5))
			// println(x, y)
			x := ((1280 * sInt) - 800) / 2
			y := ((720 * sInt) - 800) / 2
			r := image.Rect(0, 0, 800, 800)
			r = r.Add(image.Point{x, y})
			//Set the windows location to the rect.
			window.SetLocation(r)
			//Add the window to the UI.
			//Note: If the window is already added, this will just move the window and not add a duplicate.
			ui.AddWindow(window)
		}),
	)
	buttonContainer.AddChild(settingBtn)
	rootContainer.AddChild(buttonContainer)

	// construct the UI
	ui = &ebitenui.UI{
		Container: rootContainer,
	}

	return ui
}
