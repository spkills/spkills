package model

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/go-redis/redis"
	"github.com/spkills/spkills/config"
)

var redisClient *redis.Client

func InitRedis(conf config.Config) {
	var addr string
	var network string
	var dialer func() (net.Conn, error)
	_, err := os.Stat(conf.Database.SocketFile)
	if err == nil {
		network = "unix"
		dialer = func() (net.Conn, error) {
			conn, err := net.Dial("unix", conf.Redis.SocketFile)
			if err != nil {
				log.Fatalf("error: %v", err)
			}
			return conn, err
		}
	} else {
		network = "tcp"
		addr = fmt.Sprintf("%s:%s", conf.Redis.Server, conf.Redis.Port)
	}

	redisClient = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: conf.Redis.Password,
		DB:       conf.Redis.DB,
		Network:  network,
		Dialer:   dialer,
	})
	// _, err := redisClient.Ping().Result()
	// if err != nil {
	// 	log.Panic(err)
	// }
}
