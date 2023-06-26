package slot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/kifuneso-bot/internal/errors"
	"github.com/techstart35/kifuneso-bot/internal/id"
)

// 5回追加したロールを取得するmapです
// 現在のロール: 5回追加したロール となります
var FiveAddedRole = map[string]string{
	id.RoleID().SLOT_1_TICKET:  id.RoleID().SLOT_6_TICKET,
	id.RoleID().SLOT_2_TICKET:  id.RoleID().SLOT_7_TICKET,
	id.RoleID().SLOT_3_TICKET:  id.RoleID().SLOT_8_TICKET,
	id.RoleID().SLOT_4_TICKET:  id.RoleID().SLOT_9_TICKET,
	id.RoleID().SLOT_5_TICKET:  id.RoleID().SLOT_10_TICKET,
	id.RoleID().SLOT_6_TICKET:  id.RoleID().SLOT_11_TICKET,
	id.RoleID().SLOT_7_TICKET:  id.RoleID().SLOT_12_TICKET,
	id.RoleID().SLOT_8_TICKET:  id.RoleID().SLOT_13_TICKET,
	id.RoleID().SLOT_9_TICKET:  id.RoleID().SLOT_14_TICKET,
	id.RoleID().SLOT_10_TICKET: id.RoleID().SLOT_15_TICKET,
	id.RoleID().SLOT_11_TICKET: id.RoleID().SLOT_15_TICKET,
	id.RoleID().SLOT_12_TICKET: id.RoleID().SLOT_15_TICKET,
	id.RoleID().SLOT_13_TICKET: id.RoleID().SLOT_15_TICKET,
	id.RoleID().SLOT_14_TICKET: id.RoleID().SLOT_15_TICKET,
	id.RoleID().SLOT_15_TICKET: id.RoleID().SLOT_15_TICKET,
}

// 次の当たりロールを取得するmapです
// 現在の当たりロール: 新しい当たりロール となります
var NewAtariRole = map[string]string{
	id.RoleID().ATARI_1: id.RoleID().ATARI_2,
	id.RoleID().ATARI_2: id.RoleID().ATARI_3,
	id.RoleID().ATARI_3: id.RoleID().ATARI_4,
	id.RoleID().ATARI_4: id.RoleID().ATARI_5,
	id.RoleID().ATARI_5: id.RoleID().ATARI_6,
	id.RoleID().ATARI_6: id.RoleID().ATARI_7,
	id.RoleID().ATARI_7: id.RoleID().ATARI_8,
	id.RoleID().ATARI_8: id.RoleID().ATARI_9,
	id.RoleID().ATARI_9: id.RoleID().ATARI_10,
}

// 今のロールを削除し、+5したロールを付け直します
func UpdateRoleToPlus5(s *discordgo.Session, guildID, userID string, currentRoles []string) error {
	for _, role := range currentRoles {
		if afterRoleID, ok := FiveAddedRole[role]; ok {
			// 現在のロールを削除します
			if err := s.GuildMemberRoleRemove(guildID, userID, role); err != nil {
				return errors.NewError("現在の回数ロールを削除できません", err)
			}
			// 新規ロールを付与します
			if err := s.GuildMemberRoleAdd(guildID, userID, afterRoleID); err != nil {
				return errors.NewError("新規の回数ロールを付与できません", err)
			}
		}
	}

	return nil
}

// チケットロールかどうかを判定します
func IsSlotTicketRoleContainsAddedRole(roleID string) bool {
	switch roleID {
	case id.RoleID().SLOT_1_TICKET,
		id.RoleID().SLOT_2_TICKET,
		id.RoleID().SLOT_3_TICKET,
		id.RoleID().SLOT_4_TICKET,
		id.RoleID().SLOT_5_TICKET,
		id.RoleID().SLOT_6_TICKET,
		id.RoleID().SLOT_7_TICKET,
		id.RoleID().SLOT_8_TICKET,
		id.RoleID().SLOT_9_TICKET,
		id.RoleID().SLOT_10_TICKET,
		id.RoleID().SLOT_11_TICKET,
		id.RoleID().SLOT_12_TICKET,
		id.RoleID().SLOT_13_TICKET,
		id.RoleID().SLOT_14_TICKET,
		id.RoleID().SLOT_15_TICKET,
		id.RoleID().SLOT_ADDED:

		return true
	}

	return false
}

// 当たりロールかどうかを判定します
func IsAtariRole(roleID string) bool {
	switch roleID {
	case id.RoleID().ATARI_1,
		id.RoleID().ATARI_2,
		id.RoleID().ATARI_3,
		id.RoleID().ATARI_4,
		id.RoleID().ATARI_5,
		id.RoleID().ATARI_6,
		id.RoleID().ATARI_7,
		id.RoleID().ATARI_8,
		id.RoleID().ATARI_9,
		id.RoleID().ATARI_10:

		return true
	}

	return false
}
