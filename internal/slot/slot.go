package slot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/kifuneso-bot/internal/errors"
	"github.com/techstart35/kifuneso-bot/internal/id"
)

// 1回追加したチケットロールを取得するmapです
//
// 現在のロール: 5回追加したロール となります
var OneAddedTicketRole = map[string]string{
	id.RoleID().SLOT_1_TICKET:  id.RoleID().SLOT_2_TICKET,
	id.RoleID().SLOT_2_TICKET:  id.RoleID().SLOT_3_TICKET,
	id.RoleID().SLOT_3_TICKET:  id.RoleID().SLOT_4_TICKET,
	id.RoleID().SLOT_4_TICKET:  id.RoleID().SLOT_5_TICKET,
	id.RoleID().SLOT_5_TICKET:  id.RoleID().SLOT_6_TICKET,
	id.RoleID().SLOT_6_TICKET:  id.RoleID().SLOT_7_TICKET,
	id.RoleID().SLOT_7_TICKET:  id.RoleID().SLOT_8_TICKET,
	id.RoleID().SLOT_8_TICKET:  id.RoleID().SLOT_9_TICKET,
	id.RoleID().SLOT_9_TICKET:  id.RoleID().SLOT_10_TICKET,
	id.RoleID().SLOT_10_TICKET: id.RoleID().SLOT_11_TICKET,
	id.RoleID().SLOT_11_TICKET: id.RoleID().SLOT_12_TICKET,
	id.RoleID().SLOT_12_TICKET: id.RoleID().SLOT_13_TICKET,
	id.RoleID().SLOT_13_TICKET: id.RoleID().SLOT_14_TICKET,
	id.RoleID().SLOT_14_TICKET: id.RoleID().SLOT_15_TICKET,
	id.RoleID().SLOT_15_TICKET: id.RoleID().SLOT_15_TICKET,
}

