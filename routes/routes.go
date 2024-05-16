package routes

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	r.Static("/static", "./static")

	r.LoadHTMLGlob("templates/*html")

	r.GET("/", func(c *gin.Context) {
		// Renderizar la plantilla index.html
		c.HTML(http.StatusOK, "index.html", nil)

	})

	r.GET("/:page", func(c *gin.Context) {
		page := c.Param("page")

		if !strings.HasSuffix(page, ".html") {
			page += ".html"
		}

		if _, err := os.Stat("templates/" + page); err == nil {
			c.HTML(http.StatusOK, page, nil)
		} else {
			c.HTML(http.StatusNotFound, "404.html", nil)
		}

	})
}
