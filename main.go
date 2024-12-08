package main

import (
	"github.com/joho/godotenv"
	"github.com/kajiLabTeam/walking-information-storage-server-api/app/infrastructure"
)

func main() {
	// 環境変数の読み込み
	err := godotenv.Load("../.env")
	if err != nil {
		panic(err)
	}

	_, err = infrastructure.Connect()
	if err != nil {
		panic(err)
	}

	// r := gin.Default()

	// // 以下にコントローラーを追加
	// controllers.HealthCheckHandler(r)

	// r.Run(":8080")
}
