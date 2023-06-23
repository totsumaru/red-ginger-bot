package slot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/kifuneso-bot/internal/color"
	"github.com/techstart35/kifuneso-bot/internal/errors"
	"strconv"
	"strings"
)

// 3回目の数字を送信します
func SendThirdNumber(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	num := getRandomNum()
	lastDescription := strings.Replace(i.Message.Embeds[0].Description, "**", "", -1)
	lastDescription = strings.Replace(lastDescription, " ", "", -1)
	lastNum1 := strings.Split(lastDescription, "｜")[0]
	lastNum2 := strings.Split(lastDescription, "｜")[1]

	description := ""

	if lastNum1 == lastNum2 && lastNum2 == strconv.Itoa(int(num)) {
		description = fmt.Sprintf(
			"%s\n\n%s",
			fmt.Sprintf(DescriptionTmpl, lastNum1, lastNum2, strconv.Itoa(int(num))),
			"おめでとうございます🎉\nロールを付与しました！",
		)
	} else {
		description = fmt.Sprintf(
			"%s\n\n%s",
			fmt.Sprintf(DescriptionTmpl, lastNum1, lastNum2, strconv.Itoa(int(num))),
			"残念...また明日チャレンジしてみてね！",
		)
	}

	embed := &discordgo.MessageEmbed{
		Title:       Title,
		Description: description,
		Color:       color.Red,
	}

	resp := &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseUpdateMessage,
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
