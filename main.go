package main

import (
	"embed"
	"os"

	"github.com/agoodmu/demo/handlers"
	"github.com/gin-gonic/gin"
)

//go:embed templates
var templatesFolder embed.FS

var listen_port = "8080"

func init() {
	port, exists := os.LookupEnv("Port")
	if exists {
		listen_port = port
	}
}

func main() {
	router := gin.Default()

	v := router.Group("/")
	{
		v.GET("/echo", handlers.Echo(&templatesFolder, "templates/demo_request.tmpl"))
		v.StaticFile("/css/main.css", "css/main.css")
		router.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}

	v1 := router.Group("/v1")
	{
		v1.GET("/echo", handlers.Echo(&templatesFolder, "templates/demo_request.tmpl"))
	}
	router.Run("0.0.0.0:" + listen_port)
}
