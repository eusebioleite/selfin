package pages

import (
	"net/http"

	repo "github.com/eusebioleite/selfin/repository"
	"github.com/eusebioleite/selfin/views"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	users, err := repo.GetUsers()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.Header("Content-Type", "text/html; charset=utf-8")
	views.UsersPage(users).Render(c.Request.Context(), c.Writer)
}
