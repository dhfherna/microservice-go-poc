package services

import "github.com/dhfherna/microservice-go-poc/entity"

var albums = []entity.AlbumEntity{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func GetAlbums() []entity.AlbumEntity {
	return albums
}

func GetAlbumById(id string) *entity.AlbumEntity {
	for _, album := range albums {
		if album.ID == id {
			return &album
		}
	}
	return nil
}

func AddAlbum(album entity.AlbumEntity) bool {
	albums = append(albums, album)
	return true
}

func UpdateAlbum(id string, album entity.AlbumEntity) bool {
	for i, a := range albums {
		if a.ID == id {
			albums[i] = album
			return true
		}
	}
	return false
}

func DeleteAlbum(id string) bool {
	for i, a := range albums {
		if a.ID == id {
			albums = append(albums[:i], albums[i+1:]...)
			return true
		}
	}
	return false
}
