package bot

import (
//    "encoding/json"
//    "testing"
//    "log"
)

type Button struct {
	Text         string `json:"text"`
	CallbackData string `json:"callback_data"`
}

type Markup struct {
	InlineKeyboard [][]Button `json:"inline_keyboard"`
}

type Menu struct {
	ChatId      string `json:"chat_id"`
	Text        string `json:"text"`
	ReplyMarkup Markup `json:"reply_markup"`
}

var Espresso = Button{
	Text:         "Espresso",
	CallbackData: "order:Espresso",
}

var Latte = Button{
	Text:         "Latte",
	CallbackData: "order:Latte",
}

var CoffeeMenu = Menu{
	ChatId: "1",
	Text:   "select",
	ReplyMarkup: Markup{
		InlineKeyboard: [][]Button{
			{
				Espresso,
				Latte,
			},
		},
	},
}
