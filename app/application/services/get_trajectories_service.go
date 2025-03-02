package services

import (
	"fmt"
	"log"

	"github.com/kajiLabTeam/walking-information-storage-server-api/app/infrastructure"
	"github.com/kajiLabTeam/walking-information-storage-server-api/app/infrastructure/repository"
	dto_presenttion "github.com/kajiLabTeam/walking-information-storage-server-api/app/presentation/dto"
)

func GetTrajectoriesService(floorID string) (*dto_presenttion.GetTrajectoriesResponse, error) {
	db, err := infrastructure.ConnectionDB()
	if err != nil {
		log.Fatal("データベース接続エラー:", err)
	}
	defer db.Close() // 呼び出し側で DB を閉じる

	//floor_id(フロアID)を元に軌跡ID(trajectry_id)を取得
	//Todo:floorIDに置き換える
	trajectories, err := repository.GetTrajectoriesIDByFloorID(db, floorID)
	if err != nil {
		log.Printf("軌跡IDの取得エラー: %w", err)
	}
	fmt.Println("trajectories")
	fmt.Println(trajectories)

	resTrajectories := []dto_presenttion.Trajectory{}

	for _, trajectory := range trajectories {
		//軌跡ID(trajectry_id)に紐付いた推定座標(estimated_positions)/正解座標(correct_positions)を取得
		estimatedPositions, err := repository.GetEstimatedPositionsByTrajectoryID(db, trajectory.ID)
		if err != nil {
			log.Printf("推定軌跡の取得エラー: %w", err)
		}

		correctPositions, err := repository.GetCorrectPositionsByTrajectoryID(db, trajectory.ID)
		if err != nil {
			log.Printf("正解軌跡の取得エラー: %w", err)
		}
		// 推定座標(estimatedPositions)を配列にして、推定軌跡(estimatedTrajectories)に変形
		estimated := []dto_presenttion.Position{}
		correct := []dto_presenttion.Position{}

		//推定座標estimatedPositionをPosition型に変更する
		for _, estimatedPosition := range estimatedPositions {
			position := dto_presenttion.Position{
				ID:       estimatedPosition.ID,
				X:        float32(estimatedPosition.X),
				Y:        float32(estimatedPosition.Y),
				WalkedAt: estimatedPosition.CreatedAt,
			}

			estimated = append(estimated, position)
			// fmt.Printf("estimated: %v\n", estimated)
		}
		//推定座標estimatedPositionをPosition型に変更する
		for _, correctPosition := range correctPositions {
			position := dto_presenttion.Position{
				ID:       correctPosition.ID,
				X:        float32(correctPosition.X),
				Y:        float32(correctPosition.Y),
				WalkedAt: correctPosition.CreatedAt,
			}

			correct = append(correct, position)
			// fmt.Printf("correct: %v\n", correct)
		}

		resTrajectory := dto_presenttion.Trajectory{
			ID:        trajectory.ID,
			Estimated: estimated,
			Correct:   correct,
		}

		resTrajectories = append(resTrajectories, resTrajectory)
		// fmt.Printf("resTrajectories: %v\n", resTrajectories)
	}

	//正解座標(CorrectPositions)を配列にして、正解軌跡(CorrectTrajectories)にする

	if err != nil {
		log.Printf("データベースのクローズに失敗: %v", err)
	}

	floor := dto_presenttion.Floor{
		ID:          "1",
		MapImageURL: "署名付きURL",
	}

	trajectoriesResponse := dto_presenttion.GetTrajectoriesResponse{
		Floor:        floor,
		Trajectories: resTrajectories,
	}

	return &trajectoriesResponse, nil
}
