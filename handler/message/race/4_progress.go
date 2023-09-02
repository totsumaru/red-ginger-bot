package race

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/kifuneso-bot/internal/color"
	"github.com/techstart35/kifuneso-bot/internal/errors"
	"math/rand"
	"sort"
	"strings"
	"time"
)

// 途中経過のメッセージを送信します
//
// 1番のUser(ロボ)を返します。
func sendProgress(s *discordgo.Session, channelID string) (EntryUser, error) {
	currentRank := EntryUsers

	for i := 0; i < 10; i++ {
		if i != 0 {
			time.Sleep(8 * time.Second)
		}

		entryUsers, err := sendCommentary(s, currentRank, channelID, i)
		if err != nil {
			return currentRank[0], errors.NewError("実況を送信できません", err)
		}
		currentRank = entryUsers

		time.Sleep(5 * time.Second)

		// ポイント順に並べます
		sort.Slice(currentRank, func(i, j int) bool {
			return currentRank[i].Point > currentRank[j].Point
		})
	}

	// 最後に結果を送信します
	if err := sendResult(s, channelID, currentRank); err != nil {
		return currentRank[0], errors.NewError("結果を送信できません", err)
	}

	return currentRank[0], nil
}

// 実況を送信します
func sendCommentary(
	s *discordgo.Session, currentRank []EntryUser, channelID string, index int,
) ([]EntryUser, error) {
	res := make([]EntryUser, 0)
	lines := make([]string, 0)

	for ranking, entryUser := range currentRank {
		// pointとその文言を決定する
		point, text := getRandResult(ranking + 1)
		// point
		entryUser.AddPoint(point)
		res = append(res, entryUser)
		// text
		lines = append(lines, fmt.Sprintf("%s%sは%s", entryUser.Emoji, entryUser.Name, text))
	}

	description := `
途中経過 %d/10

%s
`

	embed := &discordgo.MessageEmbed{
		Description: fmt.Sprintf(description, index+1, strings.Join(lines, "\n")),
	}

	_, err := s.ChannelMessageSendEmbed(channelID, embed)
	if err != nil {
		return res, errors.NewError("メッセージを送信できません", err)
	}

	return res, nil
}

// 結果を送信します
func sendResult(s *discordgo.Session, channelID string, entryUsers []EntryUser) error {
	description := `
👑結果

%s
`

	lines := make([]string, 0)

	for index, entryUser := range entryUsers {
		line := fmt.Sprintf("%d: %s%s", index+1, entryUser.Emoji, entryUser.Name)
		lines = append(lines, line)
	}

	embed := &discordgo.MessageEmbed{
		Description: fmt.Sprintf(description, strings.Join(lines, "\n")),
		Color:       color.Yellow,
	}

	_, err := s.ChannelMessageSendEmbed(channelID, embed)
	if err != nil {
		return errors.NewError("メッセージを送信できません", err)
	}

	return nil
}

// ランダムな実況を取得します
// point: textを返します
func getRandResult(ranking int) (int, string) {
	rand.Seed(time.Now().UnixNano())

	type Res struct {
		point int
		text  string
	}

	// -2pt
	pointM2 := []Res{
		{point: -2, text: "転んでしまった"},
		{point: -2, text: "つまずいてしまった！"},
		{point: -2, text: "コースを大きく外れている"},
		{point: -2, text: "バランスを崩して転倒！"},
		{point: -2, text: "急に足を引っ張られたようだ"},
		{point: -2, text: "スタミナ切れでペースダウン"},
		{point: -2, text: "突然立ち止まってしまった"},
		{point: -2, text: "障害物に接触、大きく遅れた"},
		{point: -2, text: "足元が狂い、スピードが落ちた"},
		{point: -2, text: "急な坂でスピードが落ちてしまった"},
	}

	// -1pt
	pointM1 := []Res{
		{point: -1, text: "少し減速している"},
		{point: -1, text: "疲れているようだ"},
		{point: -1, text: "ペースが乱れている"},
		{point: -1, text: "スリップしている"},
		{point: -1, text: "集中力が切れているようだ"},
		{point: -1, text: "風に煽られている"},
		{point: -1, text: "足取りが重い"},
		{point: -1, text: "コースを外れそうだ"},
		{point: -1, text: "バランスを取り戻している"},
	}

	// 0pt
	point0 := []Res{
		{point: 0, text: "そのまま走っている"},
		{point: 0, text: "安定した走りを見せている"},
		{point: 0, text: "まだ力を温存している"},
		{point: 0, text: "他のロボと並んでいる"},
		{point: 0, text: "リズムを保っている"},
		{point: 0, text: "一定のペースで進んでいる"},
		{point: 0, text: "落ち着いた走りをしている"},
	}

	// 1pt
	point1 := []Res{
		{point: 1, text: "順調に走っている"},
		{point: 1, text: "少しずつ加速している"},
		{point: 1, text: "力強い走りを見せている"},
		{point: 1, text: "スムーズにカーブを曲がっている"},
		{point: 1, text: "状態が良さそうだ"},
		{point: 1, text: "余裕の表情を見せている"},
		{point: 1, text: "勢いがついてきた"},
	}

	// 2pt
	point2 := []Res{
		{point: 2, text: "スピードをあげてきた"},
		{point: 2, text: "一気に加速！"},
		{point: 2, text: "驚異的なスピードだ"},
		{point: 2, text: "直線で猛然と加速"},
		{point: 2, text: "まるで飛んでいるようだ"},
		{point: 2, text: "疲れを感じさせない走りをしている"},
		{point: 2, text: "最後まで諦めずに走っている"},
	}

	// 3pt
	point3 := []Res{
		{point: 3, text: "力を解放し、猛スピードで走っている"},
		{point: 3, text: "まるでエンジンをつけたような走りをしている"},
		{point: 3, text: "信じられないスピードで突き進む"},
		{point: 3, text: "観客も驚くほどのスピードで走っている"},
		{point: 3, text: "直線で加速！"},
		{point: 3, text: "完璧な走りを見せている"},
		{point: 3, text: "歴史に名を刻む走りをしている"},
	}

	// 1位は1/3の確率で-2となる
	// 最下位は1/3の確率で+3または+2となる
	switch ranking {
	case 1:
		if rand.Intn(3) == 0 {
			result := pointM2[rand.Intn(len(pointM2))]
			return result.point, result.text
		}
	case len(EntryUsers):
		if rand.Intn(3) == 0 {
			// 結果の候補です
			list := point3
			list = append(list, point2...)

			result := list[rand.Intn(len(point3)+len(point2))]
			return result.point, result.text
		}
	}

	all := make([]Res, 0)
	all = append(all, pointM2...)
	all = append(all, pointM1...)
	all = append(all, point0...)
	all = append(all, point1...)
	all = append(all, point2...)
	all = append(all, point3...)

	// ランダムなインデックスを生成
	randomIndex := rand.Intn(len(all))

	// ランダムな要素を取得
	return all[randomIndex].point, all[randomIndex].text
}
