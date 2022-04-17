package bot

import (
	"encoding/json"
	"testing"
)

var OriginJson = `{
    "chat_id": "1",
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
                }
            ]
        ]
    }
}`

var OriginMenu = Menu{
	ChatId: "1",
	Text:   "Select",
	ReplyMarkup: Markup{
		InlineKeyboard: [][]Button{{
			Espresso,
			Latte},
		},
	},
}

func TestMenu(t *testing.T) {
	want := OriginJson
	got, _ := json.MarshalIndent(OriginMenu, "", "    ")
	if want != string(got) {
		t.Errorf("want = %s; got %s", want, got)
	}
}
