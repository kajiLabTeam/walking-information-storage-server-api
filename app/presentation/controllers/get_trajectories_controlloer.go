package controllers

import (
	"github.com/gin-gonic/gin"
)

// 構造体の変数名は、パスカルケース(大文字)

//EstimationTrajectoriesId(推定軌跡)は配列
//EstimationPosition(推定座標)は２重配列

type GetTrajectoriesResponse struct {
	FloorMapImage          string                 `json:"floor_map_image"`
	EstimationTrajectories EstimationTrajectories `json:"estimation_trajectories"`
}
type EstimationTrajectories struct {
	EstimationTrajectoriesID int `json:"estimation_trajectories_id"`
	EstimationPosition
}
type EstimationPosition struct {
	ID int     `json:"id"`
	X  float32 `json:"x"`
	Y  float32 `json:"y"`
}

func GetTrajectoriesController(r *gin.Engine) {
	r.GET("/api", func(c *gin.Context) {
		estimationTrajectories := EstimationTrajectories{
			EstimationTrajectoriesID: 123,
			EstimationPosition: EstimationPosition{
				ID: 1, X: 32.4455, Y: 153.34555,
			},
		}
		res := GetTrajectoriesResponse{
			FloorMapImage:          "gsgsgfs12144",
			EstimationTrajectories: estimationTrajectories,
		}
		c.JSON(200, res)

	})
}