// 5回追加したチケットロールを取得するmapです
//
// 現在のロール: 5回追加したロール となります
var FiveAddedTicketRole = map[string]string{
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

// 10回追加したチケットロールを取得するmapです
//
// 現在のロール: 10回追加したロール となります
var TenAddedTicketRole = map[string]string{
	id.RoleID().SLOT_1_TICKET:  id.RoleID().SLOT_11_TICKET,
	id.RoleID().SLOT_2_TICKET:  id.RoleID().SLOT_12_TICKET,
	id.RoleID().SLOT_3_TICKET:  id.RoleID().SLOT_13_TICKET,
	id.RoleID().SLOT_4_TICKET:  id.RoleID().SLOT_14_TICKET,
	id.RoleID().SLOT_5_TICKET:  id.RoleID().SLOT_15_TICKET,
	id.RoleID().SLOT_6_TICKET:  id.RoleID().SLOT_15_TICKET,
	id.RoleID().SLOT_7_TICKET:  id.RoleID().SLOT_15_TICKET,
	id.RoleID().SLOT_8_TICKET:  id.RoleID().SLOT_15_TICKET,
	id.RoleID().SLOT_9_TICKET:  id.RoleID().SLOT_15_TICKET,
	id.RoleID().SLOT_10_TICKET: id.RoleID().SLOT_15_TICKET,
	id.RoleID().SLOT_11_TICKET: id.RoleID().SLOT_15_TICKET,
	id.RoleID().SLOT_12_TICKET: id.RoleID().SLOT_15_TICKET,
	id.RoleID().SLOT_13_TICKET: id.RoleID().SLOT_15_TICKET,
	id.RoleID().SLOT_14_TICKET: id.RoleID().SLOT_15_TICKET,
	id.RoleID().SLOT_15_TICKET: id.RoleID().SLOT_15_TICKET,
}

// 1回マイナスしたチケットロールを取得するmapです
var MinusTicketRole = map[string]string{
	id.RoleID().SLOT_15_TICKET: id.RoleID().SLOT_14_TICKET,
	id.RoleID().SLOT_14_TICKET: id.RoleID().SLOT_13_TICKET,
	id.RoleID().SLOT_13_TICKET: id.RoleID().SLOT_12_TICKET,
	id.RoleID().SLOT_12_TICKET: id.RoleID().SLOT_11_TICKET,
	id.RoleID().SLOT_11_TICKET: id.RoleID().SLOT_10_TICKET,
	id.RoleID().SLOT_10_TICKET: id.RoleID().SLOT_9_TICKET,
	id.RoleID().SLOT_9_TICKET:  id.RoleID().SLOT_8_TICKET,
	id.RoleID().SLOT_8_TICKET:  id.RoleID().SLOT_7_TICKET,
	id.RoleID().SLOT_7_TICKET:  id.RoleID().SLOT_6_TICKET,
	id.RoleID().SLOT_6_TICKET:  id.RoleID().SLOT_5_TICKET,
	id.RoleID().SLOT_5_TICKET:  id.RoleID().SLOT_4_TICKET,
	id.RoleID().SLOT_4_TICKET:  id.RoleID().SLOT_3_TICKET,
	id.RoleID().SLOT_3_TICKET:  id.RoleID().SLOT_2_TICKET,
	id.RoleID().SLOT_2_TICKET:  id.RoleID().SLOT_1_TICKET,
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

// 今のロールを削除し、+1したロールを付け直します
func UpdateRoleToPlus1(s *discordgo.Session, guildID, userID string, currentRoles []string) error {
	for _, role := range currentRoles {
		if afterRoleID, ok := OneAddedTicketRole[role]; ok {
			// 現在のロールを削除します
			if err := s.GuildMemberRoleRemove(guildID, userID, role); err != nil {
				return errors.NewError("現在の回数ロールを削除できません", err)
			}
			// 新規ロールを付与します
			if err := s.GuildMemberRoleAdd(guildID, userID, afterRoleID); err != nil {
				return errors.NewError("新規の回数ロールを付与できません", err)
			}

			return nil
		}
	}

	// 回数ロールを保持していない（使い切った）場合は、+5を付与する
	if err := s.GuildMemberRoleAdd(guildID, userID, id.RoleID().SLOT_1_TICKET); err != nil {
		return errors.NewError("新規の回数ロールを付与できません", err)
	}

	return nil
}

// 今のロールを削除し、+5したロールを付け直します
func UpdateRoleToPlus5(s *discordgo.Session, guildID, userID string, currentRoles []string) error {
	for _, role := range currentRoles {
		if afterRoleID, ok := FiveAddedTicketRole[role]; ok {
			// 現在のロールを削除します
			if err := s.GuildMemberRoleRemove(guildID, userID, role); err != nil {
				return errors.NewError("現在の回数ロールを削除できません", err)
			}
			// 新規ロールを付与します
			if err := s.GuildMemberRoleAdd(guildID, userID, afterRoleID); err != nil {
				return errors.NewError("新規の回数ロールを付与できません", err)
			}

			return nil
		}
	}

	// 回数ロールを保持していない（使い切った）場合は、+5を付与する
	if err := s.GuildMemberRoleAdd(guildID, userID, id.RoleID().SLOT_5_TICKET); err != nil {
		return errors.NewError("新規の回数ロールを付与できません", err)
	}

	return nil
}

// 今のロールを削除し、+10したロールを付け直します
func UpdateRoleToPlus10(s *discordgo.Session, guildID, userID string, currentRoles []string) error {
	for _, role := range currentRoles {
		if afterRoleID, ok := TenAddedTicketRole[role]; ok {
			// 現在のロールを削除します
			if err := s.GuildMemberRoleRemove(guildID, userID, role); err != nil {
				return errors.NewError("現在の回数ロールを削除できません", err)
			}
			// 新規ロールを付与します
			if err := s.GuildMemberRoleAdd(guildID, userID, afterRoleID); err != nil {
				return errors.NewError("新規の回数ロールを付与できません", err)
			}

			return nil
		}
	}

	// 回数ロールを保持していない（使い切った）場合は、+10を付与する
	if err := s.GuildMemberRoleAdd(guildID, userID, id.RoleID().SLOT_10_TICKET); err != nil {
		return errors.NewError("新規の回数ロールを付与できません", err)
	}

	return nil
}

// チケットロールかどうかを判定します
//
// Addedロールは含みません。
func IsSlotTicketRole(roleID string) bool {
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
		id.RoleID().SLOT_15_TICKET:

		return true
	}

	return false
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
