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
		if err = sendCountDown(s, m.ChannelID, 5-i); err != nil {
			return errors.NewError("カウントダウンメッセージを送信できません", err)
		}
	}

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

	// ランキングを更新します
	{
		description := `
- ランキング -
(%s更新)

%s
`

		textLine := make([]string, 0)
		races, err := db.Race{}.FindAll()
		if err != nil {
			return errors.NewError("全ての情報を取得できません", err)
		}

		for _, race := range races {
			u, err := s.User(race.ID)
			if err != nil {
				return errors.NewError("ユーザーを取得できません", err)
			}

			line := fmt.Sprintf("%s: %dpt", u.Username, race.Point)
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
