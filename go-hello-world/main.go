package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	// Load templates (SSR). Gin uses Go's html/template under the hood.
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title":   "Hello World",
			"message": "Hello World",
		})
	})

	return r
}

func main() {
	r := setupRouter()
	// Default: localhost:8080
	r.Run(":8080")
}
