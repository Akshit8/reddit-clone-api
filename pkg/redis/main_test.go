package redis

import (
	"log"
	"os"
	"testing"
)

var client CacheOperations

func TestMain(m *testing.M) {
	var err error
	client, err = NewRedisCacheClient("redis://localhost:6379")
	if err != nil {
		log.Fatal("error creating redis client: ", err)
	}
	os.Exit(m.Run())
}
