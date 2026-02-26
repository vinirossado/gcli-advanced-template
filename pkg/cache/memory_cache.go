package cache

import (
	"context"
	"log"
	"time"

	"github.com/allegro/bigcache/v3"
)

var Cache *bigcache.BigCache

func MemoryCache() {
	var err error
	Cache, err = bigcache.New(context.Background(), bigcache.DefaultConfig(10*time.Minute))
	if err != nil {
		log.Fatalf("failed to initialize in-memory cache: %v", err)
	}
}
