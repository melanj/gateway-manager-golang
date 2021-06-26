package main

import (
	"awesomeProject/controllers"
	"awesomeProject/models"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"path"
	"path/filepath"
)

func main() {

	gin.DisableConsoleColor()

	// Logging to a file.
	f, _ := os.Create("app.log")
	gin.DefaultWriter = io.MultiWriter(f)

	r := gin.Default()
	models.ConnectDataBase()

	r.NoRoute(func(c *gin.Context) {
		dir, file := path.Split(c.Request.RequestURI)
		ext := filepath.Ext(file)
		if file == "" || ext == "" {
			c.File("../html/index.html")
		} else {
			c.File("../html" + path.Join(dir, file))
		}
	})

	r.GET("/api/gateways", controllers.FindGateways)
	r.POST("/api/gateways", controllers.CreateGateway)
	r.GET("/api/gateways/:id", controllers.FindGateway)
	r.DELETE("/api/gateways/:id", controllers.DeleteGateway)

	r.Run(":8080")
}
