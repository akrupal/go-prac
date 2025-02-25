package main

import (
	"gin-poc/controller"
	middleware "gin-poc/middlewares"
	"gin-poc/service"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	setupLogOutput() //if you want to store the logs to a file
	// server := gin.Default()
	// if you go inside the Default function you will see 2 default middlewares
	// logger and recovery
	// the same can be written as
	server := gin.New()
	// server.Use(gin.Recovery(), gin.Logger())
	// lets say we want to write our own logger
	server.Use(gin.Recovery(), middleware.Logger(), middleware.BasicAuth())
	server.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "test",
		})
	})
	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})
	server.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.Save(ctx))
	})
	server.Run(":8080")
}
