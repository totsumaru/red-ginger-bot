package slot

import (
	"math/rand"
	"strings"
	"time"
)

const DescriptionTmpl = `
**%s｜%s｜%s**
`

const (
	Title = "RG SLOT"

	Prize_Lose    = "prize-lose"
	Prize_Big     = "prize-big"
	Prize_Small   = "prize-small"
	Prize_OneMore = "prize-one-more"
	Prize_Secret  = "prize-secret"

	Value_Beni   = "紅"
	Value_Shou   = "生"
	Value_Ga     = "姜"
	Value_RED    = "RED"
	Value_GIN    = "GIN"
	Value_GER    = "GER"
	Value_Cherry = "🍒"
	Value_Secret = "？"
)

const (
	String_BeniShouGa = Value_Beni + Value_Shou + Value_Ga         // 紅-生-姜
	String_RedGinGer  = Value_RED + Value_GIN + Value_GER          // RED-GIN-GER
	String_Red_3      = Value_RED + Value_RED + Value_RED          // REDx3
	String_Gin_3      = Value_GIN + Value_GIN + Value_GIN          // GINx3
	String_Ger_3      = Value_GER + Value_GER + Value_GER          // GERx3
	String_Beni_3     = Value_Beni + Value_Beni + Value_Beni       // 紅x3
	String_Shou_3     = Value_Shou + Value_Shou + Value_Shou       // 生x3
	String_Ga_3       = Value_Ga + Value_Ga + Value_Ga             // 姜x3
	String_Secret_3   = Value_Secret + Value_Secret + Value_Secret // ？x3
)

// スロットの各回の値を1つ取得します
//
// 大当たり	：1/33（2回目1/3, 3回目1/11）
// 小当たり	：1/32（2回目1/3, 3回目1/11）
// ？？？	：1/256
//
// 1回目
// - 基本的には全て同じ確率で出現
// - ？は 1/256 の確率で出現
// - チェリーは出ない
//
// 2回目
// - 基本的には1/3の確率でリーチになる
// - 1回目が？であれば2回目も？となる
//
// 3回目
// - 大当たりリーチの場合は、1/21で大当たりにする（2回目は1/3でリーチなので、掛け算して1/63で大当たりにする）
// - 小当たりリーチの場合は、1/11で小当たりにする（2回目は1/3でリーチなので、掛け算して1/33で小当たりにする）
// - ??と来たら、3回目も?とする
func getEachValue(num int, firstValue, secondValue string) string {
	rand.Seed(time.Now().UnixNano())

	switch num {
	case 1:
		return getFirstValue()
	case 2:
		return getSecondValue(firstValue)
	case 3:
		return getThirdValue(firstValue, secondValue)
	default:
		// ここは有り得ない
		return ""
	}
}

// 1つめの数字を取得します
func getFirstValue() string {
	// ？は1/256で出現
	if oneRequestValueChance(256) {
		return Value_Secret
	}

	return getRandomValue([]string{
		Value_Beni,
		Value_Shou,
		Value_Ga,
		Value_RED,
		Value_GIN,
		Value_GER,
	})
}

// 2つめの数字を取得します
func getSecondValue(firstValue string) string {
	// 1回目が？なら2回目も？を返す
	if firstValue == Value_Secret {
		return Value_Secret
	}

	// 1/3の確率でリーチを送信します
	if oneRequestValueChance(3) {
		switch firstValue {
		case Value_RED:
			return getRandomValue([]string{
				Value_RED,
				Value_GIN,
			})
		case Value_Beni:
			return getRandomValue([]string{
				Value_Beni,
				Value_Shou,
			})
		default:
			return firstValue
		}
	} else {
		switch firstValue {
		case Value_RED:
			return getRandomValue([]string{
				Value_GER,
				Value_Beni,
				Value_Shou,
				Value_Ga,
			})
		case Value_Beni:
			return getRandomValue([]string{
				Value_Ga,
				Value_RED,
				Value_GIN,
				Value_GER,
			})
		default:
			// リーチにしないために、firstValueと違う値を返す
			for {
				res := getRandomValue([]string{
					Value_Beni,
					Value_Shou,
					Value_Ga,
					Value_RED,
					Value_GIN,
					Value_GER,
				})

				if res != firstValue {
					return res
				}
			}
		}
	}
}

