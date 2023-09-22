package operate_database

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/kifuneso-bot/internal/db"
	"github.com/techstart35/kifuneso-bot/internal/errors"
	"github.com/techstart35/kifuneso-bot/internal/id"
)

// 指定のIDのレコードを削除します
func DeleteByID(s *discordgo.Session, m *discordgo.MessageCreate) error {
	if m.Author.ID != id.UserID().TOTSUMARU {
		return nil
	}

	targetID := strings.Split(m.Content, " ")[1]

	r := db.Race{
		ID: targetID,
	}
	if err := r.DeleteByID(); err != nil {
		return errors.NewError("削除できません", err)
	}

	if _, err := s.ChannelMessageSend(m.ChannelID, "完了"); err != nil {
		return errors.NewError("メッセージを送信できません", err)
	}

	return nil
}