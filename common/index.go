package common

import (
	"starRailTCG/enums"
)

var ScreenChan = make(chan enums.ActiveScreen, 1)

func ChangeScreen(newScreen enums.ActiveScreen) enums.ActiveScreen {
	ScreenChan <- newScreen
	return newScreen
}
