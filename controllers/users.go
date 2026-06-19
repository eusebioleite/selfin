package controllers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/eusebioleite/selfin/models"
	"github.com/gin-gonic/gin"
)

// GetUsers returns all users.
func GetUsers(c *gin.Context) {
	rows, err := DB.Query("SELECT id, name, login, password, image_url, enabled FROM users")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Login, &u.Password, &u.ImageURL, &u.Enabled); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		users = append(users, u)
	}
	c.JSON(http.StatusOK, users)
}

// GetUser returns a user by ID.
func GetUser(c *gin.Context) {
	id := c.Param("id")
	var u models.User
	err := DB.QueryRow("SELECT id, name, login, password, image_url, enabled FROM users WHERE id = ?", id).
		Scan(&u.ID, &u.Name, &u.Login, &u.Password, &u.ImageURL, &u.Enabled)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found."})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, u)
}

// CreateUser creates a new user.
func CreateUser(c *gin.Context) {
	var u models.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := DB.Exec("INSERT INTO users (name, login, password, image_url, enabled) VALUES (?, ?, ?, ?, ?)",
		u.Name, u.Login, u.Password, u.ImageURL, u.Enabled)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	id, _ := res.LastInsertId()
	u.ID = id
	c.JSON(http.StatusCreated, u)
}

// UpdateUser updates a user.
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var u models.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := DB.Exec("UPDATE users SET name = ?, login = ?, password = ?, image_url = ?, enabled = ? WHERE id = ?",
		u.Name, u.Login, u.Password, u.ImageURL, u.Enabled, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	idInt, _ := strconv.ParseInt(id, 10, 64)
	u.ID = idInt
	c.JSON(http.StatusOK, u)
}

// DeleteUser deletes a user.
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	_, err := DB.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "user deleted successfully."})
}
