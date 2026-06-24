package pages

import (
	"database/sql"
	"net/http"

	"github.com/eusebioleite/selfin/database"
	"github.com/eusebioleite/selfin/views"
	"github.com/gin-gonic/gin"
)

func GetLogin(c *gin.Context) {
	c.Header("Content-Type", "text/html; charset=utf-8")
	views.LoginPage().Render(c.Request.Context(), c.Writer)
}

func PostLogin(c *gin.Context) {
	login := c.PostForm("login")
	password := c.PostForm("password")

	var id int64
	err := database.DB.QueryRow("SELECT id FROM users WHERE login = ? AND password = ? AND enabled = 1", login, password).Scan(&id)

	if err != nil {
		if err == sql.ErrNoRows {
			c.String(http.StatusOK, `Invalid credentials or account disabled.`)
			return
		}
		c.String(http.StatusOK, `An error occurred during login.`)
		return
	}

	// Set auth cookie
	c.SetCookie("selfin_auth", "1", 3600*24*7, "/", "", false, true)

	c.Header("HX-Redirect", "/dashboard")
	c.Status(http.StatusOK)
}

func PostLogout(c *gin.Context) {
	// Clear auth cookie
	c.SetCookie("selfin_auth", "", -1, "/", "", false, true)
	c.Redirect(http.StatusFound, "/login")
}

func GetForgotPassword(c *gin.Context) {
	c.Header("Content-Type", "text/html; charset=utf-8")
	views.PasswordForgetPage().Render(c.Request.Context(), c.Writer)
}
