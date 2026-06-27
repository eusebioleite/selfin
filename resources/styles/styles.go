package resources

import (
	_ "embed"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

//go:embed main.css
var cssMain []byte

const cssPath = "resources/styles/main.css"

func GetCss(c *gin.Context) {
	c.Header("Content-Type", "text/css; charset=utf-8")

	if os.Getenv("DEV") == "1" {
		c.Header("Cache-Control", "no-cache")
		http.ServeFile(c.Writer, c.Request, cssPath)
		return
	}

	c.Header("Cache-Control", "public, max-age=31536000")
	c.Writer.Write(cssMain)
}
