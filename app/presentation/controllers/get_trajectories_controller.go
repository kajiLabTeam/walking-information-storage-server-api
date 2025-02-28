package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// 構造体の変数名は、パスカルケース(大文字)

//EstimationTrajectoriesId(推定軌跡)は配列
//EstimationPosition(推定座標)は２重配列

type GetTrajectoriesResponse struct {
	Floor        Floor        `json:"floor"`
	Trajectories []Trajectory `json:"trajectories"`
}

type Position struct {
	ID       int     `json:"id"`
	X        float32 `json:"x"`
	Y        float32 `json:"y"`
	WalkedAt int     `json:"walked_at"`
}

type Trajectory struct {
	ID        string     `json:"id"`
	Estimated []Position `json:"estimated"`
	Correct   []Position `json:"correct"`
}

type Floor struct {
	ID          string `json:"id"`
	MapImageURL string `json:"map_image_url"`
}

// http://localhost:8080/api/buildings/building_id/floors/floor_id/trajectories?is_correct_included=true&count=100
func GetTrajectoriesController(r *gin.Engine) {
	r.GET("/api/buildings/:building_id/floors/:floor_id/trajectories/", func(c *gin.Context) {

		buildingId := c.Param("building_id")
		floorId := c.Param("floor_id")

		isCorrectIncluded := c.Query("is_correct_included")
		count := c.Query("count")

		fmt.Println("----パスパラメータ----")
		fmt.Println(buildingId)
		fmt.Println(floorId)

		fmt.Println("----リクエストがきたよ----")
		fmt.Println(isCorrectIncluded)
		fmt.Println(count)
		res := GetTrajectoriesResponse{
			Floor: Floor{
				ID:          "1",
				MapImageURL: "署名付きURL",
			},
			Trajectories: []Trajectory{
				{
					ID: "1",
					Estimated: []Position{
						{ID: 1, X: 36.1834166, Y: 138.110467, WalkedAt: 1560000000},
						{ID: 2, X: 35.1834166, Y: 137.110467, WalkedAt: 1560000000},
					},
					Correct: []Position{
						{ID: 1, X: 30.1834166, Y: 137.110467, WalkedAt: 1560000000},
						{ID: 2, X: 31.1834166, Y: 137.110467, WalkedAt: 1560000000},
					},
				},
				{
					ID: "2",
					Estimated: []Position{
						{ID: 1, X: 36.1834166, Y: 138.110467, WalkedAt: 1560000000},
						{ID: 2, X: 35.1834166, Y: 137.110467, WalkedAt: 1560000000},
					},
					Correct: []Position{
						{ID: 1, X: 30.1834166, Y: 137.110467, WalkedAt: 1560000000},
						{ID: 2, X: 31.1834166, Y: 137.110467, WalkedAt: 1560000000},
					},
				},
			},
		}
		c.JSON(200, res)

	})
}
