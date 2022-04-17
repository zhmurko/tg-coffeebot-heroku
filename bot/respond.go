package bot

import (
//"encoding/json"
)

type From struct {
	Id           int    `json:"id"`
	IsBot        bool   `json:"is_bot"`
	FirstName    string `json:"first_name"`
	Username     string `json:"username"`
	LanguageCode string `json:"language_code"`
}

type Chat struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	Username  string `json:"username"`
	Type      string `json:"type"`
}

type Message struct {
	Id   int    `json:"message_id"`
	From From   `json:"from"`
	Chat Chat   `json:"chat"`
	Date int    `json:"date"`
	Text string `json:"text"`
}

type Update struct {
	UpdateId int     `json:"update_id"`
	Message  Message `json:"message"`
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