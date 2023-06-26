package id

import "os"

type Channel struct {
	ERR_LOG        string
	SLOT_PRIZE_LOG string
	TEST           string
}

func ChannelID() Channel {
	if os.Getenv("ENV") == "dev" {
		return Channel{
			ERR_LOG:        "1075023820395663440",
			SLOT_PRIZE_LOG: "1122852681967489125",
			TEST:           "1073275782287343637",
		}
	} else {
		return Channel{
			ERR_LOG:        "1075023820395663440",
			SLOT_PRIZE_LOG: "1122852681967489125",
			TEST:           "1073275782287343637",
		}
	}
}
