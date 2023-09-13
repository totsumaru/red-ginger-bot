package race

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/kifuneso-bot/internal/color"
	"github.com/techstart35/kifuneso-bot/internal/errors"
)

// 途中経過のメッセージを送信します
//
// User(ロボ)を順番通り返します。
func sendProgress(s *discordgo.Session, channelID string) ([]EntryUser, error) {
	currentRank := EntryUsers

	for i := 0; i < 10; i++ {
		if i != 0 {
			time.Sleep(12 * time.Second)
		}

		entryUsers, err := sendCommentary(s, currentRank, channelID, i)
		if err != nil {
			return currentRank, errors.NewError("実況を送信できません", err)
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
		return currentRank, errors.NewError("結果を送信できません", err)
	}

	return currentRank, nil
}

// 実況を送信します
func sendCommentary(
	s *discordgo.Session, currentRank []EntryUser, channelID string, index int,
) ([]EntryUser, error) {
	res := make([]EntryUser, 0)
	lines := make([]string, 0)

	for ranking, entryUser := range currentRank {
		// pointとその文言を決定する
		point, text := getRandResult(ranking+1, index)
		// point
		entryUser.AddPoint(point)
		res = append(res, entryUser)
		// text
		lines = append(lines, fmt.Sprintf("%d. %s__**%s**__ %s",
			ranking+1,
			entryUser.Emoji,
			entryUser.Name,
			text,
		))
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
		Image: &discordgo.MessageEmbedImage{
			URL: entryUsers[0].ImageURL,
		},
	}

	_, err := s.ChannelMessageSendEmbed(channelID, embed)
	if err != nil {
		return errors.NewError("メッセージを送信できません", err)
	}

	return nil
}

// ランダムな実況を取得します
// point: textを返します
func getRandResult(ranking, index int) (int, string) {
	rand.Seed(time.Now().UnixNano())

	// 9回目、10回目は1/5の確率で+6となる
	switch index + 1 {
	case 9, 10:
		if rand.Intn(5) == 0 {
			result := point6[rand.Intn(len(point6))]
			return result.point, result.text
		}
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
