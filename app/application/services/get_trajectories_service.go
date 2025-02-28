package services

import (
	"fmt"
	"log"

	"github.com/kajiLabTeam/walking-information-storage-server-api/app/infrastructure"
	"github.com/kajiLabTeam/walking-information-storage-server-api/app/infrastructure/repository"
	"github.com/kajiLabTeam/walking-information-storage-server-api/app/presentation/controllers"
)

func GetTrajectoriesService(trajectoryID string) (*controllers.GetTrajectoriesResponse, error) {
	db, err := infrastructure.ConnectionDB()
	if err != nil {
		log.Fatal("データベース接続エラー:", err)
	}
	defer db.Close() // 呼び出し側で DB を閉じる

	//floor_id(フロアID)を元に軌跡ID(trajectry_id)を取得
	trajectory, err := repository.GetTrajectoryIDByFloorID(db, "01F8VYXK67BGC1F9RP1E4S9YTV")
	if err != nil {
		log.Printf("軌跡IDの取得エラー: %w", err)
	}
	fmt.Println("trajectory")
	fmt.Println(trajectory)
	//軌跡ID(trajectry_id)に紐付いた推定座標(estimated_positions)/正解座標(correct_positions)を取得
	estimatedPosition, err := repository.GetEstimatedPositionsByTrajectoryID(db, trajectory.ID)
	if err != nil {
		log.Printf("推定軌跡の取得エラー: %w", err)
	}

	correctPositions, err := repository.GetCorrectPositionsByTrajectoryID(db, trajectory.ID)
	if err != nil {
		log.Printf("正解軌跡の取得エラー: %w", err)
	}

	//表示
	fmt.Println("estimatedPositions")
	fmt.Printf("%+v\n", estimatedPosition)
	fmt.Println("correctPositions")
	fmt.Printf("%+v\n", correctPositions)

	//推定座標(estimatedPositions)を配列にして、推定軌跡(estimatedTrajectories)に変形
	positions := []controllers.Position{}

	//正解座標(CorrectPositions)を配列にして、正解軌跡(CorrectTrajectories)にする

	if err != nil {
		log.Printf("データベースのクローズに失敗: %v", err)
	}

	//ここどうしよ
	// application.EstimatedPositionEncoding(estimatedPosition)

}
