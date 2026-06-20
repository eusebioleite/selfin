package resources

import (
	_ "embed"

	"github.com/gin-gonic/gin"
)

//go:embed main.css
var cssMain []byte

func GetCss(c *gin.Context) {
	c.Header("Content-Type", "text/css; charset=utf-8")
	c.Header("Cache-Control", "public, max-age=31536000")
	c.Writer.Write(cssMain)
}
