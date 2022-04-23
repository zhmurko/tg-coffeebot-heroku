package bot

import (
	//"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var adminId = 234140659

func Respond(c *gin.Context) {
	debug := c.Copy()
	dumpPost(debug)
	var chat Update
	err := c.ShouldBindJSON(&chat)
	if err != nil {
		log.Println(err)
	}
	id := chat.Message.Chat.Id
	log.Printf("ID: %d", id)
	log.Printf("R: %+v", chat)
	if chat.Message.Text == "/menu" {
		ReplyMenu(id)
	} else {
		SendText(id, "type /menu")
	}
}

func Pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
