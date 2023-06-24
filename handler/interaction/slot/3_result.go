package slot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/kifuneso-bot/internal/cmd"
	"github.com/techstart35/kifuneso-bot/internal/color"
	"github.com/techstart35/kifuneso-bot/internal/errors"
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
	case Prize_AL:
		description = fmt.Sprintf(
			descriptionTmpl,
			"当たり-AL🎉🎉\nロールを付与しました！",
			fmt.Sprintf(DescriptionTmpl, lastValue1, lastValue2, value),
		)
	case Prize_Small:
		description = fmt.Sprintf(
			descriptionTmpl,
			"小当たり🎉\nロールを付与しました！",
			fmt.Sprintf(DescriptionTmpl, lastValue1, lastValue2, value),
		)
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
