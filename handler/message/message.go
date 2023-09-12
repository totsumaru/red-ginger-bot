package message

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/kifuneso-bot/handler/message/info"
	"github.com/techstart35/kifuneso-bot/handler/message/quiz"
	"github.com/techstart35/kifuneso-bot/handler/message/race"
	"github.com/techstart35/kifuneso-bot/handler/message/race/operate_database"
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
	case cmd.CMD_INFO:
		if err := info.ProjectInfo(s, m); err != nil {
			errors.SendErrMsg(s, err, m.Author)
			return
		}
	case cmd.CMD_Story:
		if err := info.SendStoryPanel(s, m); err != nil {
			errors.SendErrMsg(s, err, m.Author)
			return
		}
	case cmd.CMD_QuizPanel:
		if err := quiz.SendPanel(s, m); err != nil {
			errors.SendErrMsg(s, err, m.Author)
			return
		}
	case cmd.CMD_Race:
		if err := race.SendRace(s, m); err != nil {
			errors.SendErrMsg(s, err, m.Author)
			return
		}
	}

	if strings.Contains(m.Content, cmd.CMD_DeleteRecordByID) {
		if err := operate_database.DeleteByID(s, m); err != nil {
			errors.SendErrMsg(s, err, m.Author)
			return
		}
	}

	if err := slot.AddFiveTicketRole(s, m); err != nil {
		errors.SendErrMsg(s, err, m.Author)
		return
	}

	return
}
