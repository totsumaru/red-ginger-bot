package info

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/kifuneso-bot/internal/cmd"
	"github.com/techstart35/kifuneso-bot/internal/color"
	"github.com/techstart35/kifuneso-bot/internal/errors"
	"github.com/techstart35/kifuneso-bot/internal/id"
)

// #story(スレッド)に表示するメッセージを送信します
func SendStoryPanel(s *discordgo.Session, m *discordgo.MessageCreate) error {
	btn1 := discordgo.Button{
		Style:    discordgo.PrimaryButton,
		CustomID: cmd.Interaction_CustomID_ToStory2,
		Label:    "次へ",
	}

	btn2 := discordgo.Button{
		Style: discordgo.LinkButton,
		URL: id.GenerateMessageLink(
			m.GuildID,
			id.ChannelID().CHARACTER,
			id.MessageID().CHARACTER_GG,
		),
		Label: "キャラクター紹介",
	}

	actions := discordgo.ActionsRow{
		Components: []discordgo.MessageComponent{btn1, btn2},
	}

	description := `
この物語の主人公は、宇宙で運送業を営む「**GG**」と彼の友人であるロボット「**ガク**」です。

舞台は__2200年代__、人類は地球から宇宙へと生活の場を広げていました。
`

	embed := &discordgo.MessageEmbed{
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: "https://cdn.discordapp.com/attachments/1103240223376293938/1128160753988407326/0ef1cb4ead41f82e.png",
		},
		Title: "STORY（1/7）",
		Description: fmt.Sprintf(
			description,
		),
		Color: color.Red,
	}

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
