package controllers

import (
	"net/http"

	"myapp/app/repositories"
	"myapp/app/services"
	"myapp/helpers"

	"github.com/gin-gonic/gin"
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

	// Bind JSON payload ke objek newAlbum
	if err := c.ShouldBindJSON(&newAlbum); err != nil {
		helpers.ErrorResponse(c, http.StatusBadRequest, "Invalid request payload")
		return
	}
	// TRY AND CATCH
	// Panggil service untuk membuat album baru
	albumCreated, err := services.CreateAlbum(newAlbum)
	if err != nil {
		helpers.ErrorResponse(c, http.StatusBadRequest, err.Error()) // Kirim error ke client
		return
	}

	// Kembalikan response JSON dengan album yang berhasil dibuat
	helpers.JSONResponse(c, http.StatusCreated, "Album created successfully", albumCreated)
}
