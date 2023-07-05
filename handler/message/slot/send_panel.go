package slot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/kifuneso-bot/internal/cmd"
	"github.com/techstart35/kifuneso-bot/internal/color"
	"github.com/techstart35/kifuneso-bot/internal/errors"
	"strings"
)

// パネルを送信します
//
// 新規でパネルを送信する場合は`currentPanelURL`を空に、
// パネルを更新する場合は、現在のパネルのURLを入れてください。
func SendPanel(s *discordgo.Session, m *discordgo.MessageCreate, currentPanelURL string) error {
	btn1 := discordgo.Button{
		Style:    discordgo.PrimaryButton,
		CustomID: cmd.Interaction_CustomID_Slot_Start,
		Label:    "スロットを始める",
	}

	btn2 := discordgo.Button{
		Style:    discordgo.SecondaryButton,
		CustomID: cmd.Interaction_CustomID_Slot_PrizeInfo,
		Label:    "当たりの種類",
	}

	actions := discordgo.ActionsRow{
		Components: []discordgo.MessageComponent{btn1, btn2},
	}

	description := `
毎日スロットが回せます！
楽しんでください🎉

- 毎日5回遊べます
- 1回でもコメントをすると、さらに5回遊べます(チャンネルはどこでもOK）
- 遊べる回数は毎日リセットされます
- 当たりの種類は下のボタンから確認できます
`

	embed := &discordgo.MessageEmbed{
		Image: &discordgo.MessageEmbedImage{
			URL: "https://cdn.discordapp.com/attachments/1103240223376293938/1122880972094984304/4eb3d3bc86865188.png",
		},
		Title: "RG SLOT",
		Description: fmt.Sprintf(
			description,
		),
		Color: color.Red,
	}

	if currentPanelURL == "" {
		// 新規のパネルを作成します
		messageSend := &discordgo.MessageSend{
			Components: []discordgo.MessageComponent{actions},
			Embed:      embed,
		}

		_, err := s.ChannelMessageSendComplex(m.ChannelID, messageSend)
		if err != nil {
			return errors.NewError("パネルメッセージを送信できません", err)
		}
	} else {
		// パネルを更新します

		// URL例: https://discord.com/channels/1067806759034572870/1067807967950422096/1116242093434732595
		replaced := strings.Replace(currentPanelURL, "https://discord.com/channels/", "", -1)
		ids := strings.Split(replaced, "/")

		currentPanelChannelID := ids[1]
		currentPanelMessageID := ids[2]

		edit := &discordgo.MessageEdit{
			ID:         currentPanelMessageID,
			Channel:    currentPanelChannelID,
			Components: []discordgo.MessageComponent{actions},
			Embed:      embed,
		}

		_, err := s.ChannelMessageEditComplex(edit)
		if err != nil {
			return errors.NewError("パネルを更新できません", err)
		}
	}

	// コマンドメッセージを削除
	if err := s.ChannelMessageDelete(m.ChannelID, m.ID); err != nil {
		return errors.NewError("コマンドメッセージを削除できません", err)
	}

	return nil
}
