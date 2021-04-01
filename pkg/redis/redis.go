// Package redis impl functions for accessing redis cache
package redis

import (
	"time"

	"github.com/go-redis/redis"
)

// CacheOperations defines methods for interacting with cache
type CacheOperations interface {
	Set(key string, value interface{}, duration time.Duration) error
	Delete(key string) error
	GetString(key string) (string, error)
}

type redisCache struct {
	client *redis.Client
}

// NewRedisCacheClient creates new instance of redisCache
func NewRedisCacheClient(redisURI string) (CacheOperations, error) {
	rc := &redisCache{}
	options, err := redis.ParseURL(redisURI)
	if err != nil {
		return nil, err
	}
	rc.client = redis.NewClient(options)
	_, err = rc.client.Ping().Result()
	if err != nil {
		return nil, err
	}
	return rc, nil
}

func (r *redisCache) Set(key string, value interface{}, duration time.Duration) error {
	err := r.client.Set(key, value, duration)
	if err != nil {
		return err.Err()
	}
	return nil
}

func (r *redisCache) GetString(key string) (string, error) {
	val, err := r.client.Get(key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

func (r *redisCache) Delete(key string) error {
	_, err := r.client.Del(key).Result()
	if err != nil {
		return err
	}
	return nil
}
