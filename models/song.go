package models

import "gorm.io/gorm"

type Song struct {
	gorm.Model
	Title string
	Artists []Artist `gorm:"many2many:song_artists"`
}
