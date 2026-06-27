package routes

import (
	"github.com/eusebioleite/selfin/controllers"
	"github.com/gin-gonic/gin"
)

func setupApi(r *gin.Engine) {
	api := r.Group("/api/v1")
	{
		api.GET("/users", controllers.GetUsers)
		api.GET("/users/:id", controllers.GetUser)
		api.POST("/users", controllers.CreateUser)
		api.PUT("/users/:id", controllers.UpdateUser)
		api.DELETE("/users/:id", controllers.DeleteUser)

		api.GET("/categories", controllers.GetCategories)
		api.GET("/categories/:id", controllers.GetCategory)
		api.POST("/categories", controllers.CreateCategory)
		api.PUT("/categories/:id", controllers.UpdateCategory)
		api.DELETE("/categories/:id", controllers.DeleteCategory)

		api.GET("/transactions", controllers.GetTransactions)
		api.GET("/transactions/:id", controllers.GetTransaction)
		api.POST("/transactions", controllers.CreateTransaction)
		api.PUT("/transactions/:id", controllers.UpdateTransaction)
		api.DELETE("/transactions/:id", controllers.DeleteTransaction)
	}
}
