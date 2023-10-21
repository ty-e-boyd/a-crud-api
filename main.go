package main

import (
	"a-crud-api/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "good call"})
	})

	err := r.Run()
	if err != nil {
		log.Fatal("error running server")
		return
	}
}
