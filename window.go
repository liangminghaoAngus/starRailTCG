package main

import (
	"fmt"
	"image/color"
	_ "image/gif"
	"log"
	"starRailTCG/config"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/yohamta/furex/v2"
)

type Game struct {
	i             uint8
	cfg           *config.Config
	OnboardScreen *furex.View
}

func NewGame(cfg *config.Config) *Game {

	ebiten.SetWindowSize(cfg.ScreenWidth, cfg.ScreenHeight) //窗口大小
	ebiten.SetWindowTitle(cfg.Title)                        //窗口标题

	onboard := OnBoardView(cfg.ScreenWidth, cfg.ScreenHeight)
	return &Game{
		cfg:           cfg,
		OnboardScreen: onboard,
	}
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.i++
	if g.i < 255 {
		screen.Fill(Hex2RGB("#0dceda", g.i))
	} else {
		g.i = 0
	}
	// 加载背景图片
	bgGIF, bgGIFFile, err := ebitenutil.NewImageFromFile("./static/imgs/onboard01.gif")
	if err != nil {
		log.Fatalf("ebitenutil.NewImageFromFile failed :%v", err)
	}
	log.Printf("bgGIFFile :%+v", bgGIFFile)
	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Scale(0.2, 0.2)
	opt.GeoM.Translate(float64(g.cfg.ScreenWidth-bgGIFFile.Bounds().Dx())/2, float64(g.cfg.ScreenHeight-bgGIFFile.Bounds().Dy())/2)
	screen.DrawImage(bgGIF, opt)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f\nFPS: %0.2f", ebiten.ActualTPS(), ebiten.ActualFPS()))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.cfg.ScreenWidth / 2, g.cfg.ScreenHeight / 2 //窗口分辨率
}

func Hex2RGB(color16 string, alpha uint8) color.RGBA {
	r, _ := strconv.ParseInt(color16[:2], 16, 10)
	g, _ := strconv.ParseInt(color16[2:4], 16, 18)
	b, _ := strconv.ParseInt(color16[4:], 16, 10)
	return color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: alpha}
}

func OnBoardView(width, height int) *furex.View {

	onboardView := &furex.View{
		Width:        width,
		Height:       height,
		Direction:    furex.Row,
		Justify:      furex.JustifyCenter,
		AlignItems:   furex.AlignItemCenter,
		AlignContent: furex.AlignContentCenter,
		Wrap:         furex.Wrap,
	}

	target := make([]*furex.View, 0)

	// bgImgView := &furex.View{
	// 	Width:  width,
	// 	Height: height,
	// }

	// target = append(target, bgImgView)

	// 制作按钮 todo

	for _, item := range target {
		onboardView.AddChild(item)
	}
	return onboardView

}
