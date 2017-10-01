package model

import (
	"fmt"
	"os"

	"github.com/go-redis/redis"
	"github.com/spkills/spkills/config"
)

var redisClient *redis.Client

func InitRedis(conf config.Config) {
	var addr string
	var network string
	_, err := os.Stat(conf.Redis.SocketFile)
	if err == nil {
		network = "unix"
		addr = conf.Redis.SocketFile
	} else {
		network = "tcp"
		addr = fmt.Sprintf("%s:%d", conf.Redis.Server, conf.Redis.Port)
	}

	redisClient = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: conf.Redis.Password,
		DB:       conf.Redis.DB,
		Network:  network,
	})
	// _, err := redisClient.Ping().Result()
	// if err != nil {
	// 	log.Panic(err)
	// }
}
