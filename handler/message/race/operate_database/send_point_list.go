package operate_database

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/kifuneso-bot/internal/db"
	"github.com/techstart35/kifuneso-bot/internal/errors"
	"github.com/techstart35/kifuneso-bot/internal/id"
)

// ポイントの一覧を送信します
func SendPointList(s *discordgo.Session, m *discordgo.MessageCreate) error {
	if m.Author.ID != id.UserID().TOTSUMARU {
		return nil
	}

	r := db.Race{}
	records, err := r.FindAll()
	if err != nil {
		return errors.NewError("全てのデータを取得できません", err)
	}

	list := make([]string, 0)
	for _, record := range records {
		member, err := s.GuildMember(m.GuildID, record.ID)
		if err != nil {
			return errors.NewError("メンバーを取得できません", err)
		}
		name := member.Nick
		if name == "" {
			name = member.User.Username
		}
		list = append(list, fmt.Sprintf("%s: %dpt", name, r.Point))
	}

	_, err = s.ChannelMessageSend(
		id.ChannelID().RANKING,
		strings.Join(list, "\n"),
	)
	if err != nil {
		return errors.NewError("メッセージを送信できません", err)
	}

	return nil
}
