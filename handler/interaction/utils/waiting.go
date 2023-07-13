package utils

import "github.com/bwmarrin/discordgo"

// Interactionのdeferメッセージを送信します
func SendInteractionWaitingMessage(
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
	isUpdate bool,
) error {
	responseType := discordgo.InteractionResponseDeferredChannelMessageWithSource
	if isUpdate {
		responseType = discordgo.InteractionResponseDeferredMessageUpdate
	}

	resp := &discordgo.InteractionResponse{
		Type: responseType,
		Data: &discordgo.InteractionResponseData{
			Flags: discordgo.MessageFlagsEphemeral,
		},
	}

	return s.InteractionRespond(i.Interaction, resp)
}
