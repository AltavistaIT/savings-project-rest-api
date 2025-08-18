package infra_redis

import (
	"context"

	"github.com/redis/go-redis/v9"
	"github.com/ssssshel/sp-api/src/shared/config"
	"github.com/ssssshel/sp-api/src/shared/logger"
)

type RedisConnection struct {
	RedisConn *redis.Client
}

var Ctx = context.Background()

func InitRedisConnection() (*RedisConnection, error) {
	config := config.GetConfig()

	redisConn := redis.NewClient(&redis.Options{
		Addr:     config.RedisAddr,
		Password: config.RedisPassword,
		DB:       config.RedisDB,
	})

	_, err := redisConn.Ping(Ctx).Result()
	if err != nil {
		logger.Fatal("Error connecting to redis => ", err)
		return nil, err
	}

	logger.Info("Redis connection successful")
	return &RedisConnection{RedisConn: redisConn}, nil
}
