package controllers

import (
	"net/http"
	"strconv"

	"github.com/eusebioleite/selfin/models"
	repo "github.com/eusebioleite/selfin/repository"
	"github.com/gin-gonic/gin"
)

type ArgonParams struct {
	Memory      uint32
	Iterations  uint32
	Parallelism uint8
	SaltLength  uint32
	KeyLength   uint32
}

func Auth(c *gin.Context) {
	// 1. check if user exists
	login := c.Param("login")
	user, err := repo.GetUserByLogin(login)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	// 2. validate the password

	// 2.1. set argon params
	params := &ArgonParams {
		Memory:      64 * 1024,
		Iterations:  3,
		Parallelism: 4,
		SaltLength:  16,
		KeyLength:   32,
	}

	// 2.2. generate hash
	has,
	c.JSON(http.StatusOK, user)
}