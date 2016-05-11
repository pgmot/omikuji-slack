package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.Default()

	router.POST("/omikuji", func(c *gin.Context) {
		text := c.PostForm("text")
		log.Printf("%s", text)
	})
	router.Run(":" + port)
}
