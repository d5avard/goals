package main

import (
	"github.com/d5avard/goals/goals/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	var r *gin.Engine = gin.Default()
	var gg *gin.RouterGroup = r.Group("api/goals")

	var gc *controllers.GoalsController = new(controllers.GoalsController)
	gg.GET("/", gc.All)
	gg.GET("/:id", gc.Read)
	gg.POST("/", gc.Create)
	gg.PUT("/:id", gc.Update)
	gg.DELETE("/:id", gc.Delete)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Not found"})
	})

	r.Run(":3000")
}
