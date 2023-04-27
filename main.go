package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	r.GET("/", HomepageHandler)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func HomepageHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Welcome to the Tech Company listing API with Golang"})
}
