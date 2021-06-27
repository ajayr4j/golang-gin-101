package main

import (
	"go-projects/controller"
	"go-projects/middlewares"
	"go-projects/service"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.Log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {

	setupLogOutput()
	server := gin.Default()

	gin.ForceConsoleColor()
	
	server.Static("/css","./templates/css")

	server.LoadHTMLGlob("templates/*.html")

	server.Use(middleware.BasicAuth())


	apiRoutes := server.Group("/api")

	{
	apiRoutes.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})

	apiRoutes.POST("/videos", func(ctx *gin.Context) {
		err := videoController.Save(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message":"Video Input is Valid!!"})
		}
		ctx.JSON(200, videoController.Save(ctx))
	})
	}

	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/videos", videoController.ShowAll)
	}

	port := os.Getenv("PORT")

	if port == ""{

		port = "5000"

	}
	server.Run(":"+port)
}
