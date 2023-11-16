package http

import (
	"github.com/KokopelliMusic/go-lib/logger"
	"github.com/KokopelliMusic/kalm/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SongRoutes(r *gin.RouterGroup, db *gorm.DB) {
	song := r.Group("/song")

	song.GET(":id", func(c *gin.Context)  {
		song := &models.Song{}

		id, found := c.Params.Get("id")

		if !found {
			err(c, 400, "no id")
			return
		}

		db.First(song, id)

		if song.ID == 0 {
			err(c, 404, "song not found")
			return
		}

		logger.Dump(song)

		c.JSON(200, gin.H{
			"message": "Huts!",
		})
	})
}