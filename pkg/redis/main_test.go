package redis

import (
	"testing"
)

var client CacheOperations

func TestMain(m *testing.M) {
	client = NewRedisCacheClient("localhost:6379", "", 0)
}
