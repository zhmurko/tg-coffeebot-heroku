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

func telegramURL(uri string) string {
	token, ok := os.LookupEnv("BOT_TOKEN")
	if !ok {
		token = "empty"
	}

	telegramURL := "https://api.telegram.org/bot" + token + uri

	return telegramURL
}

// Reply call a main webhook for Telegram API for /sendMessage
func Reply(b []byte) ([]byte, error) {
	url := telegramURL("/sendMessage")
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

// DeleteMessage clean up a telegram chat history from pressed buttons
func DeleteMessage(chatID int, messageID int) {
	args := fmt.Sprintf("chat_id=%d&message_id=%d", chatID, messageID)
	url := telegramURL("/deleteMessage?" + args)
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

// SendText replies with a JSON-formatted message
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
