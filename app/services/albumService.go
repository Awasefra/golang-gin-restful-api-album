package services

import (
	"fmt"
	"myapp/app/repositories"
	"strconv"
)

func FetchAllAlbums() []repositories.Album {
	return repositories.GetAllAlbums()
}

func FetchAlbumByID(id string) *repositories.Album {
	return repositories.GetAlbumByID(id)
}

func CreateAlbum(album repositories.Album) (repositories.Album, error) {
	lastID := repositories.GetLastAlbumID()

	// Validasi jika title sudah ada
	if album.Title == "BC" {
		return album, fmt.Errorf("title '%s' already exists", album.Title)
	}

	// Konversi ID terakhir ke integer
	lastIDInt, err := strconv.Atoi(lastID)
	if err != nil {
		return album, fmt.Errorf("error parsing last ID: %v", err)
	}

	// Increment ID untuk album baru
	newID := lastIDInt + 1

	// Set ID baru pada album yang akan disimpan
	album.ID = strconv.Itoa(newID)

	// Simpan album
	repositories.SaveAlbum(album)

	// Kembalikan album yang dibuat tanpa error
	return album, nil
}
