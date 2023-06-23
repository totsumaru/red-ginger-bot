package slot

import (
	"math/rand"
	"time"
)

const DescriptionTmpl = `
**%s｜%s｜%s**
`

const Title = "RG SLOT"

// ランダムな数字を1つ取得します
func getRandomNum() uint {
	elements := []int{1, 2, 3, 5, 6, 7, 8, 9}

	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(elements))

	return uint(elements[index])
}
