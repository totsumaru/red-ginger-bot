package slot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/kifuneso-bot/handler/interaction/utils"
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
	editFunc, err := utils.SendInteractionWaitingMessage(s, i, isUpdateMessage, true)
	if err != nil {
		return errors.NewError("Waitingメッセージが送信できません")
	}

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

残り回数: **%d**
`

	embed := &discordgo.MessageEmbed{
		Title:       Title,
		Description: fmt.Sprintf(description, numOfLeft(i.Member)),
		Color:       color.Red,
		Image: &discordgo.MessageEmbedImage{
			URL: "https://cdn.discordapp.com/attachments/1103240223376293938/1122881161769795714/b57d42376b9173e2.gif",
		},
	}

	// 現在のチケットロールを削除します
	if err = s.GuildMemberRoleRemove(i.GuildID, i.Member.User.ID, currentTicketRole); err != nil {
		return errors.NewError("現在の回数ロールを削除できません", err)
	}

	// まだチケットが残っている場合は新規ロールを付与します
	if newRoleID, ok := slot.MinusTicketRole[currentTicketRole]; ok {
		if err := s.GuildMemberRoleAdd(i.GuildID, i.Member.User.ID, newRoleID); err != nil {
			return errors.NewError("新規の回数ロールを付与できません", err)
		}
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

// チケットを保持していないエラーメッセージを送信します
func sendNotHaveTicketErrorMessage(
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
	isUpdateMessage bool,
) error {
	description := `
本日の回数を使い切ってしまいました..

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

// slotの残り回数を取得します
func numOfLeft(member *discordgo.Member) int {
	for _, role := range member.Roles {
		switch role {
		case id.RoleID().SLOT_1_TICKET:
			return 1
		case id.RoleID().SLOT_2_TICKET:
			return 2
		case id.RoleID().SLOT_3_TICKET:
			return 3
		case id.RoleID().SLOT_4_TICKET:
			return 4
		case id.RoleID().SLOT_5_TICKET:
			return 5
		case id.RoleID().SLOT_6_TICKET:
			return 6
		case id.RoleID().SLOT_7_TICKET:
			return 7
		case id.RoleID().SLOT_8_TICKET:
			return 8
		case id.RoleID().SLOT_9_TICKET:
			return 9
		case id.RoleID().SLOT_10_TICKET:
			return 10
		case id.RoleID().SLOT_11_TICKET:
			return 11
		case id.RoleID().SLOT_12_TICKET:
			return 12
		case id.RoleID().SLOT_13_TICKET:
			return 13
		case id.RoleID().SLOT_14_TICKET:
			return 14
		case id.RoleID().SLOT_15_TICKET:
			return 15
		}
	}

	return 0
}
