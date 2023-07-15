package id

import (
	"fmt"
	"os"
)

type Message struct {
	CHARACTER_GG string
}

func MessageID() Message {
	if os.Getenv("ENV") == "dev" {
		return Message{
			CHARACTER_GG: "1129229486420938792",
		}
	} else {
		return Message{}
	}
}

// メッセージリンクを作成します
func GenerateMessageLink(serverID, channelID, messageID string) string {
	const baseURL = "https://discord.com/channels/%s/%s/%s"
	return fmt.Sprintf(
		baseURL,
		serverID,
		channelID,
		messageID,
	)
}
