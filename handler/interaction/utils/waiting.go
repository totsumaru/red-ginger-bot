package utils

import (
	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/kifuneso-bot/internal/errors"
)

// レスポンスのEdit関数を返します
type EditFunc func(interaction *discordgo.Interaction, newresp *discordgo.WebhookEdit) (*discordgo.Message, error)

// Interactionのdeferメッセージを送信します
func SendInteractionWaitingMessage(
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
	isUpdate bool,
	isEphemeral bool,
) (EditFunc, error) {
	responseType := discordgo.InteractionResponseDeferredChannelMessageWithSource
	if isUpdate {
		responseType = discordgo.InteractionResponseDeferredMessageUpdate
	}

	resp := &discordgo.InteractionResponse{
		Type: responseType,
	}

	if isEphemeral {
		resp.Data.Flags = discordgo.MessageFlagsEphemeral
	}

	if err := s.InteractionRespond(i.Interaction, resp); err != nil {
		return nil, errors.NewError("インタラクションのレスポンスを送信できません", err)
	}

	return s.InteractionResponseEdit, nil
}
