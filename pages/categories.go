package pages

import (
	"net/http"
	"strconv"

	"github.com/eusebioleite/selfin/controllers"
	"github.com/eusebioleite/selfin/models"
	"github.com/eusebioleite/selfin/views"
	"github.com/gin-gonic/gin"
)

func GetCategories(c *gin.Context) {
	rows, err := controllers.DB.Query("SELECT id, description FROM categories ORDER BY id DESC")
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

func PostCategory(c *gin.Context) {
	description := c.PostForm("description")
	if description == "" {
		c.String(http.StatusBadRequest, "Description is required")
		return
	}

	res, err := controllers.DB.Exec("INSERT INTO categories (description) VALUES (?)", description)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error creating category")
		return
	}

	id, _ := res.LastInsertId()
	cat := models.Category{
		ID:          id,
		Description: description,
	}

	c.Header("Content-Type", "text/html; charset=utf-8")
	views.CategoryRow(cat).Render(c.Request.Context(), c.Writer)
}

func DeleteCategory(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid ID")
		return
	}

	_, err = controllers.DB.Exec("DELETE FROM categories WHERE id = ?", id)
	if err != nil {
		// Possibly a foreign key constraint issue (associated transactions)
		c.String(http.StatusConflict, `<div class="p-2 bg-red-100 text-red-600 rounded">Cannot delete category with associated transactions.</div>`)
		return
	}

	// Return empty 200 OK to remove the element via HTMX
	c.Status(http.StatusOK)
}
