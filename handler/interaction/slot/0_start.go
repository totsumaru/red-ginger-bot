package slot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/kifuneso-bot/internal/color"
	"github.com/techstart35/kifuneso-bot/internal/errors"
)

// slotの開始メッセージを送信します
func SendStartMessage(
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
	isUpdateMessage bool,
) error {
	actions := discordgo.ActionsRow{
		Components: []discordgo.MessageComponent{
			ButtonComponent(1, false),
			ButtonComponent(2, true),
			ButtonComponent(3, true),
		},
	}

	description := `
スロットを開始します。
`

	embed := &discordgo.MessageEmbed{
		Title:       Title,
		Description: description,
		Color:       color.Red,
	}

	responseType := discordgo.InteractionResponseChannelMessageWithSource
	if isUpdateMessage {
		responseType = discordgo.InteractionResponseUpdateMessage
	}

	resp := &discordgo.InteractionResponse{
		Type: responseType,
		Data: &discordgo.InteractionResponseData{
			Embeds:     []*discordgo.MessageEmbed{embed},
			Components: []discordgo.MessageComponent{actions},
			Flags:      discordgo.MessageFlagsEphemeral,
		},
	}

	if err := s.InteractionRespond(i.Interaction, resp); err != nil {
		return errors.NewError("レスポンスを送信できません", err)
	}

	return nil
}
