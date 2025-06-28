// @title gin-swagger sample
// @version 1.0
// @license.name yyamakura
package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"clean-architecture-sample/internal/config"
	"clean-architecture-sample/internal/infrastructure"
	"clean-architecture-sample/internal/interface/handler"
	"clean-architecture-sample/internal/usecase"

	// Swagger UI
	_ "clean-architecture-sample/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	cfg := config.LoadConfig() // 設定を読み込む

	// インフラ層: ファイルベースのユーザーリポジトリ
	repo := infrastructure.NewFileUserRepository(cfg.DataFile)

	// ユースケース層: ユーザー用ユースケース
	userUsecase := usecase.NewUserUseCase(repo)

	// ハンドラー層: Gin用ハンドラー
	userHandler := handler.NewUserHandler(userUsecase)

	router := gin.Default()

	router.POST("/users", userHandler.CreateUserHandler)
	router.GET("/users/:id", userHandler.GetUserHandler)
	router.PUT("/users/:id", userHandler.UpdateUserHandler)
	router.DELETE("/users/:id", userHandler.DeleteUserHandler)

	// Swagger UI
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Printf("サーバーをポート8080で起動中...")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("サーバー起動エラー: %v", err)
	}
}
