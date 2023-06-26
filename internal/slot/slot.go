package slot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/kifuneso-bot/internal/errors"
	"github.com/techstart35/kifuneso-bot/internal/id"
)

// 今のロールを削除し、+5したロールを付け直します
func UpdateRoleToPlus5(s *discordgo.Session, guildID, userID string, currentRoles []string) error {
	currentCount, currentRoles := getRemainingSlotTicketCount(currentRoles)

	if err := addPlus5SlotTicketRole(s, guildID, userID, currentCount, currentRoles); err != nil {
		return errors.NewError("チケットロールを更新できません", err)
	}

	return nil
}

// 現在のスロットチケットの残り回数と、現在のスロットチケットロールを取得します
func getRemainingSlotTicketCount(roles []string) (int, []string) {
	currentCount := 0 // 現在のスロットチケットの残り回数
	currentRoles := make([]string, 0)

	for _, role := range roles {
		switch role {
		case id.RoleID().SLOT_1_TICKET:
			currentCount += 1
			currentRoles = append(currentRoles, role)
		case id.RoleID().SLOT_2_TICKET:
			currentCount += 2
			currentRoles = append(currentRoles, role)
		case id.RoleID().SLOT_3_TICKET:
			currentCount += 3
			currentRoles = append(currentRoles, role)
		case id.RoleID().SLOT_4_TICKET:
			currentCount += 4
			currentRoles = append(currentRoles, role)
		case id.RoleID().SLOT_5_TICKET:
			currentCount += 5
			currentRoles = append(currentRoles, role)
		case id.RoleID().SLOT_10_TICKET:
			currentCount += 10
			currentRoles = append(currentRoles, role)
		case id.RoleID().SLOT_15_TICKET:
			currentCount += 15
			currentRoles = append(currentRoles, role)
		}
	}

	return currentCount, currentRoles
}

// 現在の残回数に5回追加したロールを付与します
//
// 既存の回数ロールは削除します。
func addPlus5SlotTicketRole(
	s *discordgo.Session,
	guildID string,
	userID string,
	currentCount int,
	currentRoles []string,
) error {
	// 既存の回数ロールは削除します
	for _, currentRole := range currentRoles {
		if err := s.GuildMemberRoleRemove(guildID, userID, currentRole); err != nil {
			return errors.NewError("ロールを削除できません", err)
		}
	}

	// 新規追加ロールを算出します
	addRoleIDs := make([]string, 0)

	switch currentCount {
	case 0:
		addRoleIDs = []string{id.RoleID().SLOT_5_TICKET}
	case 1:
		addRoleIDs = []string{id.RoleID().SLOT_5_TICKET, id.RoleID().SLOT_1_TICKET}
	case 2:
		addRoleIDs = []string{id.RoleID().SLOT_5_TICKET, id.RoleID().SLOT_2_TICKET}
	case 3:
		addRoleIDs = []string{id.RoleID().SLOT_5_TICKET, id.RoleID().SLOT_3_TICKET}
	case 4:
		addRoleIDs = []string{id.RoleID().SLOT_5_TICKET, id.RoleID().SLOT_4_TICKET}
	case 5:
		addRoleIDs = []string{id.RoleID().SLOT_10_TICKET}
	case 6:
		addRoleIDs = []string{id.RoleID().SLOT_10_TICKET, id.RoleID().SLOT_1_TICKET}
	case 7:
		addRoleIDs = []string{id.RoleID().SLOT_10_TICKET, id.RoleID().SLOT_2_TICKET}
	case 8:
		addRoleIDs = []string{id.RoleID().SLOT_10_TICKET, id.RoleID().SLOT_3_TICKET}
	case 9:
		addRoleIDs = []string{id.RoleID().SLOT_10_TICKET, id.RoleID().SLOT_4_TICKET}
	case 10:
		addRoleIDs = []string{id.RoleID().SLOT_15_TICKET}
	case 11, 12, 13, 14, 15:
		addRoleIDs = []string{id.RoleID().SLOT_15_TICKET}
	}

	// ロールを新規追加します
	for _, addRoleID := range addRoleIDs {
		if err := s.GuildMemberRoleAdd(guildID, userID, addRoleID); err != nil {
			return errors.NewError("ロールを付与できません", err)
		}
	}

	return nil
}
