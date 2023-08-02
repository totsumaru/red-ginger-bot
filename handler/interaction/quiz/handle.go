package quiz

import (
	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/kifuneso-bot/internal/color"
	"github.com/techstart35/kifuneso-bot/internal/errors"
)

func HandleQuiz(
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
	interactionID string,
) error {
	info, err := GetEmbedInfo(s, interactionID, i.Member.User.ID)
	if err != nil {
		return errors.NewError("送信する情報を取得できません", err)
	}

	// infoに値がなかった（空）の場合は終了
	if info.Description == "" {
		return nil
	}

	mc := make([]discordgo.MessageComponent, 0)
	for _, b := range info.Button {
		mc = append(mc, b)
	}

	actions := discordgo.ActionsRow{
		Components: mc,
	}

	embed := &discordgo.MessageEmbed{
		Description: info.Description,
		Image: &discordgo.MessageEmbedImage{
			URL: info.ImageURL,
		},
		Color: color.Red,
	}

	tp := discordgo.InteractionResponseUpdateMessage
	if info.IsNewMessage {
		tp = discordgo.InteractionResponseChannelMessageWithSource
	}
	resp := &discordgo.InteractionResponse{
		Type: tp,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{embed},
			Flags:  discordgo.MessageFlagsEphemeral,
		},
	}
	if len(actions.Components) != 0 {
		resp.Data.Components = []discordgo.MessageComponent{actions}
	}

	if err = s.InteractionRespond(i.Interaction, resp); err != nil {
		return errors.NewError("レスポンスを送信できません", err)
	}

	return nil
}
