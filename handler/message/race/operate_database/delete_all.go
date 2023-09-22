package operate_database

import (
	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/kifuneso-bot/internal/db"
	"github.com/techstart35/kifuneso-bot/internal/errors"
	"github.com/techstart35/kifuneso-bot/internal/id"
)

// 全てのレコードを削除します
func DeleteAll(s *discordgo.Session, m *discordgo.MessageCreate) error {
	if m.Author.ID != id.UserID().TOTSUMARU {
		return nil
	}

	r := db.Race{}
	if err := r.DeleteAll(); err != nil {
		return errors.NewError("全てを削除できません", err)
	}

	if _, err := s.ChannelMessageSend(m.ChannelID, "完了"); err != nil {
		return errors.NewError("メッセージを送信できません", err)
	}

	return nil
}
