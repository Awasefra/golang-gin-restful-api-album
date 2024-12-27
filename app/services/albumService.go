package services

import (
	"myapp/app/repositories"
)

func FetchAllAlbums() []repositories.Album {
	return repositories.GetAllAlbums()
}

func FetchAlbumByID(id string) *repositories.Album {
	return repositories.GetAlbumByID(id)
}

func CreateAlbum(album repositories.Album) {
	repositories.SaveAlbum(album)
}
