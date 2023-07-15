package slot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/kifuneso-bot/handler/interaction/utils"
	"github.com/techstart35/kifuneso-bot/internal/cmd"
	"github.com/techstart35/kifuneso-bot/internal/color"
	"github.com/techstart35/kifuneso-bot/internal/errors"
	"github.com/techstart35/kifuneso-bot/internal/id"
	"github.com/techstart35/kifuneso-bot/internal/slot"
	"strings"
)

// 3回目の数字を送信します
func SendThirdNumber(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	editFunc, err := utils.SendInteractionWaitingMessage(s, i, true, true)
	if err != nil {
		return errors.NewError("Waitingメッセージが送信できません")
	}

	lastDescription := strings.Replace(i.Message.Embeds[0].Description, "**", "", -1)
	lastDescription = strings.Replace(lastDescription, " ", "", -1)
	lastValue1 := strings.Split(lastDescription, "｜")[0]
	lastValue2 := strings.Split(lastDescription, "｜")[1]
	value := getEachValue(3, lastValue1, lastValue2)

	descriptionTmpl := `
%s

%s
`

	btn1 := discordgo.Button{
		Style:    discordgo.PrimaryButton,
		CustomID: cmd.Interaction_CustomID_Slot_Retry,
		Label:    "もう一回",
	}

	actions := discordgo.ActionsRow{
		Components: []discordgo.MessageComponent{btn1},
	}

	embed := &discordgo.MessageEmbed{
		Title:       Title,
		Description: "",
		Color:       color.Red,
	}

	var resultStr = lastValue1 + lastValue2 + value

	switch judgePrize(lastValue1, lastValue2, value) {
	case Prize_Lose:
		embed.Description = fmt.Sprintf(
			descriptionTmpl,
			"残念！",
			fmt.Sprintf(DescriptionTmpl, lastValue1, lastValue2, value),
		)
	case Prize_Big:
		embed.Description = fmt.Sprintf(
			descriptionTmpl,
			"大当たり🎉🎉🎉\nロールを付与しました！",
			fmt.Sprintf(DescriptionTmpl, lastValue1, lastValue2, value),
		)

		switch resultStr {
		case String_RedGinGer:
			embed.Image = &discordgo.MessageEmbedImage{
				URL: "https://cdn.discordapp.com/attachments/1103240223376293938/1122881805603835976/RGSLOT_REDGINGER.png",
			}
		case String_BeniShouGa:
			embed.Image = &discordgo.MessageEmbedImage{
				URL: "https://cdn.discordapp.com/attachments/1103240223376293938/1122881823878422669/RGSLOT_.png",
			}
		}

		newRoleID, err := updateBigPrizeRole(s, i)
		if err != nil {
			return errors.NewError("大当たりロールを更新できません", err)
		}

		if err = NoticeAtariToAdmin(s, i, newRoleID); err != nil {
			return errors.NewError("MODチャンネルに通知を送信できません", err)
		}
	case Prize_Small:
		embed.Description = fmt.Sprintf(
			descriptionTmpl,
			"小当たり🎉\n追加でもう5回遊べるよ！",
			fmt.Sprintf(DescriptionTmpl, lastValue1, lastValue2, value),
		)

		switch resultStr {
		case String_Red_3:
			embed.Image = &discordgo.MessageEmbedImage{
				URL: "https://cdn.discordapp.com/attachments/1103240223376293938/1122881687941025792/RGSLOT_RED.png",
			}
		case String_Gin_3:
			embed.Image = &discordgo.MessageEmbedImage{
				URL: "https://cdn.discordapp.com/attachments/1103240223376293938/1122881716068045023/RGSLOT_GIN.png",
			}
		case String_Ger_3:
			embed.Image = &discordgo.MessageEmbedImage{
				URL: "https://cdn.discordapp.com/attachments/1103240223376293938/1122881739166060644/RGSLOT_GER.png",
			}
		case String_Beni_3:
			embed.Image = &discordgo.MessageEmbedImage{
				URL: "https://cdn.discordapp.com/attachments/1103240223376293938/1122881757415493674/RGSLOT_.png",
			}
		case String_Shou_3:
			embed.Image = &discordgo.MessageEmbedImage{
				URL: "https://cdn.discordapp.com/attachments/1103240223376293938/1122881775505518612/RGSLOT_.png",
			}
		case String_Ga_3:
			embed.Image = &discordgo.MessageEmbedImage{
				URL: "https://cdn.discordapp.com/attachments/1103240223376293938/1122881790554689546/RGSLOT_.png",
			}
		}

		if err := slot.UpdateRoleToPlus5(s, i.GuildID, i.Member.User.ID, i.Member.Roles); err != nil {
			return errors.NewError("小当たりでロールを更新できません", err)
		}
	case Prize_OneMore:
		embed.Description = fmt.Sprintf(
			descriptionTmpl,
			"🍒が出たからもう一回遊べるよ！",
			fmt.Sprintf(DescriptionTmpl, lastValue1, lastValue2, value),
		)

		if err := slot.UpdateRoleToPlus1(s, i.GuildID, i.Member.User.ID, i.Member.Roles); err != nil {
			return errors.NewError("チェリー当たりでロールを更新できません", err)
		}
	case Prize_Secret:
		embed.Description = fmt.Sprintf(
			descriptionTmpl,
			"超大当たり🎉🎉🎉🎉\n当たりロールを獲得しました！\n\nもう10回遊べます！！",
			fmt.Sprintf(DescriptionTmpl, lastValue1, lastValue2, value),
		)
		embed.Image = &discordgo.MessageEmbedImage{
			URL: "https://cdn.discordapp.com/attachments/1103240223376293938/1123517363992666132/RGSLOT_GAKU.png",
		}

		newRoleID, err := updateBigPrizeRole(s, i)
		if err != nil {
			return errors.NewError("大当たりロールを更新できません", err)
		}

		if err = slot.UpdateRoleToPlus10(s, i.GuildID, i.Member.User.ID, i.Member.Roles); err != nil {
			return errors.NewError("???で+10回券ロールを更新できません", err)
		}

		if err = NoticeAtariToAdmin(s, i, newRoleID); err != nil {
			return errors.NewError("MODチャンネルに通知を送信できません", err)
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

// 大当たりロールを更新します
//
// 新しい当たりロールIDを返します。
func updateBigPrizeRole(s *discordgo.Session, i *discordgo.InteractionCreate) (string, error) {
	for _, currentAtariRole := range i.Member.Roles {
		// 現在1-9当たりの場合はロールの更新あり
		if newAtariRole, ok := slot.NewAtariRole[currentAtariRole]; ok {
			// 現在の当たりロールを削除
			if err := s.GuildMemberRoleRemove(i.GuildID, i.Member.User.ID, currentAtariRole); err != nil {
				return "", errors.NewError("ロールを削除できません", err)
			}
			// 新しい当たりロールを付与
			if err := s.GuildMemberRoleAdd(i.GuildID, i.Member.User.ID, newAtariRole); err != nil {
				return "", errors.NewError("ロールを削除できません", err)
			}

			return newAtariRole, nil
		}
	}

	// 初めての当たりの場合は、当たり1を付与
	if err := s.GuildMemberRoleAdd(i.GuildID, i.Member.User.ID, id.RoleID().ATARI_1); err != nil {
		return "", errors.NewError("ロールを削除できません", err)
	}

	return id.RoleID().ATARI_1, nil
}

// 当たり通知を管理者向けに送信します
func NoticeAtariToAdmin(
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
	newAtariRoleID string,
) error {
	description := `
**Slot当たり通知**

- %d回目
- <@%s>
`

	count := 0
	switch newAtariRoleID {
	case id.RoleID().ATARI_3:
		count = 3
	}

	embed := &discordgo.MessageEmbed{
		Description: fmt.Sprintf(description, count, i.Member.User.ID),
		Color:       color.Yellow,
		Author: &discordgo.MessageEmbedAuthor{
			IconURL: i.Member.User.AvatarURL(""),
			Name:    i.Member.User.Username,
		},
	}

	if _, err := s.ChannelMessageSendEmbed(id.ChannelID().MOD, embed); err != nil {
		return errors.NewError("MODチャンネルにメッセージを送信できません", err)
	}

	return nil
}
