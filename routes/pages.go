package routes

import (
	"github.com/eusebioleite/selfin/pages"
	"github.com/gin-gonic/gin"
)

func setupPages(r *gin.Engine) {
	page := r.Group("/")
	{

		page.GET("/", pages.GetHomePage)
		page.POST("/login", pages.DoLogin)
	}
}
