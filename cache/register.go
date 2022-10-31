package cache

import (
	"github.com/memcachier/mc/v3"
	"os"
)

// connection instance
var Cache *mc.Client

// inititate a connection to Memcached service
func Register() *mc.Client {
	servers := os.Getenv("MEMCACHIER_SERVERS")
	if servers == "" {
		servers = "localhost:11211"
	}
	username := os.Getenv("MEMCACHIER_USERNAME")
	password := os.Getenv("MEMCACHIER_PASSWORD")

	return mc.NewMC(servers, username, password)
}
