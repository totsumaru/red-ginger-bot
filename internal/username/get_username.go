package username

import (
	"github.com/bwmarrin/discordgo"
)

// ユーザー名を取得します
func GetUserNameFromMessageCreate(m *discordgo.MessageCreate) string {
	username := m.Author.Username
	if m.Author.GlobalName != "" {
		username = m.Author.GlobalName
	}
	if m.Member.Nick != "" {
		username = m.Member.Nick
	}

	return username
}
