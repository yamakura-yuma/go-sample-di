// @title サンプルAPI
// @version 1.0
// @description 最小構成のSwaggerサンプル
// @host localhost:8080
// @BasePath /
package main

import (
	"github.com/gin-gonic/gin"

	_ "clean-architecture-sample/docs" // Swagger UIのためのインポート
)

// PingResponse サンプルレスポンス
// @Description ping応答
// @Success 200 {object} PingResponse
// @Router /ping [get]
type PingResponse struct {
	Message string `json:"message"`
}

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, PingResponse{Message: "pong"})
	})

	r.Run(":8080")
}
