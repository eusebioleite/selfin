package pages

import (
	"fmt"
	"net/http"

	"github.com/eusebioleite/selfin/repository"
	"github.com/eusebioleite/selfin/views"
	"github.com/gin-gonic/gin"
)

func GetTransactions(c *gin.Context) {
	c.Header("Content-Type", "text/html; charset=utf-8")
	transactions, err := repository.GetTransactionsView()
	if err != nil {
		c.AbortWithError(http.StatusNotFound, fmt.Errorf("Error loading transactions table: %w", err))
		return
	}
	views.Transactions(transactions).Render(c.Request.Context(), c.Writer)
}

func GetTransactionsModal(c *gin.Context) {
	c.Header("Content-Type", "text/html; charset=utf-8")
	categories, err := repository.GetCategories()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("Error loading categories: %w", err))
		return
	}
	users, err := repository.GetUsers()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("Error loading users: %w", err))
		return
	}
	views.TransactionModal(categories, users).Render(c.Request.Context(), c.Writer)
}
