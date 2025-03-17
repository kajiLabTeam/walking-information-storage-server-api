package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kajiLabTeam/walking-information-storage-server-api/app/application/services"
)

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
		res, err := services.GetTrajectoriesService(floorId)
		if err != nil {
			return
		}

		c.JSON(200, res)
	})
}
