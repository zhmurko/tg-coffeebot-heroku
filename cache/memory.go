package cache

import (
	"github.com/memcachier/mc/v3"
	"log"
)

func RememberMe(id string, name string) {
	_, err := Cache.Set(id, name, 0, 0, 0)
	if err != nil {
		log.Println(err)
	}
}

func WhatsMyName(id string) string {
	v, _, _, err := Cache.Get(id)
	if err != nil {
		if err == mc.ErrNotFound {
			return id
		} else {
			log.Println(err)
		}
	}
	return v
}
