package operate_database

import (
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/kifuneso-bot/internal/db"
	"github.com/techstart35/kifuneso-bot/internal/errors"
	"github.com/techstart35/kifuneso-bot/internal/id"
)

// 指定のIDのポイントを追加します
func AddPoint(s *discordgo.Session, m *discordgo.MessageCreate) error {
	if m.Author.ID != id.UserID().TOTSUMARU {
		return nil
	}

	targetID := strings.Split(m.Content, " ")[1]
	addPoint, err := strconv.Atoi(strings.Split(m.Content, " ")[2])
	if err != nil {
		return errors.NewError("ポイントが不正です", err)
	}

	r := db.Race{
		ID:    targetID,
		Point: addPoint,
	}
	if err = r.Upsert(); err != nil {
		return errors.NewError("ポイントを加算できません", err)
	}

	if _, err = s.ChannelMessageSend(m.ChannelID, "完了"); err != nil {
		return errors.NewError("メッセージを送信できません", err)
	}

	return nil
}
