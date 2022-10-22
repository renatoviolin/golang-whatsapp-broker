package redis

import (
	"broker/infra/logger"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis/v9"
)

type RedisClient struct {
	Client  *redis.Client
	Expires time.Duration
}

func NewRedisClient(address string) *RedisClient {
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: "",
	})
	if err := client.Ping(context.Background()).Err(); err != nil {
		logger.Fatal("new-redis-client", err.Error())
	}
	// expires := time.Second * 600 // 5min
	expires := time.Second * 86400 // 24h
	return &RedisClient{Client: client, Expires: expires}
}

func (h *RedisClient) Save(key string, value string) error {
	_, err := h.Client.Set(context.Background(), key, value, h.Expires).Result()
	return err
}

func (h *RedisClient) Delete(key string) error {
	_, err := h.Client.Del(context.Background(), key).Result()
	return err
}

func (h *RedisClient) Get(key string) (res string, err error) {
	res, err = h.Client.Get(context.Background(), key).Result()
	if err != nil {
		return res, err
	}
	if res == "" {
		return res, fmt.Errorf("key %s not found", key)
	}

	return res, nil
}

func (h *RedisClient) GetTTL(key string) (duration int64, err error) {
	res, err := h.Client.TTL(context.Background(), key).Result()
	if err != nil || res == -2 {
		return -2, errors.New("expired or invalid key")
	}

	return int64(res / time.Second), nil
}

func (h *RedisClient) ClearAll() (err error) {
	ctx := context.Background()
	keys, err := h.Client.Keys(ctx, "*").Result()
	if err != nil || len(keys) == 0 {
		return err
	}
	_, err = h.Client.Pipelined(ctx, func(pipe redis.Pipeliner) error {
		_, err = pipe.Del(ctx, keys...).Result()
		return err
	})
	if err != nil {
		panic(err)
	}
	return nil
}

func (h *RedisClient) Touch(key string) error {
	_, err := h.Client.Expire(context.Background(), key, h.Expires).Result()
	return err
}
