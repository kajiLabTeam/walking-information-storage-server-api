package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kajiLabTeam/walking-information-storage-server-api/app/presentation/controllers"
)

func main() {
	// 環境変数の読み込み
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	// 以下にコントローラーを追加
	controllers.HealthCheckController(r)
	controllers.GetTrajectoriesController(r)

	r.Run(":8080")
}
