package controllers

import (
	"net/http"
	"strconv"

	"github.com/eusebioleite/selfin/models"
	"github.com/gin-gonic/gin"
)

// GetCategories returns all categories.
func GetCategories(c *gin.Context) {
	rows, err := DB.Query("SELECT id, description FROM categories")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var g models.Category
		if err := rows.Scan(&g.ID, &g.Description); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		categories = append(categories, g)
	}
	c.JSON(http.StatusOK, categories)
}

// GetCategory returns a category by ID.
func GetCategory(c *gin.Context) {
	id := c.Param("id")
	var g models.Category
	err := DB.QueryRow("SELECT id, description FROM categories WHERE id = ?", id).Scan(&g.ID, &g.Description)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "category not found."})
		return
	}
	c.JSON(http.StatusOK, g)
}

// CreateCategory creates a new category.
func CreateCategory(c *gin.Context) {
	var g models.Category
	if err := c.ShouldBindJSON(&g); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := DB.Exec("INSERT INTO categories (description) VALUES (?)", g.Description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	id, _ := res.LastInsertId()
	g.ID = id
	c.JSON(http.StatusCreated, g)
}

// UpdateCategory updates a category.
func UpdateCategory(c *gin.Context) {
	id := c.Param("id")
	var g models.Category
	if err := c.ShouldBindJSON(&g); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := DB.Exec("UPDATE categories SET description = ? WHERE id = ?", g.Description, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	idInt, _ := strconv.ParseInt(id, 10, 64)
	g.ID = idInt
	c.JSON(http.StatusOK, g)
}

// DeleteCategory deletes a category.
func DeleteCategory(c *gin.Context) {
	id := c.Param("id")
	_, err := DB.Exec("DELETE FROM categories WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "category deleted successfully."})
}
