package routes

import (
	"creator-hub-gin/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())

	contentGroup := router.Group("/contents")
	{
		contentGroup.GET("", controllers.GetContents)
		contentGroup.GET("/:id", controllers.GetContentDetail)
		contentGroup.POST("", controllers.CreateContent)
		contentGroup.PUT("/:id", controllers.UpdateContent)
		contentGroup.DELETE("/:id", controllers.DeleteContent)
	}

	return router
}
