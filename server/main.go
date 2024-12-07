package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kajiLabTeam/walking-information-storage-server-api/server/presentation/controllers"
)

func main() {
	// 環境変数の読み込み
	err := godotenv.Load("../.env")
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	// 以下にコントローラーを追加していく
	controllers.HealthCheckHandler(r)

	r.Run(":8000")
}
