package race

import (
	"fmt"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/kifuneso-bot/internal/db"
	"github.com/techstart35/kifuneso-bot/internal/errors"
	"github.com/techstart35/kifuneso-bot/internal/id"
)

func SendRace(s *discordgo.Session, m *discordgo.MessageCreate) error {
	// エントリーメッセージを送信
	msgID, err := sendEntryMessage(s, m.ChannelID)
	if err != nil {
		return errors.NewError("エントリーメッセージを送信できません", err)
	}

	// 4分,3分,2分,1分 のメッセージを送信
	for i := 1; i <= 4; i++ {
		time.Sleep(1 * time.Minute)
		if err = sendCountDown(s, m.GuildID, m.ChannelID, msgID, 5-i); err != nil {
			return errors.NewError("カウントダウンメッセージを送信できません", err)
		}
	}

	// 1分前-開始メッセージまでは本来1分ですが、
	// リアクション集計に時間がかかるため、10秒短くしています
	time.Sleep(50 * time.Second)

	// 開始メッセージを送信
	allUsers, err := sendStart(s, m.ChannelID, msgID)
	if err != nil {
		return errors.NewError("開始メッセージを送信できません", err)
	}

	time.Sleep(5 * time.Second)

	// 実況メッセージを送信
	entryUsers, err := sendProgress(s, m.ChannelID)
	if err != nil {
		return errors.NewError("実況メッセージを送信できません", err)
	}

	firstWinners := allUsers[entryUsers[0].Emoji]
	secondWinners := allUsers[entryUsers[1].Emoji]

	if err = sendWinnerUsers(s, m.ChannelID, firstWinners, secondWinners); err != nil {
		return errors.NewError("結果メッセージを送信できません", err)
	}

	// DBに格納します
	for _, winner := range firstWinners {
		r := db.Race{
			ID:    winner.ID,
			Point: 2, // 1位的中は2ポイント
		}

		if err = r.Upsert(); err != nil {
			return errors.NewError("Upsertに失敗しました", err)
		}
	}

	for _, winner := range secondWinners {
		r := db.Race{
			ID:    winner.ID,
			Point: 1, // 2位的中は1ポイント
		}

		if err = r.Upsert(); err != nil {
			return errors.NewError("Upsertに失敗しました", err)
		}
	}

	// ランキングを更新します
	{
		description := `
👑 ランキング 👑
（%s更新）

※ポイントが同じ場合は、先頭の数字に差があっても同じ順位とカウントします

%s
-------
`

		races, err := db.Race{}.FindAll()
		if err != nil {
			return errors.NewError("全ての情報を取得できません", err)
		}

		textLine := make([]string, 0)

		for index, race := range races {
			member, err := s.GuildMember(m.GuildID, race.ID)
			if err != nil {
				return errors.NewError("ユーザーを取得できません", err)
			}

			name := member.Nick
			if name == "" {
				name = member.User.Username
			}

			line := fmt.Sprintf("%d. %s: %dpt", index+1, name, race.Point)
			textLine = append(textLine, line)
		}

		today := time.Now().Format("2006-01-02")

		_, err = s.ChannelMessageSend(
			id.ChannelID().RANKING,
			fmt.Sprintf(description, today, strings.Join(textLine, "\n")),
		)
		if err != nil {
			return errors.NewError("メッセージを送信できません", err)
		}
	}

	return nil
}
