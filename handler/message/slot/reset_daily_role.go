package slot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/kifuneso-bot/internal/errors"
	"github.com/techstart35/kifuneso-bot/internal/id"
	"github.com/techstart35/kifuneso-bot/internal/slot"
)

// 全てのSlot回数券をリセットし、5回券を全員に付与します
//
// - #test-roomでしかコマンドを実行できません
// - Verifyロールを保持している人だけに付与します
func ResetAndAddDailyRole(s *discordgo.Session, m *discordgo.MessageCreate) error {
	if m.ChannelID != id.ChannelID().TEST {
		return nil
	}

	// 進捗メッセージを送信
	if _, err := s.ChannelMessageSend(m.ChannelID, "Process is running..."); err != nil {
		return errors.NewError("進捗メッセージを送信できません", err)
	}

	users, err := getUsersWithRole(s, m.GuildID, id.RoleID().VERIFY)
	if err != nil {
		return errors.NewError("特定のロールを持つユーザーを取得できません", err)
	}

	// この中でエラーが発生すると、その後のユーザー全てに影響が出るため、
	// エラーは通知のみでreturnはしません。
	for i, user := range users {
		// ユーザーからSlotの回数券ロールとAddedロールを削除します
		for _, role := range user.Roles {
			if slot.IsSlotTicketRoleContainsAddedRole(role) {
				if err = s.GuildMemberRoleRemove(m.GuildID, user.User.ID, role); err != nil {
					errors.SendErrMsg(s, errors.NewError("ロールを削除できません。通知のみ行います"), user.User)
				}
			}
		}

		// ユーザーに5回券ロールを付与します
		if err = s.GuildMemberRoleAdd(m.GuildID, user.User.ID, id.RoleID().SLOT_5_TICKET); err != nil {
			errors.SendErrMsg(s, errors.NewError("5回券ロールを付与できません。通知のみ行います"), user.User)
		}

		if i%50 == 0 {
			// 途中経過メッセージを送信
			if _, err = s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%d件の処理が完了しています", i)); err != nil {
				return errors.NewError("完了メッセージを送信できません", err)
			}
		}
	}

	// 完了メッセージを送信
	if _, err = s.ChannelMessageSend(m.ChannelID, "Finished!"); err != nil {
		return errors.NewError("完了メッセージを送信できません", err)
	}

	return nil
}

// 特定のロールを持つユーザーを全て取得します
func getUsersWithRole(s *discordgo.Session, guildID, roleID string) ([]*discordgo.Member, error) {
	members := make([]*discordgo.Member, 0)

	var lastID string

	for {
		guildMembers, err := s.GuildMembers(guildID, lastID, 1000)
		if err != nil {
			return nil, err
		}

		for _, member := range guildMembers {
			for _, memberRoleID := range member.Roles {
				if memberRoleID == roleID {
					members = append(members, member)
					break
				}
			}
		}

		if len(guildMembers) < 1000 {
			break
		}

		lastID = guildMembers[len(guildMembers)-1].User.ID
	}

	return members, nil
}
