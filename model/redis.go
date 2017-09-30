package model

import (
	"fmt"
	"log"

	"github.com/go-redis/redis"
	"github.com/spkills/spkills/config"
)

var redisClient *redis.Client

func InitRedis(conf config.Config) {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", conf.Redis.Server, conf.Redis.Port),
		Password: conf.Redis.Password,
		DB:       conf.Redis.DB,
	})
	_, err := redisClient.Ping().Result()
	if err != nil {
		log.Panic(err)
	}
}
