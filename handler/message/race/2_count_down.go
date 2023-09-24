package race

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/kifuneso-bot/internal/color"
	"github.com/techstart35/kifuneso-bot/internal/errors"
)

// カウントダウンメッセージを送信します
func sendCountDown(
	s *discordgo.Session,
	guildID, channelID, msgID string,
	leftMin int,
) error {
	var description = `
あと%d分で開始します

↓エントリーはこちら
%s
`

	link := fmt.Sprintf(
		"https://discord.com/channels/%s/%s/%s",
		guildID, channelID, msgID,
	)

	embed := &discordgo.MessageEmbed{
		Description: fmt.Sprintf(description, leftMin, link),
		Color:       color.Yellow,
	}

	_, err := s.ChannelMessageSendEmbed(channelID, embed)
	if err != nil {
		return errors.NewError("メッセージを送信できません", err)
	}

	return nil
}
