package quiz

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/kifuneso-bot/internal/cmd"
	"github.com/techstart35/kifuneso-bot/internal/color"
	"github.com/techstart35/kifuneso-bot/internal/errors"
)

// パネルを送信します
func SendPanel(s *discordgo.Session, m *discordgo.MessageCreate) error {
	if m.Author.Bot {
		return nil
	}

	btn1 := discordgo.Button{
		Style:    discordgo.PrimaryButton,
		CustomID: cmd.QuizButton.Start.InteractionID,
		Label:    cmd.QuizButton.Start.Label,
	}

	actions := discordgo.ActionsRow{
		Components: []discordgo.MessageComponent{btn1},
	}

	description := `
## RED GINGER Quiz 🎁

プロジェクトに関するクイズは全部で10問です。是非、全問正解に向けて挑戦してみてください！
（何度でも挑戦可能です）

__※途中で閉じるとやり直しとなります。ご注意ください！__
`

	embed := &discordgo.MessageEmbed{
		Description: fmt.Sprintf(
			description,
		),
		Color: color.Red,
	}

	// 新規のパネルを作成します
	messageSend := &discordgo.MessageSend{
		Components: []discordgo.MessageComponent{actions},
		Embed:      embed,
	}

	_, err := s.ChannelMessageSendComplex(m.ChannelID, messageSend)
	if err != nil {
		return errors.NewError("パネルメッセージを送信できません", err)
	}

	// コマンドメッセージを削除
	if err = s.ChannelMessageDelete(m.ChannelID, m.ID); err != nil {
		return errors.NewError("コマンドメッセージを削除できません", err)
	}

	return nil
}
