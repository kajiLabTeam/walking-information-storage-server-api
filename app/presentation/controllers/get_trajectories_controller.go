package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type GetTrajectoriesResponse struct {
	FloorMapImage          string                   `json:"floor_map_image"`
	EstimationTrajectories []EstimationTrajectories `json:"estimation_trajectories"`
	CorrectTrajectories    []CorrectTrajectories    `json:"correct_trajectories"`
}
type EstimationTrajectories struct {
	EstimationTrajectoriesID int                  `json:"estimation_trajectories_id"`
	EstimationPosition       []EstimationPosition `json:"estimation_position"`
}

type EstimationPosition struct {
	ID int     `json:"id"`
	X  float32 `json:"x"`
	Y  float32 `json:"y"`
}
type CorrectTrajectories struct {
	CorrectTrajectoriesID int               `json:"correct_trajectories_id"`
	CorrectPosition       []CorrectPosition `json:"correct_position"`
}
type CorrectPosition struct {
	ID int     `json:"id"`
	X  float32 `json:"x"`
	Y  float32 `json:"y"`
}

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

		estimationTrajectories := []EstimationTrajectories{
			{
				EstimationTrajectoriesID: 123,
				EstimationPosition: []EstimationPosition{
					{
						ID: 1, X: 32.4455, Y: 153.34555,
					},
					{

						ID: 1, X: 32.4455, Y: 153.34555,
					},
				},
			},
			{
				EstimationTrajectoriesID: 234,
				EstimationPosition: []EstimationPosition{
					{
						ID: 1, X: 32.4455, Y: 153.34555,
					},
					{

						ID: 1, X: 32.4455, Y: 153.34555,
					},
				},
			},
		}

		correctTrajectories := []CorrectTrajectories{
			{
				CorrectTrajectoriesID: 123,
				CorrectPosition: []CorrectPosition{
					{
						ID: 1, X: 32.4455, Y: 153.34555,
					},
					{

						ID: 1, X: 32.4455, Y: 153.34555,
					},
				},
			},
			{
				CorrectTrajectoriesID: 123,
				CorrectPosition: []CorrectPosition{
					{
						ID: 1, X: 32.4455, Y: 153.34555,
					},
					{

						ID: 1, X: 32.4455, Y: 153.34555,
					},
				},
			},
		}

		res := GetTrajectoriesResponse{

			FloorMapImage:          "gsgsgfs12144",
			EstimationTrajectories: estimationTrajectories,
			CorrectTrajectories:    correctTrajectories,
		}
		c.JSON(200, res)

	})
}
