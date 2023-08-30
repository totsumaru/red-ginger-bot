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
func sendProgress(s *discordgo.Session, channelID string) error {
	currentRank := EntryUsers

	for i := 0; i < 6; i++ {
		if i != 0 {
			time.Sleep(8 * time.Second)
		}

		entryUsers, err := sendCommentary(s, currentRank, channelID, i)
		if err != nil {
			return errors.NewError("実況を送信できません", err)
		}
		currentRank = entryUsers

		time.Sleep(2 * time.Second)

		if err := sendRank(s, currentRank, channelID); err != nil {
			return errors.NewError("現状のランキングを送信できません", err)
		}
	}

	return nil
}

// 実況を送信します
func sendCommentary(
	s *discordgo.Session,
	currentRank []EntryUser,
	channelID string,
	index int,
) ([]EntryUser, error) {
	res := make([]EntryUser, 0)
	lines := make([]string, 0)

	for _, entryUser := range currentRank {
		// pointとその文言を決定する
		point, text := getRandResult()
		// point
		entryUser.AddPoint(point)
		res = append(res, entryUser)
		// text
		lines = append(lines, fmt.Sprintf("%s%sは%s", entryUser.Emoji, entryUser.Name, text))
	}

	description := `
途中経過 %d/6

%s
`

	embed := &discordgo.MessageEmbed{
		Description: fmt.Sprintf(description, index, strings.Join(lines, "\n")),
	}

	_, err := s.ChannelMessageSendEmbed(channelID, embed)
	if err != nil {
		return res, errors.NewError("メッセージを送信できません", err)
	}

	return res, nil
}

// 現状のランキングを送信します
func sendRank(
	s *discordgo.Session,
	currentRank []EntryUser,
	channelID string,
) error {
	// ポイント順に並べます
	sort.Slice(currentRank, func(i, j int) bool {
		return currentRank[i].Point > currentRank[j].Point
	})

	lines := make([]string, 0)
	for i, entryUser := range currentRank {
		lines = append(lines, fmt.Sprintf(
			"%d位-%s%s(%dpt)",
			i+1, entryUser.Emoji, entryUser.Name, entryUser.Point,
		))
	}

	description := `
現在の順位

%s
`

	embed := &discordgo.MessageEmbed{
		Description: fmt.Sprintf(description, strings.Join(lines, "\n")),
		Color:       color.Blue,
	}

	_, err := s.ChannelMessageSendEmbed(channelID, embed)
	if err != nil {
		return errors.NewError("メッセージを送信できません", err)
	}

	return nil
}

// ランダムな実況を取得します
func getRandResult() (int, string) {
	rand.Seed(time.Now().UnixNano())

	type Res struct {
		point int
		text  string
	}

	res := []Res{
		{point: 1, text: "1pt獲得した"},
		{point: 2, text: "2pt獲得した"},
		{point: 3, text: "3pt獲得した"},
	}

	// ランダムなインデックスを生成
	randomIndex := rand.Intn(len(res))

	// ランダムな要素を取得
	return res[randomIndex].point, res[randomIndex].text
}
