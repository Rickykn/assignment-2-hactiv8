package routers

import (
	"github.com/Rickykn/assignment-2-hactiv8.git/handlers"
	"github.com/gin-gonic/gin"
)

func Server() *gin.Engine {
	engine := gin.New()
	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	engine.POST("/orders", handlers.CreateOrder)
	engine.GET("/orders", handlers.GetAllOrder)
	engine.DELETE("/orders/:id", handlers.DeleteOrder)
	return engine
}
