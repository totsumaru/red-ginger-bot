package slot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/kifuneso-bot/internal/errors"
	"github.com/techstart35/kifuneso-bot/internal/id"
	"github.com/techstart35/kifuneso-bot/internal/slot"
)

// コメントをしたユーザーに5回分のロールを付与します
//
// 1日1回までです。
func AddFiveTicketRole(s *discordgo.Session, m *discordgo.MessageCreate) error {
	// webhook対策として、nilチェックを行います
	if m.Member == nil {
		return nil
	}

	// すでに付与した人はreturnで終了します
	for _, role := range m.Member.Roles {
		if role == id.RoleID().SLOT_ADDED {
			return nil
		}
	}

	// チケットロールを更新します
	if err := slot.UpdateRoleToPlus5(s, m.GuildID, m.Author.ID, m.Member.Roles); err != nil {
		return errors.NewError("チケットロールを更新できません", err)
	}

	// ユーザーにAddedロールを付与します
	if err := s.GuildMemberRoleAdd(m.GuildID, m.Author.ID, id.RoleID().SLOT_ADDED); err != nil {
		return errors.NewError("ロールを付与できません", err)
	}

	return nil
}
