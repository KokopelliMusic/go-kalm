package http

import (
	"github.com/KokopelliMusic/go-lib/logger"
	"github.com/KokopelliMusic/kalm/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SpotifyRoutes(r *gin.RouterGroup, db *gorm.DB) {
	sp := r.Group("/spotify")

	sp.GET(":id", func(c *gin.Context)  {
		spotify := &models.SpotifyPlaylist{}

		id, found := c.Params.Get("id")

		if !found {
			err(c, 400, "no id")
			return
		}

		db.First(spotify, id)

		if spotify.ID == 0 {
			err(c, 404, "spotify playlist not found")
			return
		}

		logger.Dump(spotify)

		c.JSON(200, gin.H{
			"playlist": spotify,
		})
	})

	sp.GET("/spotify-id/:id", func(c *gin.Context) {
		spotify := &models.SpotifyPlaylist{}

		id, found := c.Params.Get("id")

		if !found {
			err(c, 400, "no id")
			return
		}

		db.First(spotify, "spotify_id = ?", id)

		if spotify.ID == 0 {
			err(c, 404, "spotify playlist not found")
			return
		}

		logger.Dump(spotify)

		c.JSON(200, gin.H{
			"playlist": spotify,
		})
	})
}

