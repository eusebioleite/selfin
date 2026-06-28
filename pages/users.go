package pages

import (
	"fmt"
	"net/http"

	"github.com/eusebioleite/selfin/repository"
	"github.com/eusebioleite/selfin/views"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	c.Header("Content-Type", "text/html; charset=utf-8")
	users, err := repository.GetUsers()
	if err != nil {
		c.AbortWithError(http.StatusNotFound, fmt.Errorf("Error loading users table: %w", err))
		return
	}
	views.Users(users).Render(c.Request.Context(), c.Writer)
}

func GetUsersModal(c *gin.Context) {
	c.Header("Content-Type", "text/html; charset=utf-8")
	views.UserModal().Render(c.Request.Context(), c.Writer)
}
