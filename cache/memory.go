package cache

import (
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
		log.Println(err)
	}
	if v == "Not found" {
		return id
	} else {
		return v
	}
}
