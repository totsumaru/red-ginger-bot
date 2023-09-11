package api

import (
	"github.com/gin-gonic/gin"
)

// ルートを設定します
func RegisterRouter(e *gin.Engine) {
	Route(e)
}

// ルートです
//
// Note: この関数は削除しても問題ありません
func Route(e *gin.Engine) {
	e.GET("/", func(c *gin.Context) {
		c.Header("hello", "world")
		c.JSON(200, gin.H{
			"message": "hello",
		})
	})
}
