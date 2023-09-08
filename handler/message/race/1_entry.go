package race

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/kifuneso-bot/internal/errors"
	"strings"
)

// エントリーメッセージを送信します
//
// メッセージIDを返します。
func sendEntryMessage(s *discordgo.Session, channelID string) (string, error) {
	lines := make([]string, 0)
	for _, user := range EntryUsers {
		line := fmt.Sprintf("%s｜%s", user.Emoji, user.Name)
		lines = append(lines, line)
	}

	var description = `
# RGグランプリ受付
次のAIロボ中から、1位を予想してください。

%s

※複数選択は無効です
※複数選んだ場合は、どれか1つでエントリーとなります
`

	embed := &discordgo.MessageEmbed{
		Description: fmt.Sprintf(description, strings.Join(lines, "\n")),
	}

	// エントリー画像を送信します
	_, err := s.ChannelMessageSendEmbed(channelID, &discordgo.MessageEmbed{
		Image: &discordgo.MessageEmbedImage{
			URL: "https://cdn.discordapp.com/attachments/1103240223376293938/1149546792036339762/RGGandPrix.png",
		},
	})

	msg, err := s.ChannelMessageSendEmbed(channelID, embed)
	if err != nil {
		return "", errors.NewError("メッセージを送信できません", err)
	}

	// リアクションを付与します
	for _, entry := range EntryUsers {
		if err = s.MessageReactionAdd(channelID, msg.ID, entry.Emoji); err != nil {
			return "", errors.NewError("絵文字を付与できません", err)
		}
	}

	return msg.ID, nil
}
