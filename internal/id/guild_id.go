package id

import "os"

func GuildID() string {
	if os.Getenv("ENV") == "dev" {
		return "1128123009165697126"
	} else {
		return "1047763497594933319"
	}
}
