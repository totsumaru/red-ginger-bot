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
	Title         = "RG SLOT"
	Prize_Lose    = 0
	Prize_Big     = 1
	Prize_Small   = 3
	Prize_OneMore = 4

	Value_Beni   = "紅"
	Value_Shou   = "生"
	Value_Ga     = "姜"
	Value_RED    = "RED"
	Value_GIN    = "GIN"
	Value_GER    = "GER"
	Value_Cherry = "🍒"
)

const (
	Prize_BeniShouGa = Value_Beni + Value_Shou + Value_Ga   // 紅-生-姜
	Prize_RedGinGer  = Value_RED + Value_GIN + Value_GER    // RED-GIN-GER
	Prize_Red_3      = Value_RED + Value_RED + Value_RED    // REDx3
	Prize_Gin_3      = Value_GIN + Value_GIN + Value_GIN    // GINx3
	Prize_Ger_3      = Value_GER + Value_GER + Value_GER    // GERx3
	Prize_Beni_3     = Value_Beni + Value_Beni + Value_Beni // 紅x3
	Prize_Shou_3     = Value_Shou + Value_Shou + Value_Shou // 生x3
	Prize_Ga_3       = Value_Ga + Value_Ga + Value_Ga       // 姜x3
)

// ランダムな値を1つ取得します
func getRandomValue(num int, firstValue, secondValue string) string {
	rand.Seed(time.Now().UnixNano())

	switch num {
	case 1:
		elements := []string{
			Value_Beni,
			Value_RED,
		}

		index := rand.Intn(len(elements))
		return elements[index]

	case 2:
		elements := []string{
			Value_Beni,
			Value_Shou,
			Value_Ga,
			Value_RED,
			Value_GIN,
			Value_GER,
		}

		if oneThirdChance() {
			// 1/3の確率でリーチを送信します
			elements = []string{
				firstValue,
				Value_Shou,
				Value_GIN,
			}
		}

		index := rand.Intn(len(elements))
		return elements[index]

	case 3:
		elements := []string{
			Value_Beni,
			Value_Shou,
			Value_Ga,
			Value_RED,
			Value_GIN,
			Value_GER,
		}

		// リーチじゃなければチェリーを入れる
		if !isReach(firstValue, secondValue) {
			elements = append(elements, Value_Cherry)
		}

		index := rand.Intn(len(elements))
		return elements[index]

	default:
		// ここはあり得ないはず
		elements := []string{
			Value_Beni,
			Value_Shou,
			Value_Ga,
			Value_RED,
			Value_GIN,
			Value_GER,
		}

		index := rand.Intn(len(elements))
		return elements[index]
	}
}

// 当たりを判定します
func judgePrize(first, second, third string) int {
	value := first + second + third

	prize := map[string]int{
		Prize_BeniShouGa: Prize_Big,   // 紅生姜
		Prize_RedGinGer:  Prize_Big,   // RED_GIN_GER
		Prize_Beni_3:     Prize_Small, // 紅x3
		Prize_Shou_3:     Prize_Small, // 生x3
		Prize_Ga_3:       Prize_Small, // 姜x3
		Prize_Red_3:      Prize_Small, // REDx3
		Prize_Gin_3:      Prize_Small, // GINx3
		Prize_Ger_3:      Prize_Small, // GERx3
	}

	if strings.Contains(value, Value_Cherry) {
		return Prize_OneMore
	}

	if prizeNum, ok := prize[value]; ok {
		return prizeNum
	}

	return Prize_Lose
}

// リーチかどうか判定します
func isReach(firstValue, secondValue string) bool {
	value := firstValue + secondValue

	prize := map[string]bool{
		Value_Beni + Value_Shou: true, // 紅生姜
		Value_RED + Value_GIN:   true, // RED_GIN_GER
		Value_Beni + Value_Beni: true, // 紅x3
		Value_Shou + Value_Shou: true, // 生x3
		Value_Ga + Value_Ga:     true, // 姜x3
		Value_RED + Value_RED:   true, // REDx3
		Value_GIN + Value_GIN:   true, // GINx3
		Value_GER + Value_GER:   true, // GERx3
	}

	if _, ok := prize[value]; ok {
		return true
	}

	return false
}

// 1/3の確率でtrueを返します
func oneThirdChance() bool {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(3) == 0
}
