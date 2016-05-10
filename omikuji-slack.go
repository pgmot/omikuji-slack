package main

import (
	"fmt"
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
		fmt.Printf("text: %s", text)
	})
	router.Run(":" + port)
}
