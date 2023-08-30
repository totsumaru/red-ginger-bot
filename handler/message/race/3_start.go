package race

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/kifuneso-bot/internal/color"
	"github.com/techstart35/kifuneso-bot/internal/errors"
)

// 開始メッセージを送信します
//
// Emoji: []UserID を返します。
func sendStart(
	s *discordgo.Session,
	channelID, entryMsgID string,
) (map[string][]*discordgo.User, error) {
	res := map[string][]*discordgo.User{}

	var description = `
-- 受付終了 --
レースを開始します！

エントリー: %d人
`
	// リアクションした人を取得します
	users, err := getReactedUsers(s, channelID, entryMsgID)
	if err != nil {
		return res, errors.NewError("リアクションしたユーザーを取得できません", err)
	}
	res = users

	userCount := 0
	for _, user := range res {
		userCount += len(user)
	}

	embed := &discordgo.MessageEmbed{
		Description: fmt.Sprintf(description, userCount),
		Color:       color.Green,
	}

	_, err = s.ChannelMessageSendEmbed(channelID, embed)
	if err != nil {
		return res, errors.NewError("メッセージを送信できません", err)
	}

	return res, nil
}

// リアクションした人を取得します
func getReactedUsers(
	s *discordgo.Session,
	channelID, msgID string,
) (map[string][]*discordgo.User, error) {
	res := map[string][]*discordgo.User{}

	// 重複確認用です
	allUsers := map[string]bool{}

	// 絵文字1つずつ、処理をします
	for _, entryUser := range EntryUsers {
		users := make([]*discordgo.User, 0)

		// 最大1000人まで参加可能（10 * 100）
		for i := 0; i < 10; i++ {
			var afterID string

			switch i {
			case 0:
				afterID = ""
			default:
				if len(users) == 0 {
					break
				}
				afterID = users[len(users)-1].ID // panic注意
			}

			us, err := s.MessageReactions(channelID, msgID, entryUser.Emoji, 100, "", afterID)
			if err != nil {
				return res, errors.NewError("リアクションを取得できません", err)
			}

			if len(us) == 0 ||
				len(us) == 1 && us[0].ID == s.State.User.ID {
				break
			}

			// bot自身は除外し、重複は排除します
			for _, u := range us {
				// botは無効
				if u.ID == s.State.User.ID {
					continue
				}
				// すでにエントリーがある場合はpass
				if ok := allUsers[u.ID]; ok {
					continue
				}
				users = append(users, u)
				allUsers[u.ID] = true
			}
		}
		res[entryUser.Emoji] = users
	}

	return res, nil
}
