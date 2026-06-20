package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware checks if the user is authenticated via a cookie.
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Try to get the "selfin_auth" cookie
		_, err := c.Cookie("selfin_auth")
		if err != nil {
			// If not found or error, redirect to login page for GET requests
			if c.Request.Method == http.MethodGet {
				// Prevent caching of authenticated routes
				c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
				c.Redirect(http.StatusFound, "/login")
				c.Abort()
				return
			}
			
			// For HTMX requests (POST, DELETE, PATCH, etc.)
			// We can return an HX-Redirect header
			if c.GetHeader("HX-Request") == "true" {
				c.Header("HX-Redirect", "/login")
				c.Status(http.StatusUnauthorized)
				c.Abort()
				return
			}

			// Generic fallback
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		c.Next()
	}
}
