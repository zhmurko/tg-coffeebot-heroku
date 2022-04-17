package bot

import (
	"github.com/gin-gonic/gin"
	"net/http"
	//    "encoding/json"
	"log"
)

func Respond(c *gin.Context) {
	dumpPost(c)
	// var chat Update
	// err := c.ShouldBindJSON(&chat)
	// if err != nil {
	// 	log.Println(err)
	// }
	// id := chat.Message.Chat.Id
	// log.Printf("ID: %d", id)
	// log.Printf("R: %+v", chat)
	// if chat.Message.Text == "/menu" {
	// 	ReplyMenu(id)
	// } else {
	// 	SendText(id, "type /menu")
	// }
}

func Pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
