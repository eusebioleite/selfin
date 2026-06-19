package fonts

import (
	_ "embed"

	"github.com/gin-gonic/gin"
)

//go:embed Inter-Variable.woff2
var interFont []byte

func GetFontInter(c *gin.Context) {
	c.Header("Content-Type", "font/woff2")
	c.Header("Cache-Control", "public, max-age=31536000")
	c.Writer.Write(interFont)
}
