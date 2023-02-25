package src

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

// 埋め込みメッセージを送信します
//
// `!r-send <Title> <fromMessageURL> <ToChannelMention>`
func SendEmbedMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	input := strings.Split(m.Content, " ")

	if len(input) != 4 {
		return
	}

	tmpToChID := input[3]
	str1 := strings.Replace(tmpToChID, "<#", "", -1)
	str2 := strings.Replace(str1, ">", "", -1)

	cmd := input[0]
	title := input[1]
	fromChannelID := strings.Split(input[2], "/")[5]
	fromMessageID := strings.Split(input[2], "/")[6]
	toChannelID := str2

	if cmd != CMDSend {
		return
	}

	msg, err := s.ChannelMessage(fromChannelID, fromMessageID)
	if err != nil {
		panic(err)
	}

	imageURL := ""
	if len(msg.Attachments) > 0 {
		imageURL = msg.Attachments[0].URL
	}

	req := &discordgo.MessageEmbed{
		Title: title,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: imageURL,
		},
		Description: msg.Content,
		Color:       Color,
	}
	_, err = s.ChannelMessageSendEmbed(toChannelID, req)
	if err != nil {
		panic(err)
	}
}
