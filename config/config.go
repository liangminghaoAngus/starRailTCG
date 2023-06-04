package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Config struct {
	init         bool   `json:"-"`
	ScreenWidth  int    `json:"screen-width"`
	ScreenHeight int    `json:"screen-height"`
	Title        string `json:"title"`
}

const (
	ModeDev  = "dev"  // 开发模式
	ModeProd = "prod" // 运行模式
)

func switchMode(mode string) (fileName string) {
	switch mode {
	case ModeDev:
		fileName = "dev"
	case ModeProd:
		fileName = "prod"
	}
	return fmt.Sprintf("%s.json", fileName)
}

func Init(mode string) *Config {
	cfg := &Config{}
	fileName := switchMode(mode)

	// todo 可以换成 path 解析
	f, err := os.Open(fmt.Sprintf("./config/%s", fileName))
	if err != nil {
		log.Fatalf("os.Open failed:%v \n", err)
	}

	if err := json.NewDecoder(f).Decode(&cfg); err != nil {
		log.Fatalf("json.NewDecode failed:%v \n", err)
	}

	cfg.init = true
	return cfg
}
