package routes

import (
	"github.com/alirezasaharkhiz/FlagChain/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterFlagRoutes(router *gin.Engine, flagController *controllers.FeatureFlagController) {
	api := router.Group("/api")
	{
		api.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
		api.POST("/flags", flagController.Create)
		api.PUT("/flags/:id/toggle", flagController.Toggle)
		api.GET("/flags", flagController.List)
		api.GET("/flags/:id/history", flagController.History)
		api.POST("/flags/:id/dependencies", flagController.AddDependency)
	}
}
