package enums

type ActiveScreen int

const (
	ScreenOnBoard  ActiveScreen = iota + 1 // 初始界面
	ScreenGameMode                         // 游戏场景
)

type AttributeInt int

// 万能 物理 火 风 冰 雷 虚数 量子
const (
	AttributeAllInt AttributeInt = iota + 1
	AttributePhysicsInt
	AttributeFireInt
	AttributeWindInt
	AttributeIceInt
	AttributeLightingInt
	AttributeQuantumInt
	AttributeImaginaryInt
)

// 命途属性
type Attribute string

// 物理 火 风 冰 雷 虚数 量子
const (
	AttributeAll       Attribute = "万能"
	AttributePhysics   Attribute = "物理"
	AttributeFire      Attribute = "火"
	AttributeWind      Attribute = "风"
	AttributeIce       Attribute = "冰"
	AttributeLighting  Attribute = "雷"
	AttributeQuantum   Attribute = "量子"
	AttributeImaginary Attribute = "虚数"
)

var attributeIntStr = map[AttributeInt]Attribute{
	AttributeAllInt:       AttributeAll,
	AttributePhysicsInt:   AttributePhysics,
	AttributeFireInt:      AttributeFire,
	AttributeWindInt:      AttributeWind,
	AttributeLightingInt:  AttributeLighting,
	AttributeIceInt:       AttributeIce,
	AttributeQuantumInt:   AttributeQuantum,
	AttributeImaginaryInt: AttributeImaginary,
}

var attributeStrInt = map[Attribute]AttributeInt{
	AttributeAll:       AttributeAllInt,
	AttributeFire:      AttributeFireInt,
	AttributeIce:       AttributeIceInt,
	AttributeImaginary: AttributeImaginaryInt,
	AttributeLighting:  AttributeLightingInt,
	AttributePhysics:   AttributePhysicsInt,
	AttributeQuantum:   AttributeQuantumInt,
	AttributeWind:      AttributeWindInt,
}

var AllAttributes = []AttributeInt{AttributeAllInt, AttributePhysicsInt, AttributeFireInt, AttributeWindInt, AttributeLightingInt, AttributeIceInt, AttributeQuantumInt, AttributeImaginaryInt}

var RoundUnit int = 8

var MaxUnits int = 16

func GetAttributeStr(attrInt AttributeInt) Attribute {
	if str, ok := attributeIntStr[attrInt]; ok {
		return str
	} else {
		return ""
	}
}

func GetAttributeInt(attrStr Attribute) AttributeInt {
	if num, ok := attributeStrInt[attrStr]; ok {
		return num
	} else {
		return -1
	}
}
