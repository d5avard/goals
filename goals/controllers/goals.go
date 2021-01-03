package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GoalsController Struct
type GoalsController struct{}

// AddGoalsRoutes Add Goals Route
func AddGoalsRoutes(r *gin.Engine) {
	var rg *gin.RouterGroup = r.Group("api/goals")
	var gc *GoalsController = new(GoalsController)
	rg.GET("/", gc.All)
	rg.GET("/:id", gc.Read)
	rg.POST("/", gc.Create)
	rg.PUT("/:id", gc.Update)
	rg.DELETE("/:id", gc.Delete)
}

// All Get all goals
func (g *GoalsController) All(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Get all goals"})
}

// Read Read a goal
func (g *GoalsController) Read(c *gin.Context) {
	var id string = c.Param("id")
	var v string = fmt.Sprintf("%s %s", "Get goal", id)
	c.JSON(http.StatusOK, gin.H{"message": v})
}

// Create Create a Goal
func (g *GoalsController) Create(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"message": "Create a new goal"})
}

// Update Update a Goal
func (g *GoalsController) Update(c *gin.Context) {
	var id string = c.Param("id")
	var v string = fmt.Sprintf("%s %s", "Update goal", id)
	c.JSON(http.StatusOK, gin.H{"message": v})
}

// Delete Delete a Goal
func (g *GoalsController) Delete(c *gin.Context) {
	var id string = c.Param("id")
	var v string = fmt.Sprintf("%s %s", "Delete goal", id)
	c.JSON(http.StatusOK, gin.H{"message": v})
}
