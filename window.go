package main

import (
	"fmt"
	"image/color"
	"starRailTCG/config"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	i   uint8
	cfg *config.Config
}

func NewGame(cfg *config.Config) *Game {

	ebiten.SetWindowSize(cfg.ScreenWidth, cfg.ScreenHeight) //窗口大小
	ebiten.SetWindowTitle(cfg.Title)                        //窗口标题

	return &Game{
		cfg: cfg,
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
