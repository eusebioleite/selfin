package js

import (
	_ "embed"

	"github.com/gin-gonic/gin"
)

//go:embed htmx.min.js
var htmxJS []byte

func GetHtmx(c *gin.Context) {
	c.Header("Content-Type", "application/javascript")
	c.Header("Cache-Control", "public, max-age=31536000")
	c.Writer.Write(htmxJS)
}
