package bot

import (
	"encoding/json"
	"testing"
	//    "log"
)

var OriginJSON = `{
    "chat_id": 1,
    "text": "Select",
    "reply_markup": {
        "inline_keyboard": [
            [
                {
                    "text": "Espresso",
                    "callback_data": "order:Espresso"
                },
                {
                    "text": "Latte",
                    "callback_data": "order:Latte"
                },
                {
                    "text": "Ice Coffee",
                    "callback_data": "order:Ice Coffee"
                }
            ]
        ]
    }
}`

var OriginMenu = Menu{
	ChatID: 1,
	Text:   "Select",
	ReplyMarkup: Markup{
		InlineKeyboard: [][]Button{{
			bEspresso,
			bLatte, 
			bIceCoffee},
		},
	},
}

func TestMenu(t *testing.T) {
	want := OriginJSON
	got, _ := json.MarshalIndent(OriginMenu, "", "    ")
	if want != string(got) {
		t.Errorf("want = %s; got %s", want, got)
	}
}

func TestJsonMenu(t *testing.T) {
	menu := Menu{ChatID: 1, ReplyMarkup: Markup{InlineKeyboard: [][]Button{}}}
	want := `{"chat_id":1,"text":"","reply_markup":{"inline_keyboard":[]}}`
	got, _ := jsonMenu(menu)
	if want != string(got) {
		t.Errorf("Error:\nwant = %s;\n got = %s", want, got)
	}
}
