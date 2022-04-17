package bot

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func Reply(b []byte) ([]byte, error) {
	token, ok := os.LookupEnv("BOT_TOKEN")
	if !ok {
		token = "empty"
	}

	telegramUrl := "https://api.telegram.org/bot" + token

	url := telegramUrl + "/sendMessage"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

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
