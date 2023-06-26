package slot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/kifuneso-bot/internal/color"
	"github.com/techstart35/kifuneso-bot/internal/errors"
	"github.com/techstart35/kifuneso-bot/internal/id"
	"github.com/techstart35/kifuneso-bot/internal/slot"
)

// slotの開始メッセージを送信します
func SendStartMessage(
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
	isUpdateMessage bool,
) error {
	var currentTicketRole string

	for _, role := range i.Member.Roles {
		if slot.IsSlotTicketRole(role) {
			currentTicketRole = role
		}
	}

	// チケットを持っていない場合はエラーを送信して終了
	if currentTicketRole == "" {
		if err := sendNotHaveTicketErrorMessage(s, i, isUpdateMessage); err != nil {
			return errors.NewError("チケット未保持メッセージを送信できません", err)
		}

		return nil
	}

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

	// まだチケットが残っている場合は新規ロールを付与します
	if newRoleID, ok := slot.MinusTicketRole[currentTicketRole]; ok {
		if err := s.GuildMemberRoleAdd(i.GuildID, i.Member.User.ID, newRoleID); err != nil {
			return errors.NewError("新規の回数ロールを付与できません", err)
		}
	}

	// 現在のチケットロールを削除します
	if err := s.GuildMemberRoleRemove(i.GuildID, i.Member.User.ID, currentTicketRole); err != nil {
		return errors.NewError("現在の回数ロールを削除できません", err)
	}

	return nil
}

// チケットを保持していないエラーメッセージを送信します
func sendNotHaveTicketErrorMessage(
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
	isUpdateMessage bool,
) error {
	description := `
本日の回数を使い切ってしまいました🥲

- 毎日5回分のチケットがもらえます
- どこかのチャンネルでコメントすると、さらに5回分のチケットがもらえます（1日1回まで）

本日まだコメントしていない人は、ぜひ <#%s> でコメントしてみてください！
`

	embed := &discordgo.MessageEmbed{
		Title:       "ERROR",
		Description: fmt.Sprintf(description, id.ChannelID().JP_CHAT),
		Color:       color.Red,
	}

	responseType := discordgo.InteractionResponseChannelMessageWithSource
	if isUpdateMessage {
		responseType = discordgo.InteractionResponseUpdateMessage
	}

	resp := &discordgo.InteractionResponse{
		Type: responseType,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{embed},
			Flags:  discordgo.MessageFlagsEphemeral,
		},
	}

	if err := s.InteractionRespond(i.Interaction, resp); err != nil {
		return errors.NewError("レスポンスを送信できません", err)
	}

	return nil
}
