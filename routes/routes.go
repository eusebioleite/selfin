package routes

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		c.Next()

		latency := time.Since(start)

		if len(c.Errors) > 0 {
			for _, e := range c.Errors.Errors() {
				log.Error().Str("error", e).Msg(c.Request.URL.Path)
			}
		} else {
			log.Info().
				Int("status", c.Writer.Status()).
				Str("method", c.Request.Method).
				Str("path", path).
				Str("query", query).
				Str("ip", c.ClientIP()).
				Dur("latency", latency).
				Msg("Request")
		}
	}
}

func SetupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(LoggerMiddleware(), gin.Recovery())

	setupApi(r)
	setupPages(r)
	setupResources(r)

	return r
}
