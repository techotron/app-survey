package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Setup router 
func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello"})
	})

	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
