package pages

import (
	"fmt"
	"net/http"

	"github.com/eusebioleite/selfin/repository"
	"github.com/eusebioleite/selfin/views"
	"github.com/gin-gonic/gin"
)

func GetCategories(c *gin.Context) {
	c.Header("Content-Type", "text/html; charset=utf-8")
	categories, err := repository.GetCategories()
	if err != nil {
		c.AbortWithError(http.StatusNotFound, fmt.Errorf("Error loading categories table: %w", err))
		return
	}
	views.Categories(categories).Render(c.Request.Context(), c.Writer)
}
