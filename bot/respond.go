package bot

import (
//"encoding/json"
)

// From converts a Telegram JSON response into go structure
type From struct {
	ID           int    `json:"id"`
	IsBot        bool   `json:"is_bot"`
	FirstName    string `json:"first_name"`
	Username     string `json:"username"`
	LanguageCode string `json:"language_code"`
}

// Chat converts a Telegram JSON response into go structure
type Chat struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	Username  string `json:"username"`
	Type      string `json:"type"`
}

// Message converts a Telegram JSON response into go structure
type Message struct {
	ID   int    `json:"message_id"`
	From From   `json:"from"`
	Chat Chat   `json:"chat"`
	Date int    `json:"date"`
	Text string `json:"text"`
}

// Update converts a Telegram JSON response into go structure
type Update struct {
	UpdateID      int           `json:"update_id"`
	Message       Message       `json:"message"`
	CallbackQuery CallbackQuery `json:"callback_query"`
}

// CallbackQuery converts a Telegram JSON response into go structure
type CallbackQuery struct {
	ID           string  `json:"id"`
	From         From    `json:"from"`
	ChatInstance string  `json:"chat_instance"`
	Data         string  `json:"data"`
	Message      Message `json:"message"`
}

// Respond
// |{
// |  "update_id": 841266257,
// |  "message": {
// |    "message_id": 698,
// |    "from": {
// |      "id": 234140659,
// |      "is_bot": false,
// |      "first_name": "Dmitry",
// |      "username": "zhmurko",
// |      "language_code": "en"
// |    },
// |    "chat": {
// |      "id": 234140659,
// |      "first_name": "Dmitry",
// |      "username": "zhmurko",
// |      "type": "private"
// |    },
// |    "date": 1650205024,
// |    "text": "test"
// |  }
// |}

// Respond with callback_query
// |{
// |  "update_id": 841266275,
// |  "callback_query": {
// |    "id": "1005626476713560102",
// |    "from": {
// |      "id": 234140659,
// |      "is_bot": false,
// |      "first_name": "Dmitry",
// |      "username": "zhmurko",
// |      "language_code": "en"
// |    },
// |    "message": {
// |      "message_id": 713,
// |      "from": {
// |        "id": 5176662516,
// |        "is_bot": true,
// |        "first_name": "Coffee, please",
// |        "username": "OneCoffeebot"
// |      },
// |      "chat": {
// |        "id": 234140659,
// |        "first_name": "Dmitry",
// |        "username": "zhmurko",
// |        "type": "private"
// |      },
// |      "date": 1650212725,
// |      "text": "Select:",
// |      "reply_markup": {
// |        "inline_keyboard": [
// |          [
// |            {
// |              "text": "Espresso",
// |              "callback_data": "order:Espresso"
// |            },
// |            {
//               "text": "Latte",
// |              "callback_data": "order:Latte"
// |            }
// |          ]
// |        ]
// |      }
// |    },
// |    "chat_instance": "-2656029114737977985",
// |    "data": "order:Espresso"
// |  }
// |}
//
