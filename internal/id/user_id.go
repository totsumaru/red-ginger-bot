package id

import "os"

type User struct {
	BOT       string
	TOTSUMARU string
}

func UserID() User {
	if os.Getenv("ENV") == "dev" {
		return User{
			BOT:       "1081377392834134156",
			TOTSUMARU: "960104306151948328",
		}
	} else {
		return User{
			BOT:       "1120187983639891978",
			TOTSUMARU: "960104306151948328",
		}
	}
}
