package slot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/kifuneso-bot/internal/cmd"
	"github.com/techstart35/kifuneso-bot/internal/color"
	"github.com/techstart35/kifuneso-bot/internal/errors"
	"github.com/techstart35/kifuneso-bot/internal/id"
	"github.com/techstart35/kifuneso-bot/internal/slot"
	"strings"
)

// 3回目の数字を送信します
func SendThirdNumber(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	lastDescription := strings.Replace(i.Message.Embeds[0].Description, "**", "", -1)
	lastDescription = strings.Replace(lastDescription, " ", "", -1)
	lastValue1 := strings.Split(lastDescription, "｜")[0]
	lastValue2 := strings.Split(lastDescription, "｜")[1]
	value := getRandomValue(3, lastValue1, lastValue2)

	descriptionTmpl := `
%s

%s
`

	description := ""

	switch judgePrize(lastValue1, lastValue2, value) {
	case Prize_Lose:
		description = fmt.Sprintf(
			descriptionTmpl,
			"残念！",
			fmt.Sprintf(DescriptionTmpl, lastValue1, lastValue2, value),
		)
	case Prize_Big:
		description = fmt.Sprintf(
			descriptionTmpl,
			"大当たり🎉🎉🎉\nロールを付与しました！",
			fmt.Sprintf(DescriptionTmpl, lastValue1, lastValue2, value),
		)

		if err := updateBigPrizeRole(s, i); err != nil {
			return errors.NewError("大当たりロールを更新できません", err)
		}
	case Prize_Small:
		description = fmt.Sprintf(
			descriptionTmpl,
			"小当たり🎉\n追加でもう5回遊べるよ！",
			fmt.Sprintf(DescriptionTmpl, lastValue1, lastValue2, value),
		)

		if err := slot.UpdateRoleToPlus5(s, i.GuildID, i.Member.User.ID, i.Member.Roles); err != nil {
			return errors.NewError("小当たりでロールを更新できません", err)
		}
	case Prize_OneMore:
		description = fmt.Sprintf(
			descriptionTmpl,
			"🍒が出たからもう一回遊べるよ！",
			fmt.Sprintf(DescriptionTmpl, lastValue1, lastValue2, value),
		)
	}

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
		Description: description,
		Color:       color.Red,
	}

	resp := &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseUpdateMessage,
		Data: &discordgo.InteractionResponseData{
			Components: []discordgo.MessageComponent{actions},
			Embeds:     []*discordgo.MessageEmbed{embed},
			Flags:      discordgo.MessageFlagsEphemeral,
		},
	}

	if err := s.InteractionRespond(i.Interaction, resp); err != nil {
		return errors.NewError("レスポンスを送信できません", err)
	}

	return nil
}

// 大当たりロールを更新します
func updateBigPrizeRole(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	for _, currentAtariRole := range i.Member.Roles {
		// 現在1-9当たりの場合はロールの更新あり
		if newAtariRole, ok := slot.NewAtariRole[currentAtariRole]; ok {
			// 現在の当たりロールを削除
			if err := s.GuildMemberRoleRemove(i.GuildID, i.Member.User.ID, currentAtariRole); err != nil {
				return errors.NewError("ロールを削除できません", err)
			}
			// 新しい当たりロールを付与
			if err := s.GuildMemberRoleAdd(i.GuildID, i.Member.User.ID, newAtariRole); err != nil {
				return errors.NewError("ロールを削除できません", err)
			}

			return nil
		}
	}

	// 初めての当たりの場合は、当たり1を付与
	if err := s.GuildMemberRoleAdd(i.GuildID, i.Member.User.ID, id.RoleID().ATARI_1); err != nil {
		return errors.NewError("ロールを削除できません", err)
	}

	return nil
}