// 3つめの数字を取得します
func getThirdValue(firstValue, secondValue string) string {
	value := firstValue + secondValue

	switch value {
	case Value_Beni + Value_Shou: // 紅_生
		if oneRequestValueChance(11) {
			return Value_Ga
		} else {
			return getRandomValue([]string{
				Value_Beni,
				Value_Shou,
				Value_RED,
				Value_GIN,
				Value_GER,
				Value_Cherry,
			})
		}
	case Value_RED + Value_GIN: // RED_GIN
		if oneRequestValueChance(11) {
			return Value_GER
		} else {
			return getRandomValue([]string{
				Value_Beni,
				Value_Shou,
				Value_Ga,
				Value_RED,
				Value_GIN,
				Value_Cherry,
			})
		}
	case Value_Beni + Value_Beni: // 紅x2
		if oneRequestValueChance(11) {
			return Value_Beni
		} else {
			return getRandomValue([]string{
				Value_Shou,
				Value_Ga,
				Value_RED,
				Value_GIN,
				Value_GER,
				Value_Cherry,
			})
		}
	case Value_Shou + Value_Shou: // 生x2
		if oneRequestValueChance(11) {
			return Value_Shou
		} else {
			return getRandomValue([]string{
				Value_Beni,
				Value_Ga,
				Value_RED,
				Value_GIN,
				Value_GER,
				Value_Cherry,
			})
		}
	case Value_Ga + Value_Ga: // 姜x2
		if oneRequestValueChance(11) {
			return Value_Ga
		} else {
			return getRandomValue([]string{
				Value_Beni,
				Value_Shou,
				Value_RED,
				Value_GIN,
				Value_GER,
				Value_Cherry,
			})
		}
	case Value_RED + Value_RED: // REDx2
		if oneRequestValueChance(11) {
			return Value_RED
		} else {
			return getRandomValue([]string{
				Value_Beni,
				Value_Shou,
				Value_Ga,
				Value_GIN,
				Value_GER,
				Value_Cherry,
			})
		}
	case Value_GIN + Value_GIN: // GINx2
		if oneRequestValueChance(11) {
			return Value_GIN
		} else {
			return getRandomValue([]string{
				Value_Beni,
				Value_Shou,
				Value_Ga,
				Value_RED,
				Value_GER,
				Value_Cherry,
			})
		}
	case Value_GER + Value_GER: // GERx2
		if oneRequestValueChance(11) {
			return Value_GER
		} else {
			return getRandomValue([]string{
				Value_Beni,
				Value_Shou,
				Value_Ga,
				Value_RED,
				Value_GIN,
				Value_Cherry,
			})
		}
	case Value_Secret + Value_Secret: // ？？
		return Value_Secret
	default:
		// NotReach
		return getRandomValue([]string{
			Value_Beni,
			Value_Shou,
			Value_Ga,
			Value_RED,
			Value_GIN,
			Value_GER,
			Value_Cherry,
		})
	}
}

// 当たりを判定します
func judgePrize(first, second, third string) string {
	value := first + second + third

	prizeKind := map[string]string{
		String_BeniShouGa: Prize_Big,    // 紅生姜
		String_RedGinGer:  Prize_Big,    // RED_GIN_GER
		String_Beni_3:     Prize_Small,  // 紅x3
		String_Shou_3:     Prize_Small,  // 生x3
		String_Ga_3:       Prize_Small,  // 姜x3
		String_Red_3:      Prize_Small,  // REDx3
		String_Gin_3:      Prize_Small,  // GINx3
		String_Ger_3:      Prize_Small,  // GERx3
		String_Secret_3:   Prize_Secret, // ？x3
	}

	if strings.Contains(value, Value_Cherry) {
		return Prize_OneMore
	}

	if prize, ok := prizeKind[value]; ok {
		return prize
	}

	return Prize_Lose
}

// 1/引数の数 の確率でtrueを返します
func oneRequestValueChance(num int) bool {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(num) == 0
}

// 配列からランダムな値を取得します
func getRandomValue(elements []string) string {
	index := rand.Intn(len(elements))
	return elements[index]
}
