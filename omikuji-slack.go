package main

import (
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"

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
		nakamiLen := len(nakami)
		if nakamiLen == 0 {
			c.String(http.StatusOK, "example) /omikuji daikichi,kichi,kyou,daikyou")
		} else {
			c.String(http.StatusOK, nakami[rand.Intn(nakamiLen)])
		}
	})
	router.Run(":" + port)
}
