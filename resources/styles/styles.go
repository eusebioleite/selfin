package resources

import (
	_ "embed"

	"github.com/gin-gonic/gin"
)

//go:embed login.css
var cssLogin []byte

func GetCssLogin(c *gin.Context) {
	c.Header("Content-Type", "text/css; charset=utf-8")
	c.Header("Cache-Control", "public, max-age=31536000")
	c.Writer.Write(cssLogin)
}
