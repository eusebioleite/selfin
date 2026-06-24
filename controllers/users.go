package controllers

import (
	"net/http"
	"strconv"

	"github.com/eusebioleite/selfin/models"
	repo "github.com/eusebioleite/selfin/repository"
	"github.com/gin-gonic/gin"
)

// GetUsers returns a json with users
func GetUsers(c *gin.Context) {
	users, err := repo.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

// GetUser returns a user by id
func GetUser(c *gin.Context) {
	// 1. parses the id parameter to string
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"info": "error converting parameter to int64.", "error": err})
		return
	}

	// 2. get user from database
	user, err := repo.GetUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, user)
}

// CreateUser creates a new user.
func CreateUser(c *gin.Context) {
	// 1. initializes a struct with data from payload
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"info": "error binding payload to struct", "error": err.Error()})
		return
	}

	// 2. creates a new user
	err := repo.NewUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// UpdateUser updates a user.
func UpdateUser(c *gin.Context) {
	// 1. parses the id parameter to string
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"info": "error converting parameter to int64.", "error": err})
		return
	}

	// 2. initializes a struct with data from payload
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 3. updates a user
	user.ID = id
	err = repo.UpdateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}

	c.JSON(http.StatusOK, user)
}

// DeleteUser deletes a user.
func DeleteUser(c *gin.Context) {
	// 1. parses the id parameter to string
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error converting parameter to int64.", "log": err})
		return
	}

	// 2. deletes the user
	err = repo.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "user deleted successfully."})
}
