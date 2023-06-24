package slot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/kifuneso-bot/internal/color"
	"github.com/techstart35/kifuneso-bot/internal/errors"
	"strings"
)

// 2回目の数字を送信します
func SendSecondNumber(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	lastDescription := strings.Replace(i.Message.Embeds[0].Description, "**", "", -1)
	lastDescription = strings.Replace(lastDescription, " ", "", -1)
	lastValue := strings.Split(lastDescription, "｜")[0]
	value := getRandomValue(2, lastValue, "")

	actions := discordgo.ActionsRow{
		Components: []discordgo.MessageComponent{
			ButtonComponent(1, true),
			ButtonComponent(2, true),
			ButtonComponent(3, false),
		},
	}

	embed := &discordgo.MessageEmbed{
		Title:       Title,
		Description: fmt.Sprintf(DescriptionTmpl, lastValue, value, "-"),
		Color:       color.Red,
	}

	resp := &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseUpdateMessage,
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
