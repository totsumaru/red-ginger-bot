package message_create

import "github.com/bwmarrin/discordgo"

// メッセージが作成された時のハンドラです
func Handler(s *discordgo.Session, m *discordgo.MessageCreate) {
	MsgToCSV(s, m)
}
