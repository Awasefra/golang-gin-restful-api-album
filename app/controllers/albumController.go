package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"myapp/app/repositories"
	"myapp/app/services"
	"myapp/helpers"
)

func GetAlbums(c *gin.Context) {
	albums := services.FetchAllAlbums()
	helpers.JSONResponse(c, http.StatusOK, "Albums fetched successfully", albums)
}

func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")
	album := services.FetchAlbumByID(id)
	if album == nil {
		helpers.ErrorResponse(c, http.StatusNotFound, "Album not found")
		return
	}
	helpers.JSONResponse(c, http.StatusOK, "Album fetched successfully", album)
}

func PostAlbums(c *gin.Context) {
	var newAlbum repositories.Album
	if err := c.ShouldBindJSON(&newAlbum); err != nil {
		helpers.ErrorResponse(c, http.StatusBadRequest, "Invalid request payload")
		return
	}
	services.CreateAlbum(newAlbum)
	helpers.JSONResponse(c, http.StatusCreated, "Album created successfully", newAlbum)
}
