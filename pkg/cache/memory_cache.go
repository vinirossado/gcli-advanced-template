package cache

import (
	"context"
	"fmt"
	"github.com/allegro/bigcache/v3"
	"time"
)

var Cache *bigcache.BigCache

func MemoryCache() {
	var err error
	Cache, err = bigcache.New(context.Background(), bigcache.DefaultConfig(10*time.Minute))

	if err != nil {
		fmt.Println("Error while creating cache", err)
		return
	}
}
