package bot

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

var adminId = 234140659

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
	id := chat.Message.Chat.Id + chat.CallbackQuery.From.Id
	text := chat.Message.Text + chat.CallbackQuery.Data
	log.Printf("ID: %d", id)
	log.Printf("R: %+v", chat)
	switch {
	case strings.HasPrefix(text, "order:"):
		var coffee string
		_, _ = fmt.Sscanf(text, "order:%s", &coffee)
		ReplyOrder(adminId, coffee, fmt.Sprint(id))
		SendText(id, "Doing: "+coffee)
	case strings.HasPrefix(text, "ready:"):
		var coffee string
		var who int
		_, _ = fmt.Sscanf(text, "ready:%s:%d", &coffee, &who)
		DeleteMessage(adminId, chat.Message.Id)
		SendText(adminId, "Ready")
		SendText(int(who), "Your "+coffee+" is ready")
	case strings.HasPrefix(text, "/"):
		switch text {
		case "/menu":
			ReplyMenu(id)
		}
	default:
		SendText(id, "type /menu")
	}
}

func Pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
