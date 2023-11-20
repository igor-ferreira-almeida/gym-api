package exercise

import (
	"context"
	"fmt"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
	"strings"
)

var (
	isLocal    = true
	isInMemory = true
)

type redisCache struct {
	client *redis.Client
}

func (c redisCache) Set(ctx context.Context, e Exercise) error {
	key := strings.ReplaceAll(strings.ToLower(fmt.Sprintf("exercise:%s:%s", e.MuscleGroup, e.Name)), " ", "_")
	item, err := e.GetByte()
	if err != nil {
		return err
	}
	set := c.client.Set(ctx, key, item, redis.KeepTTL)
	fmt.Println(set)
	return nil
}

func newRedisCache() redisCache {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return redisCache{client: client}
}

var providerRedisCache = wire.NewSet(
	newRedisCache,
)
