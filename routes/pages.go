package routes

import (
	"net/http"

	"github.com/eusebioleite/selfin/pages"
	"github.com/eusebioleite/selfin/security"
	"github.com/gin-gonic/gin"
)

func setupPages(r *gin.Engine) {

	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/dashboards")
	})

	public := r.Group("/")
	{
		public.GET("/auth", pages.GetAuth)
		public.POST("/auth", func(c *gin.Context) {
			security.AuthHandler(c)
		})
	}

	protected := r.Group("/")
	protected.Use(security.AuthMiddleware())
	{
		protected.GET("/dashboards", pages.GetDashboards)
		protected.GET("/transactions", pages.GetTransactions)
		protected.GET("/transactions/new", pages.GetTransactionsModal)
		protected.GET("/categories", pages.GetCategories)
		protected.GET("/categories/new", pages.GetCategoriesModal)
		protected.GET("/users", pages.GetUsers)
		protected.GET("/users/new", pages.GetUsersModal)
	}
}
