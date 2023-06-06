package main

import (
	"flag"
	"log"
	"starRailTCG/config"

	"github.com/hajimehoshi/ebiten/v2"
)

var globalConfig *config.Config

func main() {
	var mode string
	flag.StringVar(&mode, "mode", "", "启动模式配置文件")
	flag.Parsed()
	if mode == "" {
		mode = config.ModeDev
	}
	globalConfig = config.Init(mode)
	MainGame := NewGame(globalConfig)
	if err := ebiten.RunGame(MainGame); err != nil {
		log.Fatal(err)
	}
}
