package redis

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"codebase-go/bin/config"

	"github.com/go-redis/redis/v8"
)

var redisClient *redis.Client

func InitConnection(redisDB, redisHost, redisPort string) {
	db := 0
	parseRedisDb, err := strconv.ParseInt(redisDB, 10, 32)

	if err == nil {
		db = int(parseRedisDb)
	}

	c := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", redisHost, redisPort),
		Password: config.GetConfig().RedisPassword,
		DB:       db,
	})

	if c.Ping(context.Background()).Err() != nil {
		panic("cannot connect redis")
	}

	redisClient = c
}

func GetClient() Collections {
	return redisClient
}

type Collections interface {
	SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.BoolCmd
	EvalSha(ctx context.Context, sha1 string, keys []string, args ...interface{}) *redis.Cmd
	Del(ctx context.Context, keys ...string) *redis.IntCmd
	Conn(ctx context.Context) *redis.Conn
	Get(ctx context.Context, key string) *redis.StringCmd
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
}
