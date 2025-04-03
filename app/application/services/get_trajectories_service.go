package services

import (
	"fmt"
	"log"

	"github.com/kajiLabTeam/walking-information-storage-server-api/app/infrastructure"
	"github.com/kajiLabTeam/walking-information-storage-server-api/app/infrastructure/external/services"
	"github.com/kajiLabTeam/walking-information-storage-server-api/app/infrastructure/repository"
	dto_presentation "github.com/kajiLabTeam/walking-information-storage-server-api/app/presentation/dto"
)

func GetTrajectoriesService(floorID string) (*dto_presentation.GetTrajectoriesResponse, error) {
	db, err := infrastructure.ConnectionDB()
	if err != nil {
		log.Fatal("データベース接続エラー:", err)
	}
	defer db.Close() // 呼び出し側で DB を閉じる

	session, err := infrastructure.ConnectionMinio()
	if err != nil {
		log.Fatal("Minio接続エラー:", err)
	}

	//floor_id(フロアID)を元に軌跡ID(trajectry_id)を取得
	trajectories, err := repository.GetTrajectoriesIDByFloorID(db, floorID)
	if err != nil {
		log.Printf("軌跡IDの取得エラー: %w", err)
	}
	fmt.Println("trajectories")
	fmt.Println(trajectories)

	resTrajectories := []dto_presentation.Trajectory{}

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
		estimated := []dto_presentation.Position{}
		correct := []dto_presentation.Position{}

		//推定座標estimatedPositionをPosition型に変更する
		for _, estimatedPosition := range estimatedPositions {
			position := dto_presentation.Position{
				ID:       estimatedPosition.ID,
				X:        float32(estimatedPosition.X),
				Y:        float32(estimatedPosition.Y),
				WalkedAt: estimatedPosition.CreatedAt,
			}

			estimated = append(estimated, position)

		}
		//推定座標estimatedPositionをPosition型に変更する
		for _, correctPosition := range correctPositions {
			position := dto_presentation.Position{
				ID:       correctPosition.ID,
				X:        float32(correctPosition.X),
				Y:        float32(correctPosition.Y),
				WalkedAt: correctPosition.CreatedAt,
			}

			correct = append(correct, position)

		}

		resTrajectory := dto_presentation.Trajectory{
			ID:        trajectory.ID,
			Estimated: estimated,
			Correct:   correct,
		}

		resTrajectories = append(resTrajectories, resTrajectory)

	}

	//Todo: repositoryのfloor_information関数からフロア情報IDを取得
	floorInformationRecord, err := repository.GetFloorInformationIDByfloorID(db, floorID)
	if err != nil {
		log.Printf("フロア情報IDの取得エラー: %w", err)
	}
	fmt.Println("GetFloorInformationIDByfloorID", floorInformationRecord.ID)

	// Todo : フロア情報IDからフロアマップ画像を取得

	floorMap, err := services.Download(session, floorID, floorInformationRecord.ID)
	if err != nil {
		log.Printf("フロアマップ署名付きURLの取得エラー: %w", err)
	}
	resFloor := dto_presentation.Floor{
		ID:          floorID,
		MapImageURL: floorMap,
	}

	trajectoriesResponse := dto_presentation.GetTrajectoriesResponse{
		Floor:        resFloor,
		Trajectories: resTrajectories,
	}

	return &trajectoriesResponse, nil
}
