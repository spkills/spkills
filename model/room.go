package model

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis"
	"github.com/spkills/spkills/repository"
)

func FetchRoomNameByDB(roomId int) string {
	room, _ := repository.FindRoomG(int64(roomId))
	return room.Name
}

func FetchRoomNameByCache(roomId int) string {
	key := fmt.Sprintf("roomName:%d", roomId)
	var roomName string
	val, found := Cache.Get(key)
	if found {
		roomName = val.(string)
		//fmt.Printf("cache(db)\n")
	} else {
		room, err := repository.FindRoomG(int64(roomId))
		if err != nil {
			//
		}
		roomName = room.Name
		Cache.Set(key, roomName, 10*time.Second)
		//fmt.Printf("cache(cache)\n")
	}
	//fmt.Printf("cache %+v\n", room)
	return roomName
}
func FetchRoomNameByRedis(roomId int) string {
	key := fmt.Sprintf("roomName:%d", roomId)
	var roomName string
	val, err := redisClient.Get(key).Result()
	if err == redis.Nil {
		room, err := repository.FindRoomG(int64(roomId))
		if err != nil {
			//
		}
		roomName = room.Name
		err = redisClient.Set(key, roomName, 10*time.Second).Err()
		if err != nil {
			panic(err)
		}
		//fmt.Printf("redis(db)\n")
	} else if err != nil {
		panic(err)
	} else {
		roomName = val
		//fmt.Printf("redis(cache)\n")
	}
	//fmt.Printf("redis %+v\n", room)
	return roomName
}

func FetchRoomByDB(roomId int) *repository.Room {
	room, _ := repository.FindRoomG(int64(roomId))
	return room
}

func FetchRoomByCache(roomId int) *repository.Room {
	key := fmt.Sprintf("room:%d", roomId)
	var room *repository.Room
	var err error
	val, found := Cache.Get(key)
	if found {
		room = val.(*repository.Room)
		//fmt.Printf("cache(db)\n")
	} else {
		room, err = repository.FindRoomG(int64(roomId))
		if err != nil {
			//
		}
		Cache.Set(key, room, 10*time.Second)
		//fmt.Printf("cache(cache)\n")
	}
	//fmt.Printf("cache %+v\n", room)
	return room
}

func FetchRoomByRedis(roomId int) *repository.Room {
	key := fmt.Sprintf("room:%d", roomId)
	var room *repository.Room
	var err error
	val, err := redisClient.Get(key).Result()
	if err == redis.Nil {
		room, err = repository.FindRoomG(int64(roomId))
		if err != nil {
			//
		}
		// いったんjsonでbyteにしてから、文字列を渡す必要がある
		out, err := json.Marshal(room)
		if err != nil {
			panic(err)
		}
		err = redisClient.Set(key, string(out), 10*time.Second).Err()
		if err != nil {
			panic(err)
		}
		//fmt.Printf("redis(db)\n")
	} else if err != nil {
		panic(err)
	} else {
		// いったん文字列をbyte配列にメモリコピーしてデシリアライズする
		room = &repository.Room{}
		err := json.Unmarshal([]byte(val), room)
		if err != nil {
			panic(err)
		}
		//fmt.Printf("redis(cache)\n")
	}
	//fmt.Printf("redis %+v\n", room)
	return room
}
