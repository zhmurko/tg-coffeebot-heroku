package bot

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"encoding/json"
	"log"
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
		SendText(id, "Preare: "+text)
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
