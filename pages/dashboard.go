package pages

import (
	"github.com/eusebioleite/selfin/views"
	"github.com/gin-gonic/gin"
)

func GetDashboard(c *gin.Context) {
	c.Header("Content-Type", "text/html; charset=utf-8")
	views.DashboardPage().Render(c.Request.Context(), c.Writer)
}
