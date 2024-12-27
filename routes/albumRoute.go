package routes

import (
	"github.com/gin-gonic/gin"
	"myapp/app/controllers"
)

func AlbumRoutes(router *gin.RouterGroup) {
	router.GET("/albums", controllers.GetAlbums)
	router.GET("/albums/:id", controllers.GetAlbumByID)
	router.POST("/albums", controllers.PostAlbums)
}
