package main

import (
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

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

		nakami := strings.Split(text, ",")
		rand.Seed(time.Now().UnixNano())
		c.JSON(http.StatusOK, gin.H{
			"text":          nakami[rand.Intn(len(nakami))],
			"response_type": "in_channel",
		})
	})
	router.Run(":" + port)
}
