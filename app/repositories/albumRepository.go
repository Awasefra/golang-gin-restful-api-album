package repositories

type Album struct {
	ID     string
	Title  string
	Artist string
	Price  float64
}

var albums = []Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan", Artist: "Sarah Vaughan", Price: 39.99},
}

func GetAllAlbums() []Album {
	return albums
}

func GetAlbumByID(id string) *Album {
	for _, album := range albums {
		if album.ID == id {
			return &album
		}
	}
	return nil
}

func SaveAlbum(album Album) {
	albums = append(albums, album)
}

func GetLastAlbumID() string {
	if len(albums) == 0 {
		return "" // Jika album kosong, kembalikan string kosong
	}
	return albums[len(albums)-1].ID
}
