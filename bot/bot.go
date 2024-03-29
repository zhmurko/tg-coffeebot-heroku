package bot

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zhmurko/tg-coffeebot-heroku/cache"
	"github.com/zhmurko/tg-coffeebot-heroku/db"
	"log"
	"net/http"
	"strings"
)

var adminID = 234140659


// Respond parses an incoming json 
// and sends an output message
func Respond(c *gin.Context) {
	var chat Update
	err := c.ShouldBindJSON(&chat)
	if err != nil {
		log.Println(err)
	}
	jsonB, err := json.Marshal(chat)
	if err != nil {
		log.Println(err)
	}
	dumpPost(jsonB)
	id := chat.Message.Chat.ID + chat.CallbackQuery.From.ID
	text := chat.Message.Text + chat.CallbackQuery.Data
	log.Printf("ID: %d", id)
	log.Printf("R: %+v", chat)
	switch {
	case strings.HasPrefix(text, "order:"):
		who := chat.CallbackQuery.From.FirstName
		var coffee string
		order := strings.Split(text, ":")
		coffee = order[1]
		cache.RememberMe(fmt.Sprint(id), who)
		ReplyOrder(adminID, coffee, fmt.Sprint(id))
		SendText(id, "Doing "+coffee+" for you, "+who)
	case strings.HasPrefix(text, "ready:"):
		var coffee string
		var who int
		order := strings.Split(text, ":")
		coffee = order[1]
		_, _ = fmt.Sscanf(order[2], "%d", &who)
		messageID := chat.Message.ID + chat.CallbackQuery.Message.ID
		name := cache.WhatsMyName(fmt.Sprint(who))
		DeleteMessage(adminID, messageID)
		db.Add(who, coffee)
		SendText(adminID, "Completed "+coffee+" for "+name)
		SendText(who, "Your "+coffee+" is ready")
	case strings.HasPrefix(text, "/"):
		switch text {
		case "/menu":
			ReplyMenu(id)
		case "/stats":
			ReplyStats(id)
		default:
			SendText(id, "start order: /menu")
		}
	default:
		SendText(id, "type /menu")
	}
}

// Pong is used for healthchecks
func Pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
