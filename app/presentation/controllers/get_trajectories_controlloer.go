package controllers

import (
	"github.com/gin-gonic/gin"
)

type GetTrajectoriesResponse struct {
	Status string `json:"status"`
}

func GetTrajectoriesController(r *gin.Engine) {
	r.GET("/api", func(c *gin.Context) {
		res := GetTrajectoriesResponse{
			Status: "OK",
		}

		c.JSON(200, res)
	})
}
