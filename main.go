package main

import (
	"embed"
	"os"

	"github.com/agoodmu/demo/handlers"
	"github.com/gin-gonic/gin"
)

//go:embed templates
var templatesFolder embed.FS

var listen_port = "80"

func init() {
	port, exists := os.LookupEnv("Port")
	if exists {
		listen_port = port
	}
}

func main() {
	router := gin.Default()

	{
		v1 := router.Group("/v1")
		v1.GET("/echo", handlers.Echo(&templatesFolder, "templates/demo_http_headers.tmpl"))
	}

	router.StaticFile("/css/main.css", "css/main.css")

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run("0.0.0.0:" + listen_port)
}
