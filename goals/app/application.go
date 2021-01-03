package app

import (
	"github.com/d5avard/goals/goals/controllers"
	"github.com/gin-gonic/gin"
)

// StartApplication Start Application
func StartApplication() {
	var r *gin.Engine = gin.Default()
	SetupRouter(r)
	r.Run(":3000")
}

// SetupRouter Setupe Router
func SetupRouter(r *gin.Engine) {
	controllers.AddGoalsRoutes(r)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Not found"})
	})
}
