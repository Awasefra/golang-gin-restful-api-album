package routes

import "github.com/gin-gonic/gin"

func InitRoutes(router *gin.Engine) {
	api := router.Group("/api")
	AlbumRoutes(api)
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, your application is up and running!",
		})
	})
}
