package bot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func telegramUrl(uri string) string {
	token, ok := os.LookupEnv("BOT_TOKEN")
	if !ok {
		token = "empty"
	}

	telegramUrl := "https://api.telegram.org/bot" + token + uri

	return telegramUrl
}

func Reply(b []byte) ([]byte, error) {
	url := telegramUrl("/sendMessage")
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(b))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func DeleteMessage(chat_id int, message_id int) {
	args := fmt.Sprintf("chat_id=%d&message_id=%d", chat_id, message_id)
	url := telegramUrl("/deleteMessage?" + args)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	_, _ = ioutil.ReadAll(resp.Body)
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
