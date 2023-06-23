package id

import "os"

type Role struct {
	VERIFY       string
	LAB_MEMBER   string
	QUEST1       string
	QUEST2       string
	QUEST3       string
	QUEST_HUNTER string
	SCORE        string
}

func RoleID() Role {
	if os.Getenv("ENV") == "dev" {
		return Role{
			VERIFY:       "1088361781254176868",
			LAB_MEMBER:   "1099236979041910804",
			QUEST1:       "1116627812556746752",
			QUEST2:       "1116627842738946048",
			QUEST3:       "1116627865098788945",
			QUEST_HUNTER: "1116676942033342566",
			SCORE:        "1118704171747655741",
		}
	} else {
		return Role{
			VERIFY:       "1080249448938098708",
			LAB_MEMBER:   "1081045642916417536",
			QUEST1:       "1120191594000289813",
			QUEST2:       "1120191700577570856",
			QUEST3:       "1120191752234610751",
			QUEST_HUNTER: "1120191808702525521",
			SCORE:        "1120173744099426334",
		}
	}
}
