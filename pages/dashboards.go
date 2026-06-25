package pages

import (
	"github.com/eusebioleite/selfin/views"
	"github.com/gin-gonic/gin"
)

func GetDashboards(c *gin.Context) {
	c.Header("Content-Type", "text/html; charset=utf-8")
	views.Dashboards().Render(c.Request.Context(), c.Writer)
}
