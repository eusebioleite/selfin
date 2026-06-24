package controllers

import (
	"net/http"
	"strconv"

	"github.com/eusebioleite/selfin/models"
	repo "github.com/eusebioleite/selfin/repository"
	"github.com/gin-gonic/gin"
)

// GetTransactions returns all transactions.
func GetTransactions(c *gin.Context) {
	transactions, err := repo.GetTransactions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, transactions)
}

// GetTransaction returns a transaction by id
func GetTransaction(c *gin.Context) {
	// 1. parses the id parameter to string
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"info": "error converting parameter to int64.", "error": err})
		return
	}

	// 2. get transaction from database
	transaction, err := repo.GetTransaction(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, transaction)
}

// CreateTransaction creates a new transaction.
func CreateTransaction(c *gin.Context) {
	// 1. initializes a struct with data from payload
	var transaction models.Transaction
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"info": "error binding payload to struct", "error": err.Error()})
		return
	}

	// 2. creates a new transaction
	err := repo.NewTransaction(&transaction)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, transaction)
}

// UpdateTransaction updates a transaction.
func UpdateTransaction(c *gin.Context) {
	// 1. parses the id parameter to string
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"info": "error converting parameter to int64.", "error": err})
		return
	}

	// 2. initializes a struct with data from payload
	var transaction models.Transaction
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 3. updates a transaction
	transaction.ID = id
	err = repo.UpdateTransaction(&transaction)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}

	c.JSON(http.StatusOK, transaction)
}

// DeleteTransaction deletes a transaction.
func DeleteTransaction(c *gin.Context) {
	// 1. parses the id parameter to string
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error converting parameter to int64.", "log": err})
		return
	}

	// 2. deletes the transaction
	err = repo.DeleteTransaction(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "transaction deleted successfully."})
}
