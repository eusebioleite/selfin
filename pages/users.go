package pages

import (
	"net/http"
	"strconv"

	"github.com/eusebioleite/selfin/controllers"
	"github.com/eusebioleite/selfin/models"
	"github.com/eusebioleite/selfin/views"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	rows, err := controllers.DB.Query("SELECT id, name, login, password, image_url, enabled FROM users ORDER BY id DESC")
	if err != nil {
		c.String(http.StatusInternalServerError, "Error loading users")
		return
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Login, &u.Password, &u.ImageURL, &u.Enabled); err != nil {
			c.String(http.StatusInternalServerError, "Error parsing users")
			return
		}
		users = append(users, u)
	}

	c.Header("Content-Type", "text/html; charset=utf-8")
	views.UsersPage(users).Render(c.Request.Context(), c.Writer)
}

func PostUser(c *gin.Context) {
	name := c.PostForm("name")
	login := c.PostForm("login")
	password := c.PostForm("password")
	imageUrl := c.PostForm("image_url")
	enabledStr := c.PostForm("enabled")
	enabled := enabledStr == "true"

	if name == "" || login == "" || password == "" {
		c.String(http.StatusBadRequest, "Missing required fields")
		return
	}

	res, err := controllers.DB.Exec(
		"INSERT INTO users (name, login, password, image_url, enabled) VALUES (?, ?, ?, ?, ?)",
		name, login, password, imageUrl, enabled,
	)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error creating user")
		return
	}

	id, _ := res.LastInsertId()
	u := models.User{
		ID:       id,
		Name:     name,
		Login:    login,
		Password: password,
		ImageURL: imageUrl,
		Enabled:  enabled,
	}

	c.Header("Content-Type", "text/html; charset=utf-8")
	views.UserRow(u).Render(c.Request.Context(), c.Writer)
}

func ToggleUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid ID")
		return
	}

	// Fetch current state
	var enabled bool
	err = controllers.DB.QueryRow("SELECT enabled FROM users WHERE id = ?", id).Scan(&enabled)
	if err != nil {
		c.String(http.StatusNotFound, "User not found")
		return
	}

	// Toggle
	newEnabled := !enabled
	_, err = controllers.DB.Exec("UPDATE users SET enabled = ? WHERE id = ?", newEnabled, id)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error updating user")
		return
	}

	// We only need the ID and new state to re-render the toggle
	u := models.User{
		ID:      id,
		Enabled: newEnabled,
	}

	c.Header("Content-Type", "text/html; charset=utf-8")
	views.UserEnabledToggle(u).Render(c.Request.Context(), c.Writer)
}
