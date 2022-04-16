package bot

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

func RegisterWebhook() {

	token, ok := os.LookupEnv("BOT_TOKEN")
	if !ok {
		token = "empty"
	}

	appName, ok := os.LookupEnv("HEROKU_APP_NAME")
	if !ok {
		appName = "tg-coffeebot-heroku"
	}

	telegramUrl := "https://api.telegram.org/bot" + token
	herokuUrl := "https://" + appName + ".herokuapp.com/webhook"

	allowed_updates, _ := json.Marshal([]string{"message", "edited_channel_post", "callback_query", "chat_member"})

	urlRequest := telegramUrl + "/setWebhook?url=" + url.PathEscape(herokuUrl) + "&allowed_updates=" + url.PathEscape(string(allowed_updates))
	log.Printf("Register Webhook: %s", urlRequest)
	resp, err := http.Get(urlRequest)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	sb := string(body)
	log.Printf(sb)
	//
}
