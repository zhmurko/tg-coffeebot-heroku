package bot

import (
	"encoding/json"
	//    "testing"
	"github.com/zhmurko/tg-coffeebot-heroku/cache"
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

var IceCoffee = Button{
	Text:         "Ice Coffee",
	CallbackData: "order:Ice Coffee",
}

var Mocaccino = Button{
	Text:         "Mocaccino",
	CallbackData: "order:Mocaccino",
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
					IceCoffee,
					Mocaccino,
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

func ReplyOrder(id int, coffee string, who string) []byte {
	var Order = Button{
		Text:         coffee,
		CallbackData: "ready:" + coffee + ":" + who,
	}
	name := cache.WhatsMyName(who)
	var OrderMenu = Menu{
		ChatId: id,
		Text:   "Prepare please " + coffee + " for " + name + ":",
		ReplyMarkup: Markup{
			InlineKeyboard: [][]Button{
				{
					Order,
				},
			},
		},
	}
	menu, _ := jsonMenu(OrderMenu)
	body, err := Reply(menu)
	if err != nil {
		log.Fatalln(err)
	}
	return body
}
