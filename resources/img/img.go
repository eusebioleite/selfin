package img

import (
	_ "embed"
	"net/http"

	"github.com/gin-gonic/gin"
)

//go:embed output.png
var logo []byte

func GetLogo(c *gin.Context) {
	c.Header("Cache-Control", "public, max-age=31536000")

	c.Data(http.StatusOK, "image/png", logo)
}
