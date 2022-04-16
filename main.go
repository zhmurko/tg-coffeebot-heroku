package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/zhmurko/tg-coffeebot-heroku/bot"
	"log"
	"os"
	"sync"
)

var doOnce sync.Once

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.GET("/ping", bot.Pong)
	router.POST("/webhook", bot.Respond)

	doOnce.Do(func() {
		bot.RegisterWebhook()
	})
	router.Run(":" + port)

}
