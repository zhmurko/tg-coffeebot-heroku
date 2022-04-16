package bot

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	//"io/ioutil"
	"log"
	"net/http"
	//"net/url"
	//"os"
  "bytes"
)

func prettyprint(b []byte) ([]byte, error) {
    var out bytes.Buffer
    err := json.Indent(&out, b, "", "  ")
    return out.Bytes(), err
}

func Respond(c *gin.Context) {
  jsonData, err := c.GetRawData()
  if err != nil{
     log.Fatalln(err)
  }
  b, _ := prettyprint(jsonData)
  log.Println("Respond " + string(b))
}

func Pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
