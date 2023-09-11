package main

import (
	"fmt"
	"log"

	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/techstart35/kifuneso-bot/api"
	"github.com/techstart35/kifuneso-bot/handler"
)

const (
	location = "Asia/Tokyo"
)

func init() {
	loc, err := time.LoadLocation(location)
	if err != nil {
		loc = time.FixedZone(location, 9*60*60)
	}
	time.Local = loc

	if err = godotenv.Load(".env"); err != nil {
		panic(fmt.Sprintf(".envを読み込めません: %v", err))
	}
}

func main() {
	var Token = "Bot " + os.Getenv("APP_BOT_TOKEN")

	session, err := discordgo.New(Token)
	session.Token = Token
	if err != nil {
		panic(err)
	}

	//イベントハンドラを追加
	handler.Handler(session)

	if err = session.Open(); err != nil {
		panic(err)
	}

	// 直近の関数（main）の最後に実行される
	defer func() {
		if err = session.Close(); err != nil {
			panic(err)
		}
		return
	}()

	// Gin
	{
		engine := gin.Default()

		// CORSの設定
		// ここからCorsの設定
		engine.Use(cors.New(cors.Config{
			// アクセスを許可したいアクセス元
			AllowOrigins: []string{
				"*",
			},
			// アクセスを許可したいHTTPメソッド(以下の例だとPUTやDELETEはアクセスできません)
			AllowMethods: []string{
				"POST",
				"GET",
			},
			// 許可したいHTTPリクエストヘッダ
			AllowHeaders: []string{
				"Access-Control-Allow-Credentials",
				"Access-Control-Allow-Headers",
				"Content-Type",
				"Content-Length",
				"Accept-Encoding",
				"Authorization",
				"Token",
			},
			// cookieなどの情報を必要とするかどうか
			//AllowCredentials: true,
			// preflightリクエストの結果をキャッシュする時間
			//MaxAge: 24 * time.Hour,
		}))

		// ルートを設定する
		api.RegisterRouter(engine)

		if err = engine.Run(":8080"); err != nil {
			log.Fatal("起動に失敗しました")
		}
	}

	stopBot := make(chan os.Signal, 1)
	signal.Notify(stopBot, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-stopBot
	return
}
