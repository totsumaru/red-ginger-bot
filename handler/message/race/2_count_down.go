package race

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/kifuneso-bot/internal/errors"
)

// カウントダウンメッセージを送信します
func sendCountDown(s *discordgo.Session, channelID string, leftMin int) error {
	var description = `
あと%d分で開始します
`

	embed := &discordgo.MessageEmbed{
		Description: fmt.Sprintf(description, leftMin),
	}

	_, err := s.ChannelMessageSendEmbed(channelID, embed)
	if err != nil {
		return errors.NewError("メッセージを送信できません", err)
	}

	return nil
}
