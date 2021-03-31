// Package redis impl functions for accessing redis cache
package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

// CacheOperations defines methods for interacting with cache
type CacheOperations interface {
	Set(ctx context.Context, key string, value interface{}, duration time.Duration) error
	GetString(ctx context.Context, key string) (string, error)
}

type redisCache struct {
	client *redis.Client
}

// NewRedisCacheClient creates new instance of redisCache
func NewRedisCacheClient(address, password string, db int) CacheOperations {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       db,
	})
	return &redisCache{client: redisClient}
}

func (r *redisCache) Set(ctx context.Context, key string, value interface{}, duration time.Duration) error {
	err := r.client.Set(ctx, key, value, duration)
	if err != nil {
		return err.Err()
	}
	return nil
}

func (r *redisCache) GetString(ctx context.Context, key string) (string, error) {
	val, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}
