package cache

import (
	"github.com/memcachier/mc/v3"
	"os"
)

var Cache *mc.Client

func Register() {
	servers := os.Getenv("MEMCACHIER_SERVERS")
	if servers == "" {
		servers = "localhost:11211"
	}
	username := os.Getenv("MEMCACHIER_USERNAME")
	password := os.Getenv("MEMCACHIER_PASSWORD")

	Cache := mc.NewMC(servers, username, password)
	defer Cache.Quit()
}
