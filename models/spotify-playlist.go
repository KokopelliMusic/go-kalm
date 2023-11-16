package models

import "gorm.io/gorm"

type SpotifyPlaylist struct {
	gorm.Model
	SpotifyID string
	Favorite bool
}