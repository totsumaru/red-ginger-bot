package race

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/kifuneso-bot/internal/errors"
	"strings"
)

// 当たったユーザーを送信します
func sendWinnerUsers(
	s *discordgo.Session,
	channelID string,
	first []*discordgo.User,
	second []*discordgo.User,
) error {
	lines := make([]string, 0)

	lines = append(lines, "---当選者---\n")

	if len(first) == 0 {
		lines = append(lines, "当選者なし")
	} else {
		// 1位を追加
		lines = append(lines, "🥇｜1位---")
		for _, user := range first {
			lines = append(lines, fmt.Sprintf("<@%s>", user.ID))
		}

		// 2位を追加
		lines = append(lines, "\n🥈｜2位---")
		for _, user := range second {
			lines = append(lines, fmt.Sprintf("<@%s>", user.ID))
		}
	}

	if _, err := s.ChannelMessageSend(channelID, strings.Join(lines, "\n")); err != nil {
		return errors.NewError("メッセージを送信できません", err)
	}

	return nil
}
