package routes

import (
	fonts "github.com/eusebioleite/selfin/resources/fonts"
	js "github.com/eusebioleite/selfin/resources/js"
	styles "github.com/eusebioleite/selfin/resources/styles"
	"github.com/gin-gonic/gin"
)

func setupResources(r *gin.Engine) {
	resource := r.Group("/resources")
	{
		resource.GET("/htmx.min.js", js.GetHtmx)
		resource.GET("/fonts/inter.woff2", fonts.GetFontInter)
		resource.GET("/styles/main.css", styles.GetCss)
	}
}
