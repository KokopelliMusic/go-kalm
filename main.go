package main

import (
	"github.com/KokopelliMusic/go-lib/logger"
	"github.com/KokopelliMusic/kalm/http"
	"github.com/KokopelliMusic/kalm/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func main() {
	logger.Debug(true)

	db, err = gorm.Open(sqlite.Open("data.db"), &gorm.Config{})

	if err != nil {
		panic("failed to open database")
	}

	db.AutoMigrate(&models.Song{})
	db.AutoMigrate(&models.Artist{})

	router := http.Init(db)

	router.Run(":8081")
}