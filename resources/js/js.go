package js

import (
	_ "embed"

	"github.com/gin-gonic/gin"
)

//go:embed htmx.min.js
var htmxJS []byte

//go:embed json-enc.js
var jsonEncJS []byte

//go:embed alpine.min.js
var alpineJS []byte

func GetHtmx(c *gin.Context) {
	c.Header("Content-Type", "application/javascript")
	c.Header("Cache-Control", "public, max-age=31536000")
	c.Writer.Write(htmxJS)
}

func GetJsonEnc(c *gin.Context) {
	c.Header("Content-Type", "application/javascript")
	c.Header("Cache-Control", "public, max-age=31536000")
	c.Writer.Write(jsonEncJS)
}

func GetAlpine(c *gin.Context) {
	c.Header("Content-Type", "application/javascript")
	c.Header("Cache-Control", "public, max-age=31536000")
	c.Writer.Write(alpineJS)
}
