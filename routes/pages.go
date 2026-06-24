package routes

import (
	"net/http"

	"github.com/eusebioleite/selfin/pages"
	"github.com/gin-gonic/gin"
)

func setupPages(r *gin.Engine) {

	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/dashboard")
	})

	public := r.Group("/")
	{
		public.GET("/login", pages.GetLogin)
		public.POST("/login", pages.PostLogin)
		public.GET("/forgot-password", pages.GetForgotPassword)
	}

	protected := r.Group("/")
	protected.Use(AuthMiddleware())
	{
		protected.POST("/logout", pages.PostLogout)

		protected.GET("/dashboard", pages.GetDashboard)

		protected.GET("/categories", pages.GetCategories)
		protected.POST("/categories", pages.PostCategory)
		protected.DELETE("/categories/:id", pages.DeleteCategory)

		protected.GET("/users", pages.GetUsers)
		protected.POST("/users", pages.PostUser)
	}
}
