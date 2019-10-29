package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)



func HandlePet(c *gin.Context) {
	type Pet struct {
		Name string `json:"name"`
		Species string `json:"species"`
	}

	var p Pet
	if err := c.BindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 1,
			"message": " bad request",
		})
	}

	fmt.Println("name: " + p.Name + ", species: " + p.Species)

	c.JSON(http.StatusOK,"pet ok")
}


func main() {
	fmt.Println("Hello JSON")
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message":"pong",
		})
	})

	auth := router.Group("/auth")
	auth.POST("/pet", HandlePet)

	router.Run(":8080")
}

