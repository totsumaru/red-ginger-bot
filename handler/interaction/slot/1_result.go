package slot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/kifuneso-bot/handler/interaction/utils"
	"github.com/techstart35/kifuneso-bot/internal/color"
	"github.com/techstart35/kifuneso-bot/internal/errors"
)

// 1回目の数字を送信します
func SendFirstNumber(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	editFunc, err := utils.SendInteractionWaitingMessage(s, i, true, true)
	if err != nil {
		return errors.NewError("Waitingメッセージが送信できません")
	}

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

	webhook := &discordgo.WebhookEdit{
		Embeds:     &[]*discordgo.MessageEmbed{embed},
		Components: &[]discordgo.MessageComponent{actions},
	}
	if _, err = editFunc(i.Interaction, webhook); err != nil {
		return errors.NewError("レスポンスを送信できません", err)
	}

	return nil
}
