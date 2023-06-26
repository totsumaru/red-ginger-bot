package message

import (
	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/kifuneso-bot/handler/message/slot"
	"github.com/techstart35/kifuneso-bot/internal/cmd"
	"github.com/techstart35/kifuneso-bot/internal/errors"
)

// メッセージが作成された時のハンドラーです
func MessageCreateHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	switch m.Content {
	case cmd.CMD_Send_Slot_Panel:
		if err := slot.SendPanel(s, m, ""); err != nil {
			errors.SendErrMsg(s, err, m.Author)
			return
		}
	case cmd.CMD_Reset_Slot_Role:
		if err := slot.ResetAndAddDailyRole(s, m); err != nil {
			errors.SendErrMsg(s, err, m.Author)
			return
		}
	}

	return
}
