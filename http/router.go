package http

import (
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var r *gin.Engine

func Init(db *gorm.DB) *gin.Engine {
	r = gin.Default()

	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("[%s] %s - \"%s %s %s %d %s \"%s\" %s\"\n",
			param.TimeStamp.Format(time.RFC3339),
			param.ClientIP,
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	r.Use(gin.Recovery())

  r.Use(cors.New(cors.Config{
    AllowOrigins:     []string{"*"},
    AllowCredentials: true,
  }))

	api := r.Group("/api")

	SongRoutes(api, db)
	SpotifyRoutes(api, db)

	return r
}

func err(c *gin.Context, status int, msg string) {
	c.JSON(status, gin.H{
		"message": msg,
		"status": status,
	})
}