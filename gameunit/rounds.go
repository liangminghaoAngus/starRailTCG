package gameunit

const (
	FirstRound int = 1
)

type GameRound struct {
	Rounds       int // 回合数，从 1 开始计数
	ActivePlayer int
	isInit       bool // 判断游戏是否初始化完成

	Player0 *Player // todo
	Player1 *Player // todo
}

func NewGameRound(activePlayer int) *GameRound {
	return &GameRound{
		Rounds:       FirstRound,
		ActivePlayer: activePlayer,
	}
}

func (g *GameRound) AddPlayer(playerNumber int, player *Player) bool {
	if g.isInit {
		return false
	}

	// 根据位置添加玩家的位置
	if playerNumber == 0 {
		g.Player0 = player
	} else {
		g.Player1 = player
	}

	_ = g.InitFinish()

	return true
}

// 是否完成游戏初始化配置
func (g *GameRound) InitFinish() bool {
	if g.Player0 != nil && g.Player1 != nil {
		g.isInit = true
	} else {
		g.isInit = false
	}
	return g.isInit
}
