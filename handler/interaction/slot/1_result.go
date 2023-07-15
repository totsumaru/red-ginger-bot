package slot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/kifuneso-bot/internal/color"
	"github.com/techstart35/kifuneso-bot/internal/errors"
)

// 1回目の数字を送信します
func SendFirstNumber(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	value := getEachValue(1, "", "")

	actions := discordgo.ActionsRow{
		Components: []discordgo.MessageComponent{
			ButtonComponent(1, true),
			ButtonComponent(2, false),
			ButtonComponent(3, true),
		},
	}

	embed := &discordgo.MessageEmbed{
		Title:       Title,
		Description: fmt.Sprintf(DescriptionTmpl, value, "-", "-"),
		Color:       color.Red,
	}

	resp := &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseUpdateMessage,
		Data: &discordgo.InteractionResponseData{
			Embeds:     []*discordgo.MessageEmbed{embed},
			Flags:      discordgo.MessageFlagsEphemeral,
			Components: []discordgo.MessageComponent{actions},
		},
	}

	if err := s.InteractionRespond(i.Interaction, resp); err != nil {
		return errors.NewError("レスポンスを送信できません", err)
	}

	return nil
}
