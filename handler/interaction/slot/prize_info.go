package slot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/kifuneso-bot/internal/color"
	"github.com/techstart35/kifuneso-bot/internal/errors"
)

// 当たりの情報を送信します
func SendPrizeInfo(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	description := `
☑️大当たり
（当たりロールGET）
- RED｜GIN｜GER
- 紅｜生｜姜

☑️小当たり
（もう5回遊べます。当たりロールは無し）
- RED｜RED｜RED
- GIN｜GIN｜GIN
- GER｜GER｜GER
- 紅｜紅｜紅
- 生｜生｜生
- 姜｜姜｜姜

☑️その他
- 🍒が出ると、もう一回遊べます
-----------------
`

	embed := &discordgo.MessageEmbed{
		Description: description,
		Color:       color.Red,
	}

	resp := &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
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
