package main

import (
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
)

func HomePage(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Hello World",
	})
}

func PostHomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Post Home Page",
	})
}

func QueryString(c *gin.Context) {
	name := c.Query("name")
	age := c.Query("age")
	c.JSON(200, gin.H{
		"name": name,
		"age":  age,
	})
}

func PathParameters(c *gin.Context) {
	name := c.Param("name")
	age := c.Param("age")
	c.JSON(200, gin.H{
		"name": name,
		"age":  age,
	})
}

func PostRequest(c *gin.Context) {
	body := c.Request.Body
	value, err := io.ReadAll(body)
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(200, gin.H{
		"message": string(value),
	})
}

func main() {
	r := gin.Default()
	// r.GET("/", func(ctx *gin.Context) {
	// 	ctx.JSON(200, gin.H{
	// 		"message": "Hello World",
	// 	})
	// })
	r.GET("/", HomePage)
	r.POST("/", PostHomePage)
	r.GET("/query", QueryString)              //query?name=ajinkya&age=28
	r.GET("/path/:name/:age", PathParameters) // /path/ajinkya/28
	r.POST("/post", PostRequest)              //to read in anything passed with the body
	r.Run(":8080")
}
