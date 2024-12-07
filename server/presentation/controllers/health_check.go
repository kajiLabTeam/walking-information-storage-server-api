package controllers

import (
	"github.com/gin-gonic/gin"
)

type HealthCheckResponse struct {
	Status string `json:"status"`
}

func HealthCheckHandler(r *gin.Engine) {
	r.GET("/health", func(c *gin.Context) {
		res := HealthCheckResponse{
			Status: "OK",
		}

		c.JSON(200, res)
	})
}
