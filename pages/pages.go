package pages

import (
	"net/http"
	"time"

	"github.com/eusebioleite/selfin/views"
	"github.com/gin-gonic/gin"
)

func GetHomePage(c *gin.Context) {
	c.Header("Content-Type", "text/html; charset=utf-8")
	views.LoginPage().Render(c.Request.Context(), c.Writer)
}

func DoLogin(c *gin.Context) {
	time.Sleep(1 * time.Second)
	c.Header("Content-Type", "text/html; charset=utf-8")
	c.String(http.StatusOK, `<div class="success-msg">Login simulado com sucesso!</div>`)
}
