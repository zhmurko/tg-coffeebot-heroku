package bot

import (
	"encoding/json"
	//    "testing"
	"log"
)

type Button struct {
	Text         string `json:"text"`
	CallbackData string `json:"callback_data"`
}

type Markup struct {
	InlineKeyboard [][]Button `json:"inline_keyboard"`
}

type Menu struct {
	ChatId      int    `json:"chat_id"`
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

func jsonMenu(menu Menu) ([]byte, error) {
	return json.Marshal(menu)
}

func ReplyMenu(id int) []byte {
	var CoffeeMenu = Menu{
		ChatId: id,
		Text:   "Select:",
		ReplyMarkup: Markup{
			InlineKeyboard: [][]Button{
				{
					Espresso,
					Latte,
				},
			},
		},
	}
	menu, _ := jsonMenu(CoffeeMenu)
	body, err := Reply(menu)
	if err != nil {
		log.Fatalln(err)
	}
	return body
}
