package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
type Usuario struct {
	Nombre string `json:"nombre"`
	Email  string `json:"email"`
}

var Usuarios []Usuario
*/

func SetupRoutes(r *gin.Engine) {

	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		// Renderizar la plantilla index.html
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Title":   "Mi Aplicación",
			"Heading": "¡aplicacion web!",
			"Message": "Bienvenido a mi aplicación web con go.",
		})
	})

	r.Static("/static", "./static")
}
