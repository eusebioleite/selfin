package pages

import (
	"net/http"

	"github.com/eusebioleite/selfin/database"
	"github.com/eusebioleite/selfin/models"
	"github.com/eusebioleite/selfin/views"
	"github.com/gin-gonic/gin"
)

func GetCategories(c *gin.Context) {
	rows, err := database.DB.Query("SELECT id, description FROM categories ORDER BY id DESC")
	if err != nil {
		c.String(http.StatusInternalServerError, "Error loading categories")
		return
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var cat models.Category
		if err := rows.Scan(&cat.ID, &cat.Description); err != nil {
			c.String(http.StatusInternalServerError, "Error parsing categories")
			return
		}
		categories = append(categories, cat)
	}

	c.Header("Content-Type", "text/html; charset=utf-8")
	views.CategoriesPage(categories).Render(c.Request.Context(), c.Writer)
}
