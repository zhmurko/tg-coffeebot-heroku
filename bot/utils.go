package bot

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
)

func prettyprint(b []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Indent(&out, b, "|", "  ")
	return out.Bytes(), err
}

func dumpPost(c *gin.Context) {
	jsonData, err := c.GetRawData()
	if err != nil {
		log.Fatalln(err)
	}
	b, _ := prettyprint(jsonData)
	log.Println("Respond " + string(b))
}

