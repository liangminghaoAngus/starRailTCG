package main

import (
	"log"
	"math/rand"
	"sort"
	"starRailTCG/enums"
	"time"
)

type Units []UnitItem

func (u Units) Len() int {
	return len(u)
}

func (u Units) Less(i, j int) bool {
	f, s := u[i].Type, u[j].Type
	return f < s
}

func (u Units) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

type UnitItem struct {
	Type     enums.AttributeInt
	TypeName enums.Attribute
}

type Player struct {
	Units           Units         // 拥有的战斗单元
	OnBoardCharater interface{}   // todo 场上的角色
	Charaters       []interface{} // todo 阵容中的角色
	Active          bool          // 我的回合
	EndRound        int           // 结束回合
}

// 开始回合
func (p *Player) StartRound() {
	p.randomUnits(enums.RoundUnit)
}

func (p *Player) randomUnits(rounds int) {
	if len(p.Units) == 0 {
		p.Units = make(Units, 0)
	}

	for i := 0; i < rounds; i++ {
		rand.Seed(time.Now().UnixNano())
		ind := rand.Intn(len(enums.AllAttributes))
		enumType := enums.AllAttributes[ind]
		enumTypeName := enums.GetAttributeStr(enumType)
		p.Units = append(p.Units, UnitItem{Type: enumType, TypeName: enumTypeName})
	}
	sort.Sort(p.Units)
	log.Printf("unit list :%+v", p.Units)
}
