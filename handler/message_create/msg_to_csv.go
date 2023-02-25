package message_create

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/gocarina/gocsv"
	"io"
	"log"
	"os"
	"time"
)

const (
	CSVCommand           = "!rg.csv"
	MessageReadChannelID = "1078520828108476456"
)

// メッセージを重複なしで読み込んで、csvで出力します
func MsgToCSV(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Content != CSVCommand {
		return
	}

	// 上限の数を指定
	const limit = 100

	beforeID := ""
	messages := make([]*discordgo.Message, 0)

	// 全てのメッセージを取得します
	for {
		msgs, err := s.ChannelMessages(
			MessageReadChannelID, // channelID
			limit,                // limit
			beforeID,             // beforeID
			"",                   // afterID
			"",                   // aroundID
		)
		if err != nil {
			log.Fatal(err)
		}

		messages = append(messages, msgs...)

		// limitで指定した件数と一致しない（それ以下）は終了
		if len(msgs) != limit {
			fmt.Println("Len: ", len(msgs))
			break
		}

		// 上限まで取得した場合は未取得のものがある可能性が残っているため、
		// 取得した最後のメッセージIDをbeforeIDを設定
		beforeID = msgs[len(msgs)-1].ID
	}

	// csvのリクエストを作成します
	req := make([]OutputCSVReq, 0)
	for _, msg := range messages {
		r := OutputCSVReq{
			Timestamp: msg.Timestamp,
			UserName:  msg.Author.Username,
			Content:   msg.Content,
		}

		req = append(req, r)
	}

	// コマンドが実行されたチャンネルにcsvを送信します
	outputCSV(s, m.ChannelID, req)
}

// csvに出力するリクエストです
type OutputCSVReq struct {
	Timestamp time.Time
	UserName  string
	Content   string
}

// csv出力します
func outputCSV(s *discordgo.Session, channelID string, req []OutputCSVReq) io.Reader {
	file, _ := os.OpenFile("test.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	defer file.Close()

	if err := gocsv.MarshalFile(req, file); err != nil {
		log.Fatal(err)
	}

	f, err := os.Open("test.csv")
	if err != nil {
		log.Fatal(err)
	}

	// ファイルを送信します
	_, err = s.ChannelFileSend(channelID, "output.csv", f)
	if err != nil {
		log.Fatal(err)
	}

	return file
}
