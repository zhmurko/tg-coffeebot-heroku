package bot

import (
//	"encoding/json"
	"github.com/gin-gonic/gin"
	//"io/ioutil"
	"log"
	"net/http"
	//"net/url"
	//"os"
)

func Respond(c *gin.Context) {
  log.Printf("Respond")
}

func Pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
