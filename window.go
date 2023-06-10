package main

import (
	"fmt"
	"image/color"
	_ "image/gif"
	"starRailTCG/common"
	"starRailTCG/config"
	"starRailTCG/enums"
	"starRailTCG/screen"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/yohamta/furex/v2"
)

type Game struct {
	i             uint8
	cfg           *config.Config
	activeScreen  enums.ActiveScreen
	OnboardScreen *furex.View
}

func NewGame(cfg *config.Config) *Game {

	ebiten.SetWindowSize(cfg.ScreenWidth, cfg.ScreenHeight) //窗口大小
	ebiten.SetWindowTitle(cfg.Title)                        //窗口标题

	onboard := screen.NewOnBoardScreen()
	return &Game{
		cfg:           cfg,
		OnboardScreen: onboard,
		activeScreen:  common.ChangeScreen(enums.ScreenOnBoard),
	}
}

func (g *Game) Update() error {
	select {
	case s := <-common.ScreenChan:
		g.activeScreen = s
	default:

	}

	switch g.activeScreen {
	case enums.ScreenOnBoard:
		g.OnboardScreen.Update()
	}
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
	// sw, sh := screen.Bounds().Dx(), screen.Bounds().Dy()
	// // 加载背景图片
	// bgGIF, bgGIFFile, err := ebitenutil.NewImageFromFile("./static/imgs/onboard.gif")
	// if err != nil {
	// 	log.Fatalf("ebitenutil.NewImageFromFile failed :%v", err)
	// }
	// opt := &ebiten.DrawImageOptions{}
	// opt.GeoM.Scale(0.6, 0.6)
	// scale := ebiten.DeviceScaleFactor()
	// opt.GeoM.Scale(scale, scale)
	// opt.GeoM.Translate(float64(sw-bgGIFFile.Bounds().Dx()-40)/2, float64(sh-bgGIFFile.Bounds().Dy()-40)/2)
	// opt.Filter = ebiten.FilterLinear
	// screen.DrawImage(bgGIF, opt)
	switch g.activeScreen {
	case enums.ScreenOnBoard:
		g.OnboardScreen.Draw(screen)
	case enums.ScreenGameMode:
		// println("switch game mode screen")
	}

	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f\nFPS: %0.2f", ebiten.ActualTPS(), ebiten.ActualFPS()))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	// s := ebiten.DeviceScaleFactor()
	// return int(float64(outsideWidth) * s), int(float64(outsideHeight) * s)
	return outsideWidth, outsideHeight
}

func Hex2RGB(color16 string, alpha uint8) color.RGBA {
	r, _ := strconv.ParseInt(color16[:2], 16, 10)
	g, _ := strconv.ParseInt(color16[2:4], 16, 18)
	b, _ := strconv.ParseInt(color16[4:], 16, 10)
	return color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: alpha}
}
