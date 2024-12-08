package controllers

import (
	"github.com/gin-gonic/gin"
)

type HealthCheckResponse struct {
	Status string `json:"status"`
}

func HealthCheckController(r *gin.Engine) {
	r.GET("/health", func(c *gin.Context) {
		res := HealthCheckResponse{
			Status: "OK",
		}

		c.JSON(200, res)
	})
}
