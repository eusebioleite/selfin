package controllers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/eusebioleite/selfin/models"
	"github.com/gin-gonic/gin"
)

// GetTransactions returns all transactions.
func GetTransactions(c *gin.Context) {
	rows, err := DB.Query("SELECT id, date, amount, description, category_id, user_id FROM transactions")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var transactions []models.Transaction
	for rows.Next() {
		var t models.Transaction
		if err := rows.Scan(&t.ID, &t.Date, &t.Amount, &t.Description, &t.CategoryID, &t.UserID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		transactions = append(transactions, t)
	}
	c.JSON(http.StatusOK, transactions)
}

// GetTransaction returns a transaction by ID.
func GetTransaction(c *gin.Context) {
	id := c.Param("id")
	var t models.Transaction
	err := DB.QueryRow("SELECT id, date, amount, description, category_id, user_id FROM transactions WHERE id = ?", id).
		Scan(&t.ID, &t.Date, &t.Amount, &t.Description, &t.CategoryID, &t.UserID)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "transaction not found."})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, t)
}

// CreateTransaction creates a new transaction.
func CreateTransaction(c *gin.Context) {
	var t models.Transaction
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := DB.Exec("INSERT INTO transactions (date, amount, description, category_id, user_id) VALUES (?, ?, ?, ?, ?)",
		t.Date, t.Amount, t.Description, t.CategoryID, t.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	id, _ := res.LastInsertId()
	t.ID = id
	c.JSON(http.StatusCreated, t)
}

// UpdateTransaction updates a transaction.
func UpdateTransaction(c *gin.Context) {
	id := c.Param("id")
	var t models.Transaction
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := DB.Exec("UPDATE transactions SET date = ?, amount = ?, description = ?, category_id = ?, user_id = ? WHERE id = ?",
		t.Date, t.Amount, t.Description, t.CategoryID, t.UserID, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	idInt, _ := strconv.ParseInt(id, 10, 64)
	t.ID = idInt
	c.JSON(http.StatusOK, t)
}

// DeleteTransaction deletes a transaction.
func DeleteTransaction(c *gin.Context) {
	id := c.Param("id")
	_, err := DB.Exec("DELETE FROM transactions WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "transaction deleted successfully."})
}
