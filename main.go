package main

import (
    "github.com/zhmurko/tg-coffeebot-heroku/bot"
    "log"
    //"net/http"
    "os"
    "github.com/gin-gonic/gin"
    _ "github.com/heroku/x/hmetrics/onload"
)

func main() {
    port := os.Getenv("PORT")

    token, ok := os.LookupEnv("BOT_TOKEN")
    if !ok {
        token = "empty"
    }    

    telegramUrl := "https://api.telegram.org/bot" + token;
    _ = telegramUrl
    if port == "" {
            log.Fatal("$PORT must be set")
    }
    
    router := gin.New()
    router.Use(gin.Logger())
    router.GET("/ping", bot.Pong)
    
    router.Run(":" + port)
}