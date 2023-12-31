package id

import "os"

type Channel struct {
	ERR_LOG        string
	SLOT_PRIZE_LOG string
	TEST           string
	JP_CHAT        string
	STORY          string
	CHARACTER      string
	TEAM_MEMBER    string
	PROJECT_INFO   string
	SOZAI          string
	OFFICIAL_LINKS string
	MOD            string
	RANKING        string
}

func ChannelID() Channel {
	if os.Getenv("ENV") == "dev" {
		return Channel{
			ERR_LOG:        "1128139181957333002",
			TEST:           "1128139169743507599",
			STORY:          "1128152975743930469",
			CHARACTER:      "1128152999647268894",
			TEAM_MEMBER:    "1128153022904680580",
			PROJECT_INFO:   "1128508560393895946",
			SOZAI:          "1128178255929823264",
			OFFICIAL_LINKS: "1128508486620282920",
			MOD:            "1129618901261631648",
			RANKING:        "",
		}
	} else {
		return Channel{
			ERR_LOG:        "1075023820395663440",
			SLOT_PRIZE_LOG: "1122852681967489125",
			TEST:           "1073275782287343637",
			JP_CHAT:        "1067227422073815050",
			STORY:          "1067226670651674735",
			CHARACTER:      "1067226746824446014",
			TEAM_MEMBER:    "1077079036657946695",
			PROJECT_INFO:   "1110418665368735764",
			SOZAI:          "1116640396651548773",
			OFFICIAL_LINKS: "1070289892568997898",
			MOD:            "1067226942446772335",
			RANKING:        "1150697516162097182",
		}
	}
}
