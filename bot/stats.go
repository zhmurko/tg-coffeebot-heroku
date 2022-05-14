package bot

import (
	"fmt"
	_ "github.com/zhmurko/tg-coffeebot-heroku/cache"
	db "github.com/zhmurko/tg-coffeebot-heroku/db"
	"strings"
)

func getOrders(id int) []db.Order {
	var x db.Orders
	res := x.List(id)
	return res
}

func stats(arr []db.Order) string {
	sum := make(map[string]int)
	for _, x := range arr {
		sum[x.Coffee] += 1
	}
	var str []string
	for name, total := range sum {
		str = append(str, fmt.Sprintf("%s: %d", name, total))
	}
	return strings.Join(str[:], "\n")
}

func ReplyStats(id int) {
	orders := getOrders(id)
	message := stats(orders)
	SendText(id, message)
}
