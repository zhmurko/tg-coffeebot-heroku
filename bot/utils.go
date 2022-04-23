package bot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
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

func SendText(id int, text string) {
	data := []byte(fmt.Sprintf(`{"chat_id":%d,"text":"%s"}`, id, text))
	_, _ = Reply(data)
}

func prettyprint(b []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Indent(&out, b, "|", "  ")
	return out.Bytes(), err
}

func dumpPost(jsonB []byte) {
	b, _ := prettyprint(jsonB)
	log.Println("Respond " + string(b))
}
