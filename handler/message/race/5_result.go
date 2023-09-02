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
	selectedUsers []*discordgo.User,
) error {
	lines := make([]string, 0)

	lines = append(lines, "---当選者---\n")

	if len(selectedUsers) == 0 {
		lines = append(lines, "当選者なし")
	} else {
		for _, user := range selectedUsers {
			lines = append(lines, fmt.Sprintf("<@%s>", user.ID))
		}
	}

	if _, err := s.ChannelMessageSend(channelID, strings.Join(lines, "\n")); err != nil {
		return errors.NewError("メッセージを送信できません", err)
	}

	return nil
}
