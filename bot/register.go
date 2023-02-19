package bot

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

// RegisterWebhook a callback of our service via Telegram API 
func RegisterWebhook() {

	token, ok := os.LookupEnv("BOT_TOKEN")
	if !ok {
		token = "empty"
	}

	appName, ok := os.LookupEnv("HEROKU_APP_NAME")
	if !ok {
		appName = "tg-coffeebot-heroku"
	}

	telegramURL := "https://api.telegram.org/bot" + token
	herokuURL := "https://" + appName + ".herokuapp.com/webhook"

	allowedUpdates, _ := json.Marshal([]string{"message", "edited_channel_post", "callback_query", "chat_member"})

	urlRequest := telegramURL + "/setWebhook?url=" + url.PathEscape(herokuURL) + "&allowed_updates=" + url.PathEscape(string(allowedUpdates))
	if token == "empty" {
		log.Printf("Skip Register an Empty Webhook:\n%s", urlRequest)
	} else {
		log.Printf("Register Webhook: %s", urlRequest)
		resp, err := http.Get(urlRequest)
		if err != nil {
			log.Fatalln(err)
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		//Convert the body to type string
		sb := string(body)
		log.Printf("%s",sb)
	}
}
