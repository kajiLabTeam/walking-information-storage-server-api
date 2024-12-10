package controllers

import (
	"github.com/gin-gonic/gin"
)

// 構造体の変数名は、パスカルケース(大文字)

// EstimationTrajectoriesId(軌跡)は配列
//EstimationPosition(推定座標)は２重配列

type GetTrajectoriesResponse struct {
	FloorMapImage            string `json:"floor_map_image"`
	EstimationTrajectoriesId int    `json:"estimation_trajectories_id"`
	EstimationPosition       int    `json:"id"`
}

func GetTrajectoriesController(r *gin.Engine) {
	r.GET("/api", func(c *gin.Context) {
		res := GetTrajectoriesResponse{
			FloorMapImage:            "gsgsgfs12144",
			EstimationTrajectoriesId: 1,
			EstimationPosition:       1,
		}

		c.JSON(200, res)
	})
}
