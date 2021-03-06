package cache

import (
	"github.com/memcachier/mc/v3"
	"os"
)

var Cache *mc.Client

func Register() *mc.Client {
	servers := os.Getenv("MEMCACHIER_SERVERS")
	if servers == "" {
		servers = "localhost:11211"
	}
	username := os.Getenv("MEMCACHIER_USERNAME")
	password := os.Getenv("MEMCACHIER_PASSWORD")

	return mc.NewMC(servers, username, password)
}
