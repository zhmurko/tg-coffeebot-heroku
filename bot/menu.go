package bot

import (
	"encoding/json"
	//    "testing"
	"github.com/zhmurko/tg-coffeebot-heroku/cache"
	"log"
)

// Button in telegram chat
type Button struct {
	Text         string `json:"text"`
	CallbackData string `json:"callback_data"`
}

// Markup is a block of buttons in telegram chat
type Markup struct {
	InlineKeyboard [][]Button `json:"inline_keyboard"`
}

// Menu of full telegram message with buttons
type Menu struct {
	ChatId      int    `json:"chat_id"`
	Text        string `json:"text"`
	ReplyMarkup Markup `json:"reply_markup"`
}

// Espresso button for menu
var bEspresso = Button{
	Text:         "Espresso",
	CallbackData: "order:Espresso",
}

// Latte button for menu
var bLatte = Button{
	Text:         "Latte",
	CallbackData: "order:Latte",
}

var bIceCoffee = Button{
	Text:         "Ice Coffee",
	CallbackData: "order:Ice Coffee",
}

var bMocaccino = Button{
	Text:         "Mocaccino",
	CallbackData: "order:Mocaccino",
}

func jsonMenu(menu Menu) ([]byte, error) {
	return json.Marshal(menu)
}

// ReplyMenu prepares an Menu message with button
func ReplyMenu(id int) []byte {
	var CoffeeMenu = Menu{
		ChatId: id,
		Text:   "Select:",
		ReplyMarkup: Markup{
			InlineKeyboard: [][]Button{
				{
					bEspresso,
					bLatte,
					bIceCoffee,
					bMocaccino,
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

// ReplyOrder prepares a telegram message with order details
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
