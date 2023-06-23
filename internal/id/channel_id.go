package id

import "os"

type Channel struct {
	ERR_LOG           string
	SELF_INTRO        string
	LAB_EN            string
	LAB_JP            string
	SCORE             string
	SCORE_LOG         string
	ANIME             string
	GAME              string
	VTUBER            string
	FORUM_THREAD_ID_1 string
	FORUM_THREAD_ID_2 string
	FORUM_THREAD_ID_3 string
	FORUM_THREAD_ID_4 string
}

func ChannelID() Channel {
	if os.Getenv("ENV") == "dev" {
		return Channel{
			ERR_LOG:           "1088362090793795664",
			SELF_INTRO:        "1101876439873245285",
			LAB_EN:            "1118397928948637787",
			LAB_JP:            "1118397946585682010",
			SCORE:             "1107004046293880994",
			SCORE_LOG:         "1114118994031562853",
			ANIME:             "1115447803062272001",
			GAME:              "1115447819793334323",
			VTUBER:            "1115447872159219813",
			FORUM_THREAD_ID_1: "1115818296051257494", // Share you creation
			FORUM_THREAD_ID_2: "1115818160780754974", // What artists would you like us to pay attention to?
			FORUM_THREAD_ID_3: "1115817938721714196", // Manga Otaku Gathering
			FORUM_THREAD_ID_4: "1115541095166779403", // Your favorite character
		}
	} else {
		return Channel{
			ERR_LOG:           "1002229253892472852",
			SELF_INTRO:        "1107553532766326794",
			LAB_EN:            "1078087982092927056",
			LAB_JP:            "1093831890844270673",
			SCORE:             "1107594248880148503",
			SCORE_LOG:         "1115159179380932658",
			ANIME:             "1079923148268523672",
			GAME:              "1106958955634237580",
			VTUBER:            "1106959304180912250",
			FORUM_THREAD_ID_1: "1115793426122813480", // Share you creation
			FORUM_THREAD_ID_2: "1107681573945487462", // What artists would you like us to pay attention to?
			FORUM_THREAD_ID_3: "1116532489050136616", // Manga Otaku Gathering
			FORUM_THREAD_ID_4: "1107516189468536963", // Your favorite character
		}
	}
}
