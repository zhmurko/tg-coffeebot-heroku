package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/zhmurko/tg-coffeebot-heroku/bot"
	"github.com/zhmurko/tg-coffeebot-heroku/cache"
	"os"
	"sync"
)

var doOnce sync.Once

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	router := gin.New()
	router.SetTrustedProxies(nil)
	router.Use(gin.Logger())
	router.GET("/ping", bot.Pong)
	router.POST("/webhook", bot.Respond)

	doOnce.Do(func() {
		bot.RegisterWebhook()
		cache.Register()
	})
	router.Run(":" + port)

}
