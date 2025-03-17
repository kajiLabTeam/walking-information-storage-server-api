package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	//"github.com/kajiLabTeam/walking-information-storage-server-api/app/infrastructure"

	"github.com/kajiLabTeam/walking-information-storage-server-api/app/application/services"
	"github.com/kajiLabTeam/walking-information-storage-server-api/app/presentation/controllers"
)

func main() {
	// 環境変数の読み込み
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	services.GetTrajectoriesService("01F8VYXK67BGC1F9RP1E4S9YTV")

	// 以下にコントローラーを追加
	controllers.HealthCheckController(r)
	controllers.GetTrajectoriesController(r)

	r.Run(":8080")
}
