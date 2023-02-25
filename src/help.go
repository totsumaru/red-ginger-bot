package src

import (
	"github.com/bwmarrin/discordgo"
)

// helpを送信します
func Help(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Content != CMDHelp {
		return
	}

	msg := `
!r-help
!r-send [タイトル(なければ"none")] [チャンネルリンク] [メッセージID]
`

	req := &discordgo.MessageEmbed{
		Title:       "HELP",
		Description: msg,
	}
	_, err := s.ChannelMessageSendEmbed(m.ChannelID, req)
	if err != nil {
		panic(err)
	}
}
